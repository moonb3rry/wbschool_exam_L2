package main

import (
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	input := `Hello
world
Hello again
Grep is great
Hello world
Ignore case
IGNORE CASE
`

	tests := []struct {
		name       string
		searchTerm string
		flags      grepFlags
		expected   string
	}{
		{
			name:       "Count matches",
			searchTerm: "Hello",
			flags:      grepFlags{count: true},
			expected:   "3",
		},
		{
			name:       "Ignore case",
			searchTerm: "ignore case",
			flags:      grepFlags{ignoreCase: true},
			expected:   "Ignore case\nIGNORE CASE\n",
		},
		{
			name:       "Fixed string match",
			searchTerm: "great",
			flags:      grepFlags{fixed: true},
			expected:   "Grep is great\n",
		},
		{
			name:       "Invert match",
			searchTerm: "Hello",
			flags:      grepFlags{invertMatch: true},
			expected:   "world\nGrep is great\nIgnore case\nIGNORE CASE\n",
		},
		{
			name:       "Line number with output",
			searchTerm: "Hello",
			flags:      grepFlags{lineNum: true},
			expected:   "1:Hello\n3:Hello again\n5:Hello world\n",
		},
		{
			name:       "After context",
			searchTerm: "Hello",
			flags:      grepFlags{after: 1},
			expected:   "Hello\nworld\nHello again\nGrep is great\nHello world\nIgnore case\n",
		},
		{
			name:       "Before context",
			searchTerm: "Hello",
			flags:      grepFlags{before: 1},
			expected:   "Hello\nworld\nHello again\nGrep is great\nHello world\n",
		},
		{
			name:       "Context around",
			searchTerm: "Hello",
			flags:      grepFlags{context: 1},
			expected:   "Hello\nworld\nHello again\nGrep is great\nHello world\nIgnore case\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := &strings.Builder{}
			grep(input, test.searchTerm, test.flags, output)
			if output.String() != test.expected {
				t.Errorf("%s: grep with searchTerm '%s' and flags %+v, expected '%s', got '%s'", test.name, test.searchTerm, test.flags, test.expected, output.String())
			}
		})
	}
}
