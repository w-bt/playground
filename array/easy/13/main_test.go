package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1 - small array with duplicates",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "Example 2 - larger array with duplicates",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
		{
			name: "All elements same",
			args: args{
				nums: []int{5, 5, 5, 5, 5},
			},
			want: 1,
		},
		{
			name: "No duplicates",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "Empty array",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "Single element",
			args: args{
				nums: []int{7},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
