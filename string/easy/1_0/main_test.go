package main

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"hello world", "hello world", "dlrow olleh"},
		{"OpenAI", "OpenAI", "IAnepO"},
		{"single a", "a", "a"},
		{"HelloWorld", "HelloWorld", "dlroWolleH"},
		{"single x", "x", "x"},
		{"palindrome", "madam", "madam"},
		{"special chars", "abc123!@#", "#@!321cba"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseString(tt.input)
			if got != tt.expected {
				t.Errorf("ReverseString(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
