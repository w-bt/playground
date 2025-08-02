package main

import "testing"

func TestMaxProduct(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1 - Positive and negative values",
			args: args{nums: []int{2, 3, -2, 4}},
			want: 6, // subarray [2, 3]
		},
		{
			name: "Example 2 - With zero in the middle",
			args: args{nums: []int{-2, 0, -1}},
			want: 0, // max product is [0]
		},
		{
			name: "Single negative number",
			args: args{nums: []int{-5}},
			want: -5,
		},
		{
			name: "All positive numbers",
			args: args{nums: []int{1, 2, 3, 4}},
			want: 24, // entire array
		},
		{
			name: "Two negatives separated by zero",
			args: args{nums: []int{-1, 0, -2}},
			want: 0,
		},
		{
			name: "Alternating signs",
			args: args{nums: []int{2, -5, -2, -4, 3}},
			want: 24, // subarray [-2, -4, 3]
		},
		{
			name: "Multiple zeros",
			args: args{nums: []int{0, -2, 0, -3, 0}},
			want: 0,
		},
		{
			name: "Product resets due to 0",
			args: args{nums: []int{2, 3, 0, 4, 5}},
			want: 20, // subarray [4, 5]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProduct(tt.args.nums); got != tt.want {
				t.Errorf("MaxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
