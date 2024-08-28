package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "Valid email address",
			email:    "example@example.com",
			expected: true,
		},
		{
			name:     "Valid email with subdomain",
			email:    "user@mail.example.com",
			expected: true,
		},
		{
			name:     "Valid email with special characters",
			email:    "user+name@example.co.uk",
			expected: true,
		},
		{
			name:     "Invalid email without @",
			email:    "example.com",
			expected: false,
		},
		{
			name:     "Invalid email without domain",
			email:    "user@.com",
			expected: false,
		},
		{
			name:     "Invalid email with multiple @",
			email:    "user@@example.com",
			expected: false,
		},
		{
			name:     "Invalid email with spaces",
			email:    "user name@example.com",
			expected: false,
		},
		{
			name:     "Invalid email with invalid characters",
			email:    "user@exa*mple.com",
			expected: false,
		},
		{
			name:     "Invalid email with missing domain extension",
			email:    "user@example",
			expected: false,
		},
		{
			name:     "Empty email address",
			email:    "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidEmail(tt.email)
			assert.Equal(t, tt.expected, result)
		})
	}
}
