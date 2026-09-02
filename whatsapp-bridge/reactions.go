package main

import (
	"fmt"
	"strings"
	"time"

	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

// Message reactions.
//
// A reaction arrives as an ordinary events.Message whose payload is a
// ReactionMessage: no text, no media. Before this file existed the bridge
// dropped every one of them on the floor — extractTextContent has no branch
// for ReactionMessage, so content came back empty, and handleMessage bailed at
// the "no content and no media" guard. The events were always there (hundreds
// a day in the receive log); nothing was ever persisted.
//
// Reactions are deliberately NOT mirrored into the messages table, which is
// where poll votes go. Volume is the reason: a busy account sees hundreds of
// reactions a day, and every existing reader of messages.db — list_messages,
// the MCP tools, the digest scripts — would suddenly have to filter them out
// of chat history. Reactions live in their own table and get read on purpose.
//
// State, not history. WhatsApp allows one reaction per (message, sender):
// changing it or taking it back sends another ReactionMessage. So the row is
// keyed by target message plus sender and updated in place. An empty emoji is
// how the protocol says "reaction removed", and it is stored as such instead
// of deleting the row, so a reader can tell "never reacted" from "un-reacted".

// StoreReaction records the current reaction state for (chat, target message,
// sender).
//
// The timestamp guard exists because delivery order is not guaranteed: a
// retraction that arrives before the reaction it retracts must not win. It
// compares the driver's own timestamp encoding, which is lexicographically
// ordered only while the UTC offset is stable — good enough for ordering two
// reactions seconds apart, and the failure mode is a stale emoji, never a lost
// row.
func (store *MessageStore) StoreReaction(
	id, chatJID, targetID, sender, emoji string,
	timestamp time.Time,
	isFromMe, targetFromMe bool,
) error {
	if targetID == "" {
		return fmt.Errorf("reaction %s carries no target message id", id)
	}

	_, err := store.db.Exec(
		`INSERT INTO reactions
		(id, chat_jid, target_id, target_from_me, sender, emoji, timestamp, is_from_me)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(chat_jid, target_id, sender) DO UPDATE SET
			id = excluded.id,
			emoji = excluded.emoji,
			timestamp = excluded.timestamp,
			is_from_me = excluded.is_from_me,
			target_from_me = excluded.target_from_me
		WHERE excluded.timestamp >= reactions.timestamp`,
		id, chatJID, targetID, targetFromMe, sender, emoji, timestamp, isFromMe,
	)
	return err
}

// GetReaction returns the current emoji a sender has on a target message.
// An empty emoji with found=true means the reaction was taken back.
func (store *MessageStore) GetReaction(chatJID, targetID, sender string) (emoji string, found bool) {
	err := store.db.QueryRow(
		`SELECT emoji FROM reactions WHERE chat_jid = ? AND target_id = ? AND sender = ?`,
		chatJID, targetID, sender,
	).Scan(&emoji)
	return emoji, err == nil
}

// handleReaction persists an incoming reaction and reports whether the event
// was one, so handleMessage can stop treating it as a regular message.
func handleReaction(messageStore *MessageStore, msg *events.Message, chatJID, sender string, logger waLog.Logger) bool {
	reaction := msg.Message.GetReactionMessage()
	if reaction == nil {
		return false
	}

	key := reaction.GetKey()
	targetID := key.GetID()
	if targetID == "" {
		// Nothing actionable: a reaction with no target cannot be attributed.
		// Swallow it anyway — it is still a reaction, not chat content.
		logger.Warnf("Reaction %s carries no target message key; dropping", msg.Info.ID)
		return true
	}

	emoji := strings.TrimSpace(reaction.GetText())

	if err := messageStore.StoreReaction(
		msg.Info.ID, chatJID, targetID, sender, emoji,
		msg.Info.Timestamp, msg.Info.IsFromMe, key.GetFromMe(),
	); err != nil {
		logger.Warnf("Failed to store reaction %s on message %s: %v", msg.Info.ID, targetID, err)
		return true
	}

	if emoji == "" {
		logger.Infof("Reaction removed by %s on message %s", sender, targetID)
	} else {
		logger.Infof("Reaction %s from %s on message %s", emoji, sender, targetID)
	}
	return true
}
