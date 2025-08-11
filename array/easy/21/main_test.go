package main

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty array",
			args: args{nums: []int{}},
			want: []string{},
		},
		{
			name: "Single element",
			args: args{nums: []int{5}},
			want: []string{"5"},
		},
		{
			name: "Consecutive sequence",
			args: args{nums: []int{1, 2, 3}},
			want: []string{"1->3"},
		},
		{
			name: "Example 1 from LeetCode",
			args: args{nums: []int{0, 1, 2, 4, 5, 7}},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			name: "Example 2 from LeetCode",
			args: args{nums: []int{0, 2, 3, 4, 6, 8, 9}},
			want: []string{"0", "2->4", "6", "8->9"},
		},
		{
			name: "No consecutive numbers",
			args: args{nums: []int{1, 3, 5}},
			want: []string{"1", "3", "5"},
		},
		{
			name: "Mixed consecutive and singles",
			args: args{nums: []int{-3, -2, -1, 1, 3, 4, 5, 7}},
			want: []string{"-3->-1", "1", "3->5", "7"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SummaryRanges(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SummaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
