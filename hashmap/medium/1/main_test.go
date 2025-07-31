package main

import (
	"reflect"
	"sort"
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
			name: "Basic case with multiple frequent numbers",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "Single element array",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "Two elements with same frequency",
			args: args{
				nums: []int{4, 4, 6, 6, 7},
				k:    2,
			},
			want: []int{4, 6},
		},
		{
			name: "All unique elements, k equals array size",
			args: args{
				nums: []int{9, 8, 7, 6, 5},
				k:    5,
			},
			want: []int{5, 6, 7, 8, 9},
		},
		{
			name: "All same elements",
			args: args{
				nums: []int{2, 2, 2, 2},
				k:    1,
			},
			want: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TopKFrequent(tt.args.nums, tt.args.k)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
