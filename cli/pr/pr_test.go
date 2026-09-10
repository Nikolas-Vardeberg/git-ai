package pr

import (
	"encoding/json"
	"gitai/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePRDescription(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"data": map[string]string{
				"prDescription": "## Pull Request\n\nImplemented feature X.",
			},
		})
	}))
	defer ts.Close()

	oldBaseURL := server.ServerConfig.BaseURL
	server.ServerConfig.BaseURL = ts.URL
	defer func() {
		server.ServerConfig.BaseURL = oldBaseURL
	}()

	desc, err := CreatePRDescription("diff --git a/file b/file")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "## Pull Request\n\nImplemented feature X."
	if desc != expected {
		t.Errorf("expected %q, got %q", expected, desc)
	}
}
