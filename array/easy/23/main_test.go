package main

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "empty",
			args: args{s: []byte("")},
		},
		{
			name: "single-char",
			args: args{s: []byte("a")},
		},
		{
			name: "two-chars",
			args: args{s: []byte("ab")},
		},
		{
			name: "odd-length",
			args: args{s: []byte("abc")},
		},
		{
			name: "even-length",
			args: args{s: []byte("abcd")},
		},
		{
			name: "palindrome",
			args: args{s: []byte("level")},
		},
		{
			name: "all-same",
			args: args{s: []byte("aaaaa")},
		},
		{
			name: "with-spaces",
			args: args{s: []byte("a b c")},
		},
		{
			name: "with-punctuation",
			args: args{s: []byte("h!i?")},
		},
		{
			name: "digits",
			args: args{s: []byte("12345")},
		},
		{
			name: "mixed-case",
			args: args{s: []byte("AbCdE")},
		},
		{
			name: "leading-trailing-space",
			args: args{s: []byte(" hello ")},
		},
		{
			name: "longer",
			args: args{s: []byte("thequickbrownfox")},
		},
		{
			name: "repeated-pattern",
			args: args{s: []byte("abababab")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseString(tt.args.s)
		})
	}
}
