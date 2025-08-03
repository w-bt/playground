package main

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "Example 2 - Single element",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "All elements same frequency",
			args: args{
				nums: []int{4, 5, 6, 7},
				k:    2,
			},
			want: []int{4, 5},
		},
		{
			name: "Multiple duplicates, pick top 1",
			args: args{
				nums: []int{9, 9, 9, 1, 2, 2, 3},
				k:    1,
			},
			want: []int{9},
		},
		{
			name: "k equals number of unique elements",
			args: args{
				nums: []int{5, 7, 5, 7, 8},
				k:    3,
			},
			want: []int{5, 7, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
