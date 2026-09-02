package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"google.golang.org/protobuf/proto"
)

func TestValidatePollRequest(t *testing.T) {
	valid := SendPollRequest{
		Recipient: "15551234567@s.whatsapp.net",
		Question:  "Deploy today?",
		Options:   []string{"Yes", "No"},
	}

	tests := []struct {
		name    string
		mutate  func(*SendPollRequest)
		wantErr string
	}{
		{name: "valid single choice"},
		{
			name:   "valid multi choice",
			mutate: func(r *SendPollRequest) { r.AllowMultiple = true },
		},
		{
			name:    "missing recipient",
			mutate:  func(r *SendPollRequest) { r.Recipient = "  " },
			wantErr: "recipient is required",
		},
		{
			name:    "empty question",
			mutate:  func(r *SendPollRequest) { r.Question = "   " },
			wantErr: "question is required",
		},
		{
			name:    "question too long",
			mutate:  func(r *SendPollRequest) { r.Question = strings.Repeat("a", maxPollQuestionLen+1) },
			wantErr: "question is 256 characters",
		},
		{
			name:    "too few options",
			mutate:  func(r *SendPollRequest) { r.Options = []string{"Only one"} },
			wantErr: "between 2 and 12 options",
		},
		{
			name: "too many options",
			mutate: func(r *SendPollRequest) {
				opts := make([]string, maxPollOptions+1)
				for i := range opts {
					opts[i] = string(rune('a' + i))
				}
				r.Options = opts
			},
			wantErr: "between 2 and 12 options",
		},
		{
			name:    "empty option",
			mutate:  func(r *SendPollRequest) { r.Options = []string{"Yes", "   "} },
			wantErr: "option 2 is empty",
		},
		{
			name: "option too long",
			mutate: func(r *SendPollRequest) {
				r.Options = []string{"Yes", strings.Repeat("b", maxPollOptionLen+1)}
			},
			wantErr: "option 2 is 101 characters",
		},
		{
			name:    "duplicate options",
			mutate:  func(r *SendPollRequest) { r.Options = []string{"Yes", "No", "Yes"} },
			wantErr: "duplicates option 1",
		},
		{
			name:    "duplicate options after trimming",
			mutate:  func(r *SendPollRequest) { r.Options = []string{"Yes", "  Yes  "} },
			wantErr: "duplicates option 1",
		},
		{
			name:    "selectable count above option count",
			mutate:  func(r *SendPollRequest) { r.SelectableOptionsCount = 5 },
			wantErr: "must be between 1 and 2",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := valid
			req.Options = append([]string(nil), valid.Options...)
			if tc.mutate != nil {
				tc.mutate(&req)
			}

			err := validatePollRequest(req)
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("expected request to validate, got %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("expected error containing %q, got nil", tc.wantErr)
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("expected error containing %q, got %q", tc.wantErr, err.Error())
			}
		})
	}
}

// Multi-byte questions must be measured in runes, not bytes, or an accented
// Portuguese question would be rejected well before WhatsApp's actual limit.
func TestValidatePollRequestCountsRunesNotBytes(t *testing.T) {
	req := SendPollRequest{
		Recipient: "15551234567@s.whatsapp.net",
		Question:  strings.Repeat("ç", maxPollQuestionLen),
		Options:   []string{"Sim", "Não"},
	}
	if err := validatePollRequest(req); err != nil {
		t.Fatalf("expected 255-rune question to validate, got %v", err)
	}
}

func TestSelectableCount(t *testing.T) {
	options := []string{"A", "B", "C"}

	tests := []struct {
		name string
		req  SendPollRequest
		want int
	}{
		{
			name: "defaults to single choice",
			req:  SendPollRequest{Options: options},
			want: 1,
		},
		{
			name: "allow multiple selects every option",
			req:  SendPollRequest{Options: options, AllowMultiple: true},
			want: 3,
		},
		{
			name: "explicit count wins over allow multiple",
			req:  SendPollRequest{Options: options, AllowMultiple: true, SelectableOptionsCount: 2},
			want: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.req.selectableCount(); got != tc.want {
				t.Fatalf("expected selectable count %d, got %d", tc.want, got)
			}
		})
	}
}

func TestResolvePollVoteOptions(t *testing.T) {
	options := []string{"Yes", "No", "Maybe"}
	hashes := whatsmeow.HashPollOptions(options)

	t.Run("resolves known hashes back to text", func(t *testing.T) {
		selected := resolvePollVoteOptions(options, [][]byte{hashes[2], hashes[0]})
		want := []string{"Maybe", "Yes"}
		if len(selected) != len(want) {
			t.Fatalf("expected %d selections, got %v", len(want), selected)
		}
		for i := range want {
			if selected[i] != want[i] {
				t.Fatalf("expected selection %d to be %q, got %q", i, want[i], selected[i])
			}
		}
	})

	t.Run("unknown hash is preserved as a prefix", func(t *testing.T) {
		unknown := whatsmeow.HashPollOptions([]string{"Removed option"})[0]
		selected := resolvePollVoteOptions(options, [][]byte{unknown})
		if len(selected) != 1 {
			t.Fatalf("expected one selection, got %v", selected)
		}
		if !strings.HasPrefix(selected[0], "<unknown:") {
			t.Fatalf("expected unknown-hash marker, got %q", selected[0])
		}
	})

	t.Run("empty vote yields no selections", func(t *testing.T) {
		if selected := resolvePollVoteOptions(options, nil); len(selected) != 0 {
			t.Fatalf("expected no selections, got %v", selected)
		}
	})

	// A poll we never observed leaves us with no option list; the vote must
	// still round-trip as hashes rather than silently vanish.
	t.Run("no known options still records the vote", func(t *testing.T) {
		selected := resolvePollVoteOptions(nil, [][]byte{hashes[0]})
		if len(selected) != 1 || !strings.HasPrefix(selected[0], "<unknown:") {
			t.Fatalf("expected unknown-hash marker, got %v", selected)
		}
	})
}

func TestFormatPollContent(t *testing.T) {
	got := formatPollContent("Deploy today?", []string{"Yes", "No"})
	want := "[poll] Deploy today? | Yes / No"
	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestFormatPollVoteContent(t *testing.T) {
	tests := []struct {
		name     string
		question string
		selected []string
		want     string
	}{
		{
			name:     "single selection",
			question: "Deploy today?",
			selected: []string{"Yes"},
			want:     "[poll-vote] Deploy today? -> Yes",
		},
		{
			name:     "multiple selections",
			question: "Which stacks?",
			selected: []string{"Go", "Python"},
			want:     "[poll-vote] Which stacks? -> Go, Python",
		},
		{
			name:     "cleared vote",
			question: "Deploy today?",
			selected: nil,
			want:     "[poll-vote] Deploy today? -> (vote cleared)",
		},
		{
			name:     "unknown poll",
			question: "",
			selected: []string{"Yes"},
			want:     "[poll-vote] (unknown poll) -> Yes",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := formatPollVoteContent(tc.question, tc.selected); got != tc.want {
				t.Fatalf("expected %q, got %q", tc.want, got)
			}
		})
	}
}

// Polls have no text field. Without explicit rendering they would hit
// StoreMessage's empty-content guard and never reach the database.
func TestExtractTextContentRendersPoll(t *testing.T) {
	msg := &waProto.Message{
		PollCreationMessage: &waE2E.PollCreationMessage{
			Name: proto.String("Deploy today?"),
			Options: []*waE2E.PollCreationMessage_Option{
				{OptionName: proto.String("Yes")},
				{OptionName: proto.String("No")},
			},
			SelectableOptionsCount: proto.Uint32(1),
		},
	}

	got := extractTextContent(msg)
	want := "[poll] Deploy today? | Yes / No"
	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestStorePollRoundTrip(t *testing.T) {
	ms := newTestMessageStore(t)

	chatJID := "15551234567@s.whatsapp.net"
	options := []string{"Yes", "No", "Maybe"}
	if err := ms.StorePoll("POLL1", chatJID, "Deploy today?", options, 1, time.Unix(1710000000, 0), true); err != nil {
		t.Fatalf("failed to store poll: %v", err)
	}

	question, got, err := ms.GetPoll("POLL1", chatJID)
	if err != nil {
		t.Fatalf("failed to read poll back: %v", err)
	}
	if question != "Deploy today?" {
		t.Fatalf("expected question to round-trip, got %q", question)
	}
	if len(got) != len(options) {
		t.Fatalf("expected %d options, got %v", len(options), got)
	}
	for i := range options {
		if got[i] != options[i] {
			t.Fatalf("expected option %d to be %q, got %q", i, options[i], got[i])
		}
	}

	// Re-sending the same poll ID must update rather than fail on the PK.
	if err := ms.StorePoll("POLL1", chatJID, "Deploy tomorrow?", options, 1, time.Unix(1710000100, 0), true); err != nil {
		t.Fatalf("expected upsert to succeed, got %v", err)
	}
	if question, _, err = ms.GetPoll("POLL1", chatJID); err != nil || question != "Deploy tomorrow?" {
		t.Fatalf("expected upserted question, got %q (err %v)", question, err)
	}
}

func TestStorePollVotePersistsSelection(t *testing.T) {
	ms := newTestMessageStore(t)

	chatJID := "120363428713578541@g.us"
	if err := ms.StorePollVote("VOTE1", "POLL1", chatJID, "15551234567", []string{"Yes"}, time.Unix(1710000200, 0)); err != nil {
		t.Fatalf("failed to store vote: %v", err)
	}

	var voter, encoded string
	err := ms.db.QueryRow(
		`SELECT voter, selected_options FROM poll_votes WHERE id = ? AND chat_jid = ?`,
		"VOTE1", chatJID,
	).Scan(&voter, &encoded)
	if err != nil {
		t.Fatalf("failed to read vote back: %v", err)
	}
	if voter != "15551234567" {
		t.Fatalf("expected voter to round-trip, got %q", voter)
	}

	var selected []string
	if err := json.Unmarshal([]byte(encoded), &selected); err != nil {
		t.Fatalf("failed to decode stored selection: %v", err)
	}
	if len(selected) != 1 || selected[0] != "Yes" {
		t.Fatalf("expected [Yes], got %v", selected)
	}
}

func TestPollHandlerRejectsInvalidRequest(t *testing.T) {
	const token = "supersecrettoken1234567890abcdef"

	handler := newRESTMux(newTestClient(&mockLIDStore{}), newTestMessageStore(t), 8080, token, nil)

	body := `{"recipient":"15551234567@s.whatsapp.net","question":"Deploy?","options":["Yes"]}`
	req := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8080/api/poll", strings.NewReader(body))
	req.RemoteAddr = "127.0.0.1:54321"
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for a one-option poll, got %d", resp.Code)
	}

	var decoded SendMessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if decoded.Success {
		t.Fatal("expected success=false for an invalid poll")
	}
	if !strings.Contains(decoded.Message, "between 2 and 12 options") {
		t.Fatalf("expected validation reason in response, got %q", decoded.Message)
	}
}

func TestPollHandlerRejectsNonPost(t *testing.T) {
	const token = "supersecrettoken1234567890abcdef"

	handler := newRESTMux(newTestClient(&mockLIDStore{}), newTestMessageStore(t), 8080, token, nil)
	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8080/api/poll", nil)
	req.RemoteAddr = "127.0.0.1:54321"
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	if resp.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405 for GET, got %d", resp.Code)
	}
}

// The poll endpoint must sit behind the same bearer-token wall as /api/send.
func TestPollHandlerRequiresAuth(t *testing.T) {
	const token = "supersecrettoken1234567890abcdef"

	handler := newRESTMux(newTestClient(&mockLIDStore{}), newTestMessageStore(t), 8080, token, nil)
	body := `{"recipient":"15551234567@s.whatsapp.net","question":"Deploy?","options":["Yes","No"]}`
	req := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8080/api/poll", strings.NewReader(body))
	req.RemoteAddr = "127.0.0.1:54321"
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	if resp.Code == http.StatusOK {
		t.Fatal("expected unauthenticated poll request to be rejected")
	}
}
