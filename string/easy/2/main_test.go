package main

import (
	"reflect"
	"testing"
)

func TestCharFrequency(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
		{"empty string", args{""}, map[rune]int{}},
		{"single character", args{"a"}, map[rune]int{'a': 1}},
		{"repeated characters", args{"aaabbb"}, map[rune]int{'a': 3, 'b': 3}},
		{"mixed characters", args{"abcabc"}, map[rune]int{'a': 2, 'b': 2, 'c': 2}},
		{"with spaces", args{"hello world"}, map[rune]int{'h': 1, 'e': 1, 'l': 3, 'o': 2, ' ': 1, 'w': 1, 'r': 1, 'd': 1}},
		{"unicode characters", args{"你好世界"}, map[rune]int{'你': 1, '好': 1, '世': 1, '界': 1}},
		{"special characters", args{"!@#$%^&*()"}, map[rune]int{'!': 1, '@': 1, '#': 1, '$': 1, '%': 1, '^': 1, '&': 1, '*': 1, '(': 1, ')': 1}},
		{"numbers", args{"123123"}, map[rune]int{'1': 2, '2': 2, '3': 2}},
		{"mixed alphanumeric", args{"abc123!@#"}, map[rune]int{'a': 1, 'b': 1, 'c': 1, '1': 1, '2': 1, '3': 1, '!': 1, '@': 1, '#': 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CharFrequency(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CharFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
