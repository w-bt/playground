package main

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1 - Valid aabcbc",
			args: args{s: "aabcbc"},
			want: true,
		},
		{
			name: "Example 2 - Valid abcabcababcc",
			args: args{s: "abcabcababcc"},
			want: true,
		},
		{
			name: "Example 3 - Invalid abccba",
			args: args{s: "abccba"},
			want: false,
		},
		{
			name: "Edge Case - Empty string",
			args: args{s: ""},
			want: true, // since t is initially "" and no operations are needed
		},
		{
			name: "Single abc is valid",
			args: args{s: "abc"},
			want: true,
		},
		{
			name: "Multiple valid insertions",
			args: args{s: "abcabcabc"},
			want: true,
		},
		{
			name: "Starts with invalid pattern",
			args: args{s: "aabc"},
			want: false,
		},
		{
			name: "Invalid with extra character",
			args: args{s: "aabcbcc"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
