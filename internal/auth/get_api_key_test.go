package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectError bool
		expectedErr error
	}{
		{
			name: "valid api key",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-key"},
			},
			expectedKey: "my-secret-key",
			expectError: false,
		},
		{
			name:        "missing authorization header",
			headers:     http.Header{},
			expectedKey: "",
			expectError: true,
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed authorization header",
			headers: http.Header{
				"Authorization": []string{"Bearer token"},
			},
			expectedKey: "",
			expectError: true,
		},
		{
			name: "empty api key",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey: "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			t.Logf("TEST CASE: %s", tt.name)
			t.Logf("INPUT: Authorization=%v", tt.headers.Get("Authorization"))

			apiKey, err := GetAPIKey(tt.headers)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}

				if tt.expectedErr != nil && err != tt.expectedErr {
					t.Fatalf("expected error %v but got %v", tt.expectedErr, err)
				}

				return
			}

			t.Logf("OUTPUT: apiKey=%s", apiKey)

			if err != nil {
				t.Fatalf("did not expect error but got %v", err)
			}

			if apiKey != tt.expectedKey {
				t.Fatalf("expected api key %q but got %q", tt.expectedKey, apiKey)
			}
			t.Logf("SUCCESS ✔")
		})
	}
}
