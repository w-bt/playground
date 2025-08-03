package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "Example 2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "Example 3 - duplicate values",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
		{
			name: "Negative numbers",
			args: args{
				nums:   []int{-1, -2, -3, -4, -5},
				target: -8,
			},
			want: []int{2, 4},
		},
		{
			name: "Mixed numbers",
			args: args{
				nums:   []int{-10, 20, 10, -5, 5},
				target: 15,
			},
			want: []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
