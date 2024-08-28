package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsInputSpaceFree(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Single Word",
			input:    "username",
			expected: true,
		},
		{
			name:     "Multiple Words",
			input:    "user name",
			expected: false,
		},
		{
			name:     "Empty String",
			input:    "",
			expected: true,
		},
		{
			name:     "Leading and Trailing Spaces",
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
