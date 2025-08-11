package main

import "testing"

func TestFindTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "basic-sorted",
			args: args{s: "abcd", t: "abcde"},
			want: byte('e'),
		},
		{
			name: "empty-s",
			args: args{s: "", t: "y"},
			want: byte('y'),
		},
		{
			name: "permuted",
			args: args{s: "abcd", t: "cadbe"},
			want: byte('e'),
		},
		{
			name: "duplicate-extra-b",
			args: args{s: "aabb", t: "ababb"},
			want: byte('b'),
		},
		{
			name: "duplicate-extra-a",
			args: args{s: "aabb", t: "baaab"},
			want: byte('a'),
		},
		{
			name: "all-same-letters",
			args: args{s: "zzzz", t: "zzzzz"},
			want: byte('z'),
		},
		{
			name: "single-char",
			args: args{s: "x", t: "xx"},
			want: byte('x'),
		},
		{
			name: "mixed-permutation-extra-m",
			args: args{s: "abcxyz", t: "zxycbam"},
			want: byte('m'),
		},
		{
			name: "keyboard-letters-extra-p",
			args: args{s: "qwertyuiop", t: "poiuytqrewp"},
			want: byte('p'),
		},
		{
			name: "many-a-extra-b",
			args: args{s: "aaaaaaaaab", t: "baaaaaaaaba"},
			want: byte('b'),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("FindTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
