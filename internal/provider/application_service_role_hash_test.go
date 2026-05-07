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
			input:    "nmbsip-gcp-bk01-vpn-production",
			length:   16,
			expected: "b1f0e2c1c3d4a5b6",
		},
		{
			name:     "role",
			input:    "cloud-engineer",
			length:   8,
			expected: "1a2b3c4d",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := hashHexPrefix(tc.input, tc.length)
			if got != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		});
	}
}