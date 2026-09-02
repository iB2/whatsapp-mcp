package main

import (
	"testing"
	"time"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/proto/waCommon"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

// buildReactionMessage mirrors buildTextMessage, but the payload is a
// ReactionMessage pointing at targetID. An empty emoji is a retraction.
func buildReactionMessage(chat, sender, senderAlt types.JID, isFromMe bool, id, targetID, emoji string, targetFromMe bool, ts time.Time) *events.Message {
	return &events.Message{
		Info: types.MessageInfo{
			MessageSource: types.MessageSource{
				Chat:      chat,
				Sender:    sender,
				SenderAlt: senderAlt,
				IsFromMe:  isFromMe,
				IsGroup:   false,
			},
			ID:        id,
			Timestamp: ts,
		},
		Message: &waProto.Message{
			ReactionMessage: &waProto.ReactionMessage{
				Key: &waCommon.MessageKey{
					ID:        proto.String(targetID),
					FromMe:    proto.Bool(targetFromMe),
					RemoteJID: proto.String(chat.String()),
				},
				Text: proto.String(emoji),
			},
		},
	}
}

func queryReaction(t *testing.T, ms *MessageStore, chatJID, targetID, sender string) (emoji string, targetFromMe bool, found bool) {
	t.Helper()
	err := ms.db.QueryRow(
		"SELECT emoji, target_from_me FROM reactions WHERE chat_jid = ? AND target_id = ? AND sender = ?",
		chatJID, targetID, sender,
	).Scan(&emoji, &targetFromMe)
	return emoji, targetFromMe, err == nil
}

func queryReactionCount(t *testing.T, ms *MessageStore) int {
	t.Helper()
	var n int
	if err := ms.db.QueryRow("SELECT COUNT(*) FROM reactions").Scan(&n); err != nil {
		t.Fatalf("counting reactions: %v", err)
	}
	return n
}

// A reaction must land in the reactions table and must NOT become chat
// content — that separation is the whole design (see reactions.go).
func TestHandleMessage_ReactionStoredOutsideMessages(t *testing.T) {
	client := newTestClient(&mockLIDStore{})
	ms := newTestMessageStore(t)
	logger := testLogger()

	msg := buildReactionMessage(phonePN, phonePN, types.EmptyJID, false,
		"reaction-001", "episode-msg-001", "\U0001F44D", true, time.Now())

	handleMessage(client, ms, msg, logger)

	emoji, targetFromMe, found := queryReaction(t, ms, phonePN.String(), "episode-msg-001", phonePN.User)
	if !found {
		t.Fatal("reaction was not stored")
	}
	if emoji != "\U0001F44D" {
		t.Errorf("emoji = %q, want thumbs up", emoji)
	}
	if !targetFromMe {
		t.Error("target_from_me should be true: the reaction is on a message we sent")
	}

	// The regression this guards: reactions leaking into chat history would
	// flood every existing reader of messages.db.
	if count := queryMessageCount(ms, phonePN.String()); count != 0 {
		t.Errorf("expected 0 messages stored for a reaction, got %d", count)
	}
}

// Reacting again replaces the emoji in place: one row per (target, sender).
func TestHandleMessage_ReactionChangeReplacesInPlace(t *testing.T) {
	client := newTestClient(&mockLIDStore{})
	ms := newTestMessageStore(t)
	logger := testLogger()

	base := time.Now()
	handleMessage(client, ms, buildReactionMessage(phonePN, phonePN, types.EmptyJID, false,
		"reaction-001", "episode-msg-001", "\U0001F44D", true, base), logger)
	handleMessage(client, ms, buildReactionMessage(phonePN, phonePN, types.EmptyJID, false,
		"reaction-002", "episode-msg-001", "\U0001F4A9", true, base.Add(time.Minute)), logger)

	if n := queryReactionCount(t, ms); n != 1 {
		t.Fatalf("expected 1 reaction row after a change, got %d", n)
	}
	emoji, _, _ := queryReaction(t, ms, phonePN.String(), "episode-msg-001", phonePN.User)
	if emoji != "\U0001F4A9" {
		t.Errorf("emoji = %q, want the newer reaction", emoji)
	}
}

// Taking a reaction back stores an empty emoji rather than deleting the row,
// so a reader can distinguish "never reacted" from "un-reacted".
func TestHandleMessage_ReactionRemovalKeepsRowWithEmptyEmoji(t *testing.T) {
	client := newTestClient(&mockLIDStore{})
	ms := newTestMessageStore(t)
	logger := testLogger()

	base := time.Now()
	handleMessage(client, ms, buildReactionMessage(phonePN, phonePN, types.EmptyJID, false,
		"reaction-001", "episode-msg-001", "\U0001F44D", true, base), logger)
	handleMessage(client, ms, buildReactionMessage(phonePN, phonePN, types.EmptyJID, false,
		"reaction-002", "episode-msg-001", "", true, base.Add(time.Minute)), logger)

	emoji, _, found := queryReaction(t, ms, phonePN.String(), "episode-msg-001", phonePN.User)
	if !found {
		t.Fatal("row should survive a retraction")
	}
	if emoji != "" {
		t.Errorf("emoji = %q, want empty after retraction", emoji)
	}
}

// Delivery order is not guaranteed. A retraction that arrives late but carries
// an older timestamp must not overwrite the newer reaction.
func TestStoreReaction_OlderTimestampDoesNotWin(t *testing.T) {
	ms := newTestMessageStore(t)
	base := time.Now()

	if err := ms.StoreReaction("r2", "chat@s.whatsapp.net", "target-1", "5521999999999", "\U0001F44D", base.Add(time.Minute), false, true); err != nil {
		t.Fatalf("storing newer reaction: %v", err)
	}
	if err := ms.StoreReaction("r1", "chat@s.whatsapp.net", "target-1", "5521999999999", "", base, false, true); err != nil {
		t.Fatalf("storing older retraction: %v", err)
	}

	emoji, found := ms.GetReaction("chat@s.whatsapp.net", "target-1", "5521999999999")
	if !found {
		t.Fatal("reaction missing")
	}
	if emoji != "\U0001F44D" {
		t.Errorf("emoji = %q, want the newer reaction to survive the out-of-order retraction", emoji)
	}
}

// Two people reacting to the same message are two rows, not a fight over one.
func TestStoreReaction_DistinctSendersCoexist(t *testing.T) {
	ms := newTestMessageStore(t)
	now := time.Now()

	if err := ms.StoreReaction("r1", "g@g.us", "target-1", "5521111111111", "\U0001F44D", now, false, true); err != nil {
		t.Fatalf("first sender: %v", err)
	}
	if err := ms.StoreReaction("r2", "g@g.us", "target-1", "5521222222222", "❤️", now, false, true); err != nil {
		t.Fatalf("second sender: %v", err)
	}

	if n := queryReactionCount(t, ms); n != 2 {
		t.Errorf("expected 2 rows for 2 senders, got %d", n)
	}
}

func TestStoreReaction_RejectsMissingTarget(t *testing.T) {
	ms := newTestMessageStore(t)
	if err := ms.StoreReaction("r1", "g@g.us", "", "5521111111111", "\U0001F44D", time.Now(), false, true); err == nil {
		t.Error("a reaction with no target message id must be rejected")
	}
}

// GetReaction must not invent a reaction for a message nobody touched.
func TestGetReaction_MissingIsNotFound(t *testing.T) {
	ms := newTestMessageStore(t)
	if _, found := ms.GetReaction("g@g.us", "never-reacted", "5521111111111"); found {
		t.Error("expected not found for a message with no reaction")
	}
}
