package validation

import (
	"testing"
)

func TestIsInputSpaceFree(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Input without spaces",
			input:    "NoSpacesHere",
			expected: true,
		},
		{
			name:     "Input with a space",
			input:    "Has Space",
			expected: false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "Input with leading space",
			input:    " LeadingSpace",
			expected: false,
		},
		{
			name:     "Input with trailing space",
			input:    "TrailingSpace ",
			expected: false,
		},
		{
			name:     "Input with multiple spaces",
			input:    "Multiple  Spaces",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsInputSpaceFree(tt.input)
			if result != tt.expected {
				t.Errorf("IsInputSpaceFree(%q) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
