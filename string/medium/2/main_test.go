package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Example 1 - multiple anagram groups",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			name: "Example 2 - empty string",
			args: args{strs: []string{""}},
			want: [][]string{
				{""},
			},
		},
		{
			name: "Example 3 - single character",
			args: args{strs: []string{"a"}},
			want: [][]string{
				{"a"},
			},
		},
		{
			name: "No anagram group",
			args: args{strs: []string{"abc", "def", "ghi"}},
			want: [][]string{
				{"abc"}, {"def"}, {"ghi"},
			},
		},
		{
			name: "All are the same anagram group",
			args: args{strs: []string{"listen", "silent", "enlist"}},
			want: [][]string{
				{"listen", "silent", "enlist"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.args.strs)
			if !reflect.DeepEqual(normalize(got), normalize(tt.want)) {
				t.Errorf("GroupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func normalize(output [][]string) [][]string {
	for i := range output {
		sort.Strings(output[i]) // Sort each group
	}
	sort.Slice(output, func(i, j int) bool { // Sort groups by first element
		return output[i][0] < output[j][0]
	})
	return output
}
