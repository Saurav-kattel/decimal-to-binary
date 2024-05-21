package test

import (
	"side-project/decimal-bin/converter"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		input       string
		expected    int
		expectError bool
	}{
		{"123\n", 123, false},
		{"456\n", 456, false},
		{"-789\n", -789, false},
		{"abc\n", 0, true},
		{"12.34\n", 0, true},
		{"", 0, true},
	}

	for _, test := range tests {
		converter := &converter.Converter{}
		err := converter.ParseInput(test.input)
		if (err != nil) != test.expectError {
			t.Errorf("parseInput(%q) error = %v, expectError %v", test.input, err, test.expectError)
		}
		if !test.expectError && converter.UserInput != test.expected {
			t.Errorf("parseInput(%q) = %d, expected %d", test.input, converter.UserInput, test.expected)
		}
	}
}
