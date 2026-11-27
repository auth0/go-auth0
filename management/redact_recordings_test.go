package management

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedactSensitiveJSONBody(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		expected string
	}{
		{
			name:     "redacts client secrets in a list response",
			body:     `{"total":1,"clients":[{"client_id":"abc","client_secret":"super-secret-value"}]}`,
			expected: `{"clients":[{"client_id":"abc","client_secret":"[REDACTED]"}],"total":1}`,
		},
		{
			name:     "redacts recovery codes and signing secrets",
			body:     `{"recovery_code":"ABC","recovery_codes":["A","B"],"signing_secret":"xyz"}`,
			expected: `{"recovery_code":"[REDACTED]","recovery_codes":["[REDACTED]","[REDACTED]"],"signing_secret":"[REDACTED]"}`,
		},
		{
			name:     "redacts secrets in a barcode uri",
			body:     `{"barcode_uri":"otpauth://totp/Tenant:user?secret=ABC&issuer=Tenant"}`,
			expected: `{"barcode_uri":"otpauth://totp/Tenant:user?secret=[REDACTED]&issuer=Tenant"}`,
		},
		{
			name:     "keeps connection options untouched",
			body:     `{"options":{"client_secret":"fixture"}}`,
			expected: `{"options":{"client_secret":"fixture"}}`,
		},
		{
			name:     "returns bodies without sensitive fields unchanged",
			body:     `{"id":12345678901234567890, "name":"test"}`,
			expected: `{"id":12345678901234567890, "name":"test"}`,
		},
		{
			name:     "returns non JSON bodies unchanged",
			body:     "not json",
			expected: "not json",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, redactSensitiveJSONBody(t, testCase.body))
		})
	}
}
