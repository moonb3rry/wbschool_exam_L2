package main

import (
	"errors"
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errors.New("invalid string")},
		{"", "", nil},
		{"qwe\\4\\5", "qwe45", nil},
		{"qwe\\\\5", "qwe\\\\\\\\\\", nil},
	}

	for _, test := range tests {
		result, err := unpack(test.input)
		if test.err != nil && err == nil {
			t.Errorf("unpack(%s) expected error, got nil", test.input)
		}
		if test.err == nil && err != nil {
			t.Errorf("unpack(%s) unexpected error: %s", test.input, err)
		}
		if result != test.expected {
			t.Errorf("unpack(%s) = %s, want %s", test.input, result, test.expected)
		}
	}
}
