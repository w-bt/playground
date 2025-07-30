package main

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1: mix of positives and negatives",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "Example 2: no valid triplet",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "Example 3: all zeroes",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "Edge case: more than three zeroes",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "Multiple valid combinations",
			nums: []int{-2, 0, 1, 1, 2},
			want: [][]int{
				{-2, 0, 2},
				{-2, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ThreeSum(tt.nums); !reflect.DeepEqual(gotResult, tt.want) {
				t.Errorf("ThreeSum() = %v, want %v", gotResult, tt.want)
			}
		})
	}
}
