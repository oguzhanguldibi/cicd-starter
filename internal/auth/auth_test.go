package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("successful API key extraction", func(t *testing.T) {
		headers := http.Header{}
		headers.Add("Authorization", "ApiKey test-key-123")

		key, err := GetAPIKey(headers)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if key != "test-key-123" {
			t.Errorf("expected key 'test-key-123', got '%s'", key)
		}
	})

	t.Run("missing authorization header", func(t *testing.T) {
		headers := http.Header{}

		_, err := GetAPIKey(headers)

		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("expected ErrNoAuthHeaderIncluded, got %v", err)
		}
	})
}
