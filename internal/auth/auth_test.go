package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError string // Changed to string for error message comparison
	}{
		{
			name:          "Valid Authorization Header",
			headers:       http.Header{"Authorization": []string{"ApiKey abc123"}},
			expectedKey:   "abc123",
			expectedError: "",
		},
		{
			name:          "No Authorization Header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name:          "Malformed Authorization Header",
			headers:       http.Header{"Authorization": []string{"Bearer abc123"}},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
		{
			name:          "Incomplete Authorization Header",
			headers:       http.Header{"Authorization": []string{"ApiKey"}},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
		{
			name:          "Empty Authorization Value",
			headers:       http.Header{"Authorization": []string{""}},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded.Error(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)

			if key != tc.expectedKey {
				t.Errorf("expected key %v, got %v", tc.expectedKey, key)
			}

			if err != nil {
				if err.Error() != tc.expectedError {
					t.Errorf("expected error %v, got %v", tc.expectedError, err.Error())
				}
			} else if tc.expectedError != "" {
				t.Errorf("expected error %v, got nil", tc.expectedError)
			}
		})
	}
}
