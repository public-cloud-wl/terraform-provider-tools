package provider

import "testing"

func TestHashHexPrefix(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		length   int
		expected string
	}{
		{
			name:     "project code",
			input:    "nm-toto-01-production",
			length:   16,
			expected: "9cfa3a3ba2a87e14",
		},
		{
			name:     "role",
			input:    "super-role",
			length:   8,
			expected: "c98e1693",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := hashHexPrefix(tc.input, tc.length)
			if got != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}
