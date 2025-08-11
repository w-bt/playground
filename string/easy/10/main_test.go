package main

import "testing"

func TestWordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1 (true)",
			args: args{pattern: "abba", s: "dog cat cat dog"},
			want: true,
		},
		{
			name: "Example 2 (false)",
			args: args{pattern: "abba", s: "dog cat cat fish"},
			want: false,
		},
		{
			name: "Example 3 (false same char->different words)",
			args: args{pattern: "aaaa", s: "dog cat cat dog"},
			want: false,
		},
		{
			name: "Example 4 (false different chars->same word)",
			args: args{pattern: "abba", s: "dog dog dog dog"},
			want: false,
		},
		{
			name: "All same char matches same word",
			args: args{pattern: "aaaa", s: "dog dog dog dog"},
			want: true,
		},
		{
			name: "Single char and single word",
			args: args{pattern: "a", s: "dog"},
			want: true,
		},
		{
			name: "Length mismatch",
			args: args{pattern: "ab", s: "dog"},
			want: false,
		},
		{
			name: "Bijection holds with distinct words",
			args: args{pattern: "abc", s: "red blue green"},
			want: true,
		},
		{
			name: "Two pattern chars map to same word (violation)",
			args: args{pattern: "abc", s: "red red green"},
			want: false,
		},
		{
			name: "Alternating pattern valid",
			args: args{pattern: "abab", s: "one two one two"},
			want: true,
		},
		{
			name: "Bijection broken mid-string",
			args: args{pattern: "abca", s: "dog cat cat dog"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("WordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
