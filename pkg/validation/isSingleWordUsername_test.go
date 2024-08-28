package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSingleWordUsername(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid single word username",
			input:    "username",
			expected: true,
		},
		{
			name:     "Username with spaces",
			input:    "user name",
			expected: false,
		},
		{
			name:     "Empty username",
			input:    "",
			expected: true,
		},
		{
			name:     "Username with leading and trailing spaces",
			input:    " username ",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSingleWordUsername(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
