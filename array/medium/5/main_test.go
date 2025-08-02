package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		intervals [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1 - Overlapping intervals",
			args: args{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "Example 2 - End meets start",
			args: args{intervals: [][]int{{1, 4}, {4, 5}}},
			want: [][]int{{1, 5}},
		},
		{
			name: "No overlap",
			args: args{intervals: [][]int{{1, 2}, {3, 4}, {5, 6}}},
			want: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name: "All intervals overlap into one",
			args: args{intervals: [][]int{{1, 10}, {2, 3}, {4, 5}, {6, 7}}},
			want: [][]int{{1, 10}},
		},
		{
			name: "Unsorted input",
			args: args{intervals: [][]int{{5, 6}, {1, 3}, {2, 4}}},
			want: [][]int{{1, 4}, {5, 6}},
		},
		{
			name: "Single interval",
			args: args{intervals: [][]int{{1, 2}}},
			want: [][]int{{1, 2}},
		},
		{
			name: "Empty input",
			args: args{intervals: [][]int{}},
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
