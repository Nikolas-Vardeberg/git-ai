package groq

import (
	"encoding/json"
	"gitai/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePrDescriptionWithGroq(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/pr-description" {
			t.Errorf("expected path /api/pr-description, got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}

		var req CommitRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		if req.GitDiff != "test diff" {
			t.Errorf("expected gitDiff 'test diff', got '%s'", req.GitDiff)
		}

		resp := prDescriptionResp{
			Data: &prDescriptionData{
				PrDescription: "Generated PR Description",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	oldBaseURL := server.ServerConfig.BaseURL
	server.ServerConfig.BaseURL = ts.URL
	defer func() {
		server.ServerConfig.BaseURL = oldBaseURL
	}()

	desc, err := CreatePrDescriptionWithGroq("test diff")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if desc != "Generated PR Description" {
		t.Errorf("expected 'Generated PR Description', got '%s'", desc)
	}
}

func TestCreatePrDescriptionWithGroq_FallbackFields(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := prDescriptionResp{
			Data: &prDescriptionData{
				Description: "Fallback Description",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	oldBaseURL := server.ServerConfig.BaseURL
	server.ServerConfig.BaseURL = ts.URL
	defer func() {
		server.ServerConfig.BaseURL = oldBaseURL
	}()

	desc, err := CreatePrDescriptionWithGroq("test diff")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if desc != "Fallback Description" {
		t.Errorf("expected 'Fallback Description', got '%s'", desc)
	}
}

func TestCreatePrDescriptionWithGroq_ErrorFromServer(t *testing.T) {
	errMsg := "rate limit exceeded"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := prDescriptionResp{
			Error: &errMsg,
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	oldBaseURL := server.ServerConfig.BaseURL
	server.ServerConfig.BaseURL = ts.URL
	defer func() {
		server.ServerConfig.BaseURL = oldBaseURL
	}()

	_, err := CreatePrDescriptionWithGroq("test diff")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err.Error() != errMsg {
		t.Errorf("expected error '%s', got '%s'", errMsg, err.Error())
	}
}
