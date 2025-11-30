package demon

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"testing"

	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/models"
	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/repo/mock"
)

func TestCreateSession(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	mockRepo := &mock.MockSessionRepo{
		CreateSessionFunc: func(ctx context.Context, session models.Session) (int64, error) {
			if session.Label != "test-session" {
				t.Errorf("expected label 'test-session', got %s", session.Label)
			}
			return 123, nil
		},
	}

	payload := createSessionReq{
		Label:    "test-session",
		Note:     "test note",
		Tracked:  true,
		Estimate: 60,
	}
	data, _ := json.Marshal(payload)
	req := Request{
		Method: "START",
		Data:   data,
	}

	go createSession(server, mockRepo, req)

	var resp envelope
	decoder := json.NewDecoder(client)
	if err := decoder.Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp["status"] != "success" {
		t.Errorf("expected status success, got %v", resp["status"])
	}

	sessionMap, ok := resp["session"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected session object in response, got %T", resp["session"])
	}

	if id, ok := sessionMap["id"].(float64); !ok || int64(id) != 123 {
		t.Errorf("expected session id 123, got %v", sessionMap["id"])
	}

}
