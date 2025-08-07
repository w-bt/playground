package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "common prefix exists",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "no common prefix",
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "single string",
			args: args{strs: []string{"alone"}},
			want: "alone",
		},
		{
			name: "all strings identical",
			args: args{strs: []string{"same", "same", "same"}},
			want: "same",
		},
		{
			name: "some strings empty",
			args: args{strs: []string{"", "empty", "emp"}},
			want: "",
		},
		{
			name: "all empty strings",
			args: args{strs: []string{"", "", ""}},
			want: "",
		},
		{
			name: "partial match",
			args: args{strs: []string{"interspecies", "interstellar", "interstate"}},
			want: "inters",
		},
		{
			name: "case sensitive",
			args: args{strs: []string{"Test", "test", "testing"}},
			want: "",
		},
		{
			name: "prefix is entire shortest word",
			args: args{strs: []string{"a", "ab", "abc"}},
			want: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
