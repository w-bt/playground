package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple palindrome",
			args: args{s: "racecar"},
			want: true,
		},
		{
			name: "not a palindrome",
			args: args{s: "hello"},
			want: false,
		},
		{
			name: "mixed case with spaces and punctuation",
			args: args{s: "A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "only punctuation",
			args: args{s: "!!!???"},
			want: true, // becomes empty string, which is valid palindrome
		},
		{
			name: "numeric palindrome",
			args: args{s: "12321"},
			want: true,
		},
		{
			name: "numeric non-palindrome",
			args: args{s: "12345"},
			want: false,
		},
		{
			name: "empty string",
			args: args{s: ""},
			want: true,
		},
		{
			name: "single character",
			args: args{s: "a"},
			want: true,
		},
		{
			name: "capital letters only",
			args: args{s: "ABBA"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
