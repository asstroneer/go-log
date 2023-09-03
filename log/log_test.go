package log

import "testing"

func TestCompareLogLevel(t *testing.T) {
	type test struct {
		input       int
		systemLevel string
		expected    bool
	}

	tests := []test{
		{
			input:       2,
			systemLevel: "info",
			expected:    true,
		},
		{
			input:       3,
			systemLevel: "error",
			expected:    false,
		},
		{
			input:       1,
			systemLevel: "error",
			expected:    true,
		},
	}

	for _, tt := range tests {
		SetLogLevel(tt.systemLevel)
		actual := compareLogLevel(tt.input)
		if actual != tt.expected {
			t.Errorf("compareLogLevel(%q, %q) expected: %t, actual: %t", tt.input, tt.systemLevel, tt.expected, actual)
		}
	}
}

func TestLevelToCode(t *testing.T) {
	type test struct {
		input    string
		expected int
	}

	tests := []test{
		{
			input:    "fatal",
			expected: 0,
		},
		{
			input:    "info",
			expected: 3,
		},
		{
			input:    "error",
			expected: 1,
		},
		{
			input:    "warning",
			expected: 2,
		},
		{
			input:    "debug",
			expected: 4,
		},
		{
			input:    "unknown",
			expected: 100,
		},
	}

	for _, tt := range tests {
		actual := logLevelToCode(tt.input)
		if actual != tt.expected {
			t.Errorf("levelToCode(%q) expected: %d, actual: %d", tt.input, tt.expected, actual)
		}
	}
}
