package main

import (
	"strings"
	"testing"
)

func TestSortLines(t *testing.T) {
	tests := []struct {
		input    string
		opts     sortOptions
		expected string
	}{
		{
			input:    "3 fish\n10 dogs\n2 cats\n1 mouse",
			opts:     sortOptions{column: 0, numeric: true, reverse: false, unique: false},
			expected: "1 mouse\n2 cats\n3 fish\n10 dogs\n",
		},
		{
			input:    "apple 2\nbanana 9\ncherry 12",
			opts:     sortOptions{column: 1, numeric: false, reverse: true, unique: false},
			expected: "cherry 12\napple 2\nbanana 9\n",
		},
		{
			input:    "kiwi\nkiwi\napple\napple",
			opts:     sortOptions{column: 0, numeric: false, reverse: false, unique: true},
			expected: "apple\nkiwi\n",
		},
	}

	for _, test := range tests {
		reader := strings.NewReader(test.input)
		writer := &strings.Builder{}

		lines, _ := readLines(reader)
		sortLines(lines, test.opts)
		writeLines(lines, writer, test.opts)

		result := strings.TrimSpace(writer.String())
		expected := strings.TrimSpace(test.expected)
		if result != expected {
			t.Errorf("Expected sorted lines to be:\n%s\nbut got:\n%s", expected, result)
		}
	}
}
