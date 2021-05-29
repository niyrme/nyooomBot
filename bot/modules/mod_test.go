package modules_test

import (
	. "nyooomBot/bot/modules"
	"testing"
)

func TestContains(t *testing.T) {
	// TODO: Add more tests
	// TODO: Add tests for every module
	tests := []struct {
		inputS   []string
		inputStr string
		expected bool
	}{
		{ // Test contains
			inputS:   []string{"a", "b", "c"},
			inputStr: "a",
			expected: true,
		},
		{ // Test empty
			inputS:   []string{},
			inputStr: "test",
			expected: false,
		},
		{ // Test single item in slice
			inputS:   []string{"giraffe"},
			inputStr: "giraffe",
			expected: true,
		},
		{ // Test capitalization
			inputS:   []string{"GiraFFE"},
			inputStr: "giraffe",
			expected: false,
		},
	}

	for _, test := range tests {
		if output := Contains(test.inputS, test.inputStr); output != test.expected {
			t.Errorf("Test failed. want: %v; got: %v", test.expected, output)
		}
	}
}
