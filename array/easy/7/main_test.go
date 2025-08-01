package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1 - Contains duplicate",
			args: args{nums: []int{1, 2, 3, 1}},
			want: true,
		},
		{
			name: "Example 2 - All distinct",
			args: args{nums: []int{1, 2, 3, 4}},
			want: false,
		},
		{
			name: "Example 3 - Multiple duplicates",
			args: args{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			want: true,
		},
		{
			name: "Single element",
			args: args{nums: []int{99}},
			want: false,
		},
		{
			name: "Empty array",
			args: args{nums: []int{}},
			want: false,
		},
		{
			name: "Large distinct values",
			args: args{nums: []int{-1000000000, 0, 1000000000}},
			want: false,
		},
		{
			name: "Edge case - large array with duplicates at end",
			args: args{nums: append(make([]int, 99999), 42, 42)},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("ContainsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
