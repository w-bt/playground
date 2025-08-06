package main

import "testing"

func TestStrStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1 - needle at start",
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			name: "Example 2 - no occurrence",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
		{
			name: "Needle equals haystack",
			args: args{
				haystack: "abc",
				needle:   "abc",
			},
			want: 0,
		},
		{
			name: "Needle longer than haystack",
			args: args{
				haystack: "short",
				needle:   "longerstring",
			},
			want: -1,
		},
		{
			name: "Needle in middle",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "Needle in middle",
			args: args{
				haystack: "babbbbbabb",
				needle:   "bbab",
			},
			want: 5,
		},
		{
			name: "Repeated patterns",
			args: args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			want: -1,
		},
		{
			name: "Needle at end",
			args: args{
				haystack: "mississippi",
				needle:   "issip",
			},
			want: 4,
		},
		{
			name: "Needle at end",
			args: args{
				haystack: "aabaaabaaac",
				needle:   "aabaaac",
			},
			want: 4,
		},
		{
			name: "Needle at end",
			args: args{
				haystack: "mississippi",
				needle:   "pi",
			},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("StrStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
