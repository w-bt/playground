package main

import "testing"

func TestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple parentheses",
			args: args{s: "()"},
			want: true,
		},
		{
			name: "multiple pairs",
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			name: "mismatched pair",
			args: args{s: "(]"},
			want: false,
		},
		{
			name: "interleaved brackets",
			args: args{s: "([)]"},
			want: false,
		},
		{
			name: "nested brackets",
			args: args{s: "{[]}"},
			want: true,
		},
		{
			name: "single bracket",
			args: args{s: "["},
			want: false,
		},
		{
			name: "closing first",
			args: args{s: "])"},
			want: false,
		},
		{
			name: "empty string",
			args: args{s: ""},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("ValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
