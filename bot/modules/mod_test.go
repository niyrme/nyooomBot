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

func TestAnswerCommand(t *testing.T) {
	tests := []struct {
		cmd      string
		args     []string
		expected string
	}{
		{ // Test missing argument
			cmd:      "desc",
			args:     []string{},
			expected: ModDesc.How,
		},
		{ // Test unknown command
			cmd:      "desc",
			args:     []string{"doesNotExist"},
			expected: "Unknown command doesNotExist",
		},
		{ // Test no recursion
			cmd:      "help",
			args:     []string{"help"},
			expected: ModHelp.How,
		},
		{ // Test keys
			cmd:      "h",
			args:     []string{"ping"},
			expected: ModPing.How,
		},
		{ // Test multiple args
			cmd:      "help",
			args:     []string{"description", "ping"},
			expected: ModDesc.How,
		},
	}

	for _, test := range tests {
		if output := AnswerCommand(test.cmd, test.args); output != test.expected {
			t.Errorf("Test failed. want: %v; got: %v", test.expected, output)
		}
	}
}
