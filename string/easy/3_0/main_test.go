package main

import "testing"

func TestIsAnagram(t *testing.T) {
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
			name: "Valid anagram",
			args: args{"anagram", "nagaram"},
			want: true,
		},
		{
			name: "Not anagram (different chars)",
			args: args{"rat", "car"},
			want: false,
		},
		{
			name: "Valid anagram (listen/silent)",
			args: args{"listen", "silent"},
			want: true,
		},
		{
			name: "Different lengths",
			args: args{"abc", "ab"},
			want: false,
		},
		{
			name: "Same chars, different counts",
			args: args{"aabb", "ab"},
			want: false,
		},
		{
			name: "Identical strings",
			args: args{"test", "test"},
			want: true,
		},
		{
			name: "Completely different",
			args: args{"hello", "world"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
