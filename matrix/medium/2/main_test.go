package main

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1: 3x3 square matrix",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "Example 2: 3x4 rectangular matrix",
			args: args{
				matrix: [][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
				},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "1x1 matrix",
			args: args{
				matrix: [][]int{{42}},
			},
			want: []int{42},
		},
		{
			name: "1 row only",
			args: args{
				matrix: [][]int{{1, 2, 3, 4}},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "1 column only",
			args: args{
				matrix: [][]int{
					{1},
					{2},
					{3},
					{4},
				},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Empty matrix",
			args: args{
				matrix: [][]int{},
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
