package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

// Native WhatsApp polls.
//
// Polls are the only interactive primitive this bridge exposes on purpose:
// they are a first-class protocol message (PollCreationMessage), unlike
// buttons/list messages, which are a Business-API surface that personal
// accounts get flagged for emitting.
//
// Vote plumbing is the non-obvious part. A vote arrives as a
// PollUpdateMessage whose payload is encrypted with a secret derived from
// the *poll's* MessageSecret, and once decrypted it identifies the choice
// only by the SHA-256 of the option text — never the text itself. So to
// render "who voted what" the bridge has to keep the original option list
// around: hence the polls table, written both when we send a poll and when
// we observe someone else's.
const (
	maxPollQuestionLen = 255
	maxPollOptionLen   = 100
	minPollOptions     = 2
	maxPollOptions     = 12
)

// SendPollRequest is the request body for POST /api/poll.
type SendPollRequest struct {
	Recipient string   `json:"recipient"`
	Question  string   `json:"question"`
	Options   []string `json:"options"`
	// AllowMultiple is the ergonomic switch: false (default) = single choice,
	// true = voters may pick every option.
	AllowMultiple bool `json:"allow_multiple,omitempty"`
	// SelectableOptionsCount overrides AllowMultiple when > 0, for callers
	// that want "pick up to N".
	SelectableOptionsCount int `json:"selectable_options_count,omitempty"`
}

// selectableCount maps the two request knobs onto whatsmeow's single
// selectableOptionCount field.
func (req SendPollRequest) selectableCount() int {
	if req.SelectableOptionsCount > 0 {
		return req.SelectableOptionsCount
	}
	if req.AllowMultiple {
		return len(normalizePollOptions(req.Options))
	}
	return 1
}

// normalizePollOptions trims each option. WhatsApp hashes the option text
// verbatim, so trailing whitespace would silently produce a hash that no
// longer matches what we stored.
func normalizePollOptions(options []string) []string {
	out := make([]string, 0, len(options))
	for _, opt := range options {
		out = append(out, strings.TrimSpace(opt))
	}
	return out
}

// validatePollRequest enforces WhatsApp's limits before we hit the wire.
func validatePollRequest(req SendPollRequest) error {
	if strings.TrimSpace(req.Recipient) == "" {
		return fmt.Errorf("recipient is required")
	}

	question := strings.TrimSpace(req.Question)
	if question == "" {
		return fmt.Errorf("question is required")
	}
	if utf8.RuneCountInString(question) > maxPollQuestionLen {
		return fmt.Errorf("question is %d characters, limit is %d",
			utf8.RuneCountInString(question), maxPollQuestionLen)
	}

	options := normalizePollOptions(req.Options)
	if len(options) < minPollOptions || len(options) > maxPollOptions {
		return fmt.Errorf("poll needs between %d and %d options, got %d",
			minPollOptions, maxPollOptions, len(options))
	}

	seen := make(map[string]int, len(options))
	for i, opt := range options {
		if opt == "" {
			return fmt.Errorf("option %d is empty", i+1)
		}
		if utf8.RuneCountInString(opt) > maxPollOptionLen {
			return fmt.Errorf("option %d is %d characters, limit is %d",
				i+1, utf8.RuneCountInString(opt), maxPollOptionLen)
		}
		// Duplicate option text is not merely ugly: both options hash to the
		// same SHA-256, so an incoming vote could not be attributed to one of
		// them. Reject at the door rather than store an ambiguous poll.
		if first, dup := seen[opt]; dup {
			return fmt.Errorf("option %d duplicates option %d (%q); votes are identified by a hash of the option text, so duplicates are indistinguishable",
				i+1, first+1, opt)
		}
		seen[opt] = i
	}

	if n := req.selectableCount(); n < 1 || n > len(options) {
		return fmt.Errorf("selectable_options_count must be between 1 and %d, got %d", len(options), n)
	}

	return nil
}

// ---------- persistence ----------

// StorePoll records a poll's question and options so later votes (which carry
// only option hashes) can be rendered as text. Called for both directions:
// polls we send and polls we merely observe.
func (store *MessageStore) StorePoll(id, chatJID, question string, options []string, selectableCount int, timestamp time.Time, isFromMe bool) error {
	encoded, err := json.Marshal(options)
	if err != nil {
		return fmt.Errorf("failed to encode poll options: %w", err)
	}

	_, err = store.db.Exec(
		`INSERT INTO polls (id, chat_jid, question, options, selectable_count, timestamp, is_from_me)
		 VALUES (?, ?, ?, ?, ?, ?, ?)
		 ON CONFLICT(id, chat_jid) DO UPDATE SET
			question = excluded.question,
			options = excluded.options,
			selectable_count = excluded.selectable_count,
			timestamp = excluded.timestamp,
			is_from_me = excluded.is_from_me`,
		id, chatJID, question, string(encoded), selectableCount, timestamp, isFromMe,
	)
	return err
}

// GetPoll returns a stored poll's question and option list. Returns
// sql.ErrNoRows when the poll predates the bridge (its options are then
// unrecoverable — the hashes are one-way).
func (store *MessageStore) GetPoll(id, chatJID string) (string, []string, error) {
	var question, encoded string
	err := store.db.QueryRow(
		`SELECT question, options FROM polls WHERE id = ? AND chat_jid = ?`,
		id, chatJID,
	).Scan(&question, &encoded)
	if err != nil {
		return "", nil, err
	}

	var options []string
	if err := json.Unmarshal([]byte(encoded), &options); err != nil {
		return question, nil, fmt.Errorf("failed to decode stored poll options: %w", err)
	}
	return question, options, nil
}

// StorePollVote records one decrypted vote. A voter changing their mind
// produces a new PollUpdateMessage with a new ID, so rows accumulate and the
// latest vote per voter is the one with the highest timestamp.
func (store *MessageStore) StorePollVote(voteID, pollID, chatJID, voter string, selected []string, votedAt time.Time) error {
	encoded, err := json.Marshal(selected)
	if err != nil {
		return fmt.Errorf("failed to encode selected options: %w", err)
	}

	_, err = store.db.Exec(
		`INSERT INTO poll_votes (id, poll_id, chat_jid, voter, selected_options, timestamp)
		 VALUES (?, ?, ?, ?, ?, ?)
		 ON CONFLICT(id, chat_jid) DO UPDATE SET
			poll_id = excluded.poll_id,
			voter = excluded.voter,
			selected_options = excluded.selected_options,
			timestamp = excluded.timestamp`,
		voteID, pollID, chatJID, voter, string(encoded), votedAt,
	)
	return err
}

// ---------- rendering ----------

// pollOptionNames pulls the option text out of a PollCreationMessage.
func pollOptionNames(poll *waE2E.PollCreationMessage) []string {
	options := make([]string, 0, len(poll.GetOptions()))
	for _, opt := range poll.GetOptions() {
		options = append(options, opt.GetOptionName())
	}
	return options
}

// formatPollContent renders a poll for the messages table, so that every
// existing reader of messages.db (list_messages, the MCP tools, the digest
// scripts) sees something meaningful instead of an empty row.
func formatPollContent(question string, options []string) string {
	if len(options) == 0 {
		return fmt.Sprintf("[poll] %s", question)
	}
	return fmt.Sprintf("[poll] %s | %s", question, strings.Join(options, " / "))
}

// formatPollVoteContent renders a decrypted vote in the same spirit.
func formatPollVoteContent(question string, selected []string) string {
	if question == "" {
		question = "(unknown poll)"
	}
	if len(selected) == 0 {
		// An empty selection is how WhatsApp represents "I took my vote back".
		return fmt.Sprintf("[poll-vote] %s -> (vote cleared)", question)
	}
	return fmt.Sprintf("[poll-vote] %s -> %s", question, strings.Join(selected, ", "))
}

// resolvePollVoteOptions maps the SHA-256 hashes in a decrypted vote back to
// option text. Hashes with no match (poll created before the bridge saw it,
// or an option added later) are rendered as a short hash prefix rather than
// dropped, so the row still records that *something* was voted.
func resolvePollVoteOptions(options []string, selectedHashes [][]byte) []string {
	knownHashes := whatsmeow.HashPollOptions(options)

	selected := make([]string, 0, len(selectedHashes))
	for _, hash := range selectedHashes {
		matched := ""
		for i, known := range knownHashes {
			if bytes.Equal(hash, known) {
				matched = options[i]
				break
			}
		}
		if matched == "" {
			limit := len(hash)
			if limit > 8 {
				limit = 8
			}
			matched = fmt.Sprintf("<unknown:%x>", hash[:limit])
		}
		selected = append(selected, matched)
	}
	return selected
}

// ---------- send ----------

// sendWhatsAppPoll mirrors sendWhatsAppMessage's JID handling: resolve the
// recipient for sending, but persist under the pre-LID-resolution JID so the
// poll lands in the same chat row as everything else.
func sendWhatsAppPoll(client *whatsmeow.Client, messageStore *MessageStore, req SendPollRequest) (bool, string) {
	if !client.IsConnected() {
		return false, "Not connected to WhatsApp"
	}

	question := strings.TrimSpace(req.Question)
	options := normalizePollOptions(req.Options)
	selectable := req.selectableCount()

	var settingsLookupJID types.JID
	var err error
	if strings.Contains(req.Recipient, "@") {
		settingsLookupJID, err = types.ParseJID(req.Recipient)
		if err != nil {
			return false, fmt.Sprintf("Error parsing JID: %v", err)
		}
	} else {
		settingsLookupJID = types.JID{
			User:   req.Recipient,
			Server: "s.whatsapp.net",
		}
	}
	storageJID := settingsLookupJID

	recipientJID, err := resolveRecipientJID(client, req.Recipient)
	if err != nil {
		return false, err.Error()
	}

	// BuildPollCreation attaches the MessageSecret that whatsmeow later needs
	// to decrypt votes; SendMessage persists it to the msgsecret store.
	msg := client.BuildPollCreation(question, options, selectable)

	settings, err := messageStore.GetChatEphemeralSettings(resolveUserJID(client, settingsLookupJID, types.EmptyJID).String())
	if err != nil && err != sql.ErrNoRows {
		return false, fmt.Sprintf("Error loading chat settings: %v", err)
	}
	if err == nil {
		applyChatEphemeralSettings(msg, settings)
	}

	resp, err := client.SendMessage(context.Background(), recipientJID, msg)
	if err != nil {
		return false, fmt.Sprintf("Error sending poll: %v", err)
	}

	// whatsmeow does not re-emit events.Message for our own sends, so persist
	// explicitly — same reasoning as sendWhatsAppMessage.
	if messageStore != nil && client.Store != nil && client.Store.ID != nil {
		persistJID := resolveUserJID(client, storageJID, types.EmptyJID)
		chatJID := persistJID.String()
		timestamp := resp.Timestamp
		if timestamp.IsZero() {
			timestamp = time.Now()
		}

		if chatErr := messageStore.StoreChat(chatJID, "", timestamp); chatErr != nil {
			fmt.Printf("Warning: failed to store outbound chat metadata: %v\n", chatErr)
		}
		if pollErr := messageStore.StorePoll(resp.ID, chatJID, question, options, selectable, timestamp, true); pollErr != nil {
			fmt.Printf("Warning: failed to persist outbound poll: %v\n", pollErr)
		}
		if storeErr := messageStore.StoreMessage(
			resp.ID, chatJID, client.Store.ID.User, formatPollContent(question, options), timestamp, true,
			"", "", "", nil, nil, nil, 0, "",
		); storeErr != nil {
			fmt.Printf("Warning: failed to persist outbound poll message: %v\n", storeErr)
		}
	}

	return true, fmt.Sprintf("Poll sent to %s", req.Recipient)
}

// ---------- receive ----------

// handlePollCreation records an observed poll's options. Without this, votes
// on polls we did not send would decrypt to unresolvable hashes.
func handlePollCreation(messageStore *MessageStore, msg *events.Message, chatJID string, logger waLog.Logger) {
	poll := msg.Message.GetPollCreationMessage()
	if poll == nil {
		return
	}

	if err := messageStore.StorePoll(
		msg.Info.ID, chatJID, poll.GetName(), pollOptionNames(poll),
		int(poll.GetSelectableOptionsCount()), msg.Info.Timestamp, msg.Info.IsFromMe,
	); err != nil {
		logger.Warnf("Failed to store poll %s: %v", msg.Info.ID, err)
	}
}

// handlePollUpdate decrypts an incoming vote and persists it both structurally
// (poll_votes) and readably (messages). Returns true when the event was a poll
// vote, so the caller can stop treating it as a regular message.
func handlePollUpdate(client *whatsmeow.Client, messageStore *MessageStore, msg *events.Message, chatJID, sender string, logger waLog.Logger) bool {
	pollUpdate := msg.Message.GetPollUpdateMessage()
	if pollUpdate == nil {
		return false
	}

	pollID := pollUpdate.GetPollCreationMessageKey().GetID()

	vote, err := client.DecryptPollVote(context.Background(), msg)
	if err != nil {
		// Most common cause: the poll was created before this device paired,
		// so its message secret was never stored. Nothing to recover.
		logger.Warnf("Failed to decrypt poll vote %s (poll %s): %v", msg.Info.ID, pollID, err)
		return true
	}

	question, options, err := messageStore.GetPoll(pollID, chatJID)
	if err != nil && err != sql.ErrNoRows {
		logger.Warnf("Failed to load poll %s for vote %s: %v", pollID, msg.Info.ID, err)
	}

	selected := resolvePollVoteOptions(options, vote.GetSelectedOptions())

	if err := messageStore.StorePollVote(
		msg.Info.ID, pollID, chatJID, sender, selected, msg.Info.Timestamp,
	); err != nil {
		logger.Warnf("Failed to store poll vote %s: %v", msg.Info.ID, err)
	}

	if err := messageStore.StoreMessage(
		msg.Info.ID, chatJID, sender, formatPollVoteContent(question, selected),
		msg.Info.Timestamp, msg.Info.IsFromMe, "", "", "", nil, nil, nil, 0, "",
	); err != nil {
		logger.Warnf("Failed to persist poll vote message %s: %v", msg.Info.ID, err)
	}

	logger.Infof("Poll vote from %s on poll %s: %s", sender, pollID, strings.Join(selected, ", "))
	return true
}
