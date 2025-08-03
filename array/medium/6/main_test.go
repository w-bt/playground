package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "All positive numbers",
			args: args{nums: []int{1, 2, 3, 4}},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "Includes zero",
			args: args{nums: []int{-1, 1, 0, -3, 3}},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			name: "Two elements",
			args: args{nums: []int{5, 6}},
			want: []int{6, 5},
		},
		{
			name: "All same elements",
			args: args{nums: []int{2, 2, 2, 2}},
			want: []int{8, 8, 8, 8},
		},
		{
			name: "Multiple zeros",
			args: args{nums: []int{0, 0, 1, 2}},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "Negative numbers",
			args: args{nums: []int{-1, -2, -3, -4}},
			want: []int{-24, -12, -8, -6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
