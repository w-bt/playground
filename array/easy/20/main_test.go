package main

import "testing"

func TestMoveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Example case 1",
			args: args{
				nums: []int{0, 1, 0, 3, 12}, // expected after move: [1, 3, 12, 0, 0]
			},
		},
		{
			name: "Example case 2",
			args: args{
				nums: []int{0}, // expected after move: [0]
			},
		},
		{
			name: "All zeros",
			args: args{
				nums: []int{0, 0, 0, 0}, // expected after move: [0, 0, 0, 0]
			},
		},
		{
			name: "No zeros",
			args: args{
				nums: []int{1, 2, 3, 4}, // expected after move: [1, 2, 3, 4]
			},
		},
		{
			name: "Zeros at start",
			args: args{
				nums: []int{0, 0, 1, 2}, // expected after move: [1, 2, 0, 0]
			},
		},
		{
			name: "Zeros at end",
			args: args{
				nums: []int{1, 2, 0, 0}, // expected after move: [1, 2, 0, 0]
			},
		},
		{
			name: "Mixed zeros in middle",
			args: args{
				nums: []int{4, 0, 5, 0, 6}, // expected after move: [4, 5, 6, 0, 0]
			},
		},
		{
			name: "Single element non-zero",
			args: args{
				nums: []int{7}, // expected after move: [7]
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MoveZeroes(tt.args.nums)
		})
	}
}
