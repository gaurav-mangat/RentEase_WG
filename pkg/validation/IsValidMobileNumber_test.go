package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidMobileNumber(t *testing.T) {
	tests := []struct {
		name     string
		number   string
		expected bool
	}{
		{
			name:     "Valid Mobile Number",
			number:   "9876543210",
			expected: true,
		},
		{
			name:     "Valid Mobile Number Starting with 6",
			number:   "6123456789",
			expected: true,
		},
		{
			name:     "Invalid Mobile Number Starting with 5",
			number:   "5123456789",
			expected: false,
		},
		{
			name:     "Invalid Mobile Number Too Short",
			number:   "987654321",
			expected: false,
		},
		{
			name:     "Invalid Mobile Number Too Long",
			number:   "98765432101",
			expected: false,
		},
		{
			name:     "Invalid Mobile Number with Letters",
			number:   "98765abc90",
			expected: false,
		},
		{
			name:     "Empty Mobile Number",
			number:   "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidMobileNumber(tt.number)
			assert.Equal(t, tt.expected, result)
		})
	}
}
