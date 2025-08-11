package main

import "testing"

func TestIsSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example-true",
			args: args{s: "abc", t: "ahbgdc"},
			want: true,
		},
		{
			name: "example-false",
			args: args{s: "axc", t: "ahbgdc"},
			want: false,
		},
		{
			name: "empty-s-true",
			args: args{s: "", t: "anything"},
			want: true,
		},
		{
			name: "both-empty-true",
			args: args{s: "", t: ""},
			want: true,
		},
		{
			name: "s-longer-than-t-false",
			args: args{s: "abc", t: "ab"},
			want: false,
		},
		{
			name: "exact-equal-true",
			args: args{s: "abc", t: "abc"},
			want: true,
		},
		{
			name: "repeated-chars-enough-true",
			args: args{s: "aaa", t: "aaaaa"},
			want: true,
		},
		{
			name: "repeated-chars-not-enough-false",
			args: args{s: "aaaa", t: "aaa"},
			want: false,
		},
		{
			name: "char-not-present-false",
			args: args{s: "z", t: "abc"},
			want: false,
		},
		{
			name: "order-matters-false",
			args: args{s: "ace", t: "aec"},
			want: false,
		},
		{
			name: "duplicate-with-gaps-true",
			args: args{s: "abbc", t: "abdbc"},
			want: true,
		},
		{
			name: "many-duplicates-true",
			args: args{s: "abc", t: "aabbcc"},
			want: true,
		},
		{
			name: "single-char-present-true",
			args: args{s: "a", t: "a"},
			want: true,
		},
		{
			name: "single-char-absent-false",
			args: args{s: "a", t: ""},
			want: false,
		},
		{
			name: "match-at-end-true",
			args: args{s: "t", t: "bbbbbt"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
