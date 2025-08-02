package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1 - Mixed positive and negative",
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			want: 6, // subarray [4, -1, 2, 1]
		},
		{
			name: "Example 2 - Single positive element",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "Example 3 - All positive elements",
			args: args{nums: []int{5, 4, -1, 7, 8}},
			want: 23, // entire array
		},
		{
			name: "All negative values",
			args: args{nums: []int{-5, -4, -3, -2, -1}},
			want: -1, // pick the largest single number
		},
		{
			name: "Single negative element",
			args: args{nums: []int{-10}},
			want: -10,
		},
		{
			name: "Large negative in middle",
			args: args{nums: []int{1, 2, 3, -100, 4, 5, 6}},
			want: 15, // subarray [4, 5, 6]
		},
		{
			name: "Zero breaks negative",
			args: args{nums: []int{-2, 0, -1}},
			want: 0,
		},
		{
			name: "Single element zero",
			args: args{nums: []int{0}},
			want: 0,
		},
		{
			name: "Multiple peaks",
			args: args{nums: []int{1, -2, 3, 5, -1, 2, -1, 2, -3, 4}},
			want: 11, // subarray [3, 5, -1, 2, -1, 2, -3, 4]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("MaxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
