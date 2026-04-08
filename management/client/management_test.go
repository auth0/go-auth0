package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBaseURL(t *testing.T) {
	tests := []struct {
		domain      string
		expectedURL string
	}{
		{"example.com", "https://example.com/api/v2"},
		{"example.com/", "https://example.com/api/v2"},
		{"https://example.com", "https://example.com/api/v2"},
		{"https://example.com/", "https://example.com/api/v2"},
	}

	for _, tt := range tests {
		t.Run(tt.domain, func(t *testing.T) {
			m, err := New(tt.domain)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedURL, m.baseURL)
		})
	}
}
