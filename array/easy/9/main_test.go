package main

import "testing"

func TestMissingNumber(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1 - Missing 2 from [0,1,3]",
			args: args{nums: []int{3, 0, 1}},
			want: 2,
		},
		{
			name: "Example 2 - Missing 2 from [0,1]",
			args: args{nums: []int{0, 1}},
			want: 2,
		},
		{
			name: "Example 3 - Missing 8 from [0..9] without 8",
			args: args{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}},
			want: 8,
		},
		{
			name: "Single element - Missing 0",
			args: args{nums: []int{1}},
			want: 0,
		},
		{
			name: "Missing 1 from [0]",
			args: args{nums: []int{0}},
			want: 1,
		},
		{
			name: "Full range except last (missing 5)",
			args: args{nums: []int{0, 1, 2, 3, 4}},
			want: 5,
		},
		{
			name: "Full range except first (missing 0)",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MissingNumber(tt.args.nums); got != tt.want {
				t.Errorf("MissingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
