package main

import "testing"

func TestFirstUniqueChar(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Unique at beginning",
			args: args{"leetcode"},
			want: 0,
		},
		{
			name: "Unique in middle",
			args: args{"loveleetcode"},
			want: 2,
		},
		{
			name: "No unique character",
			args: args{"aabb"},
			want: -1,
		},
		{
			name: "Single character",
			args: args{"z"},
			want: 0,
		},
		{
			name: "All unique characters",
			args: args{"abcde"},
			want: 0,
		},
		{
			name: "Unique at end",
			args: args{"aabbc"},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstUniqueChar(tt.args.s); got != tt.want {
				t.Errorf("FirstUniqueChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
