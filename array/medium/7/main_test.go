package main

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Target found in rotated array",
			args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0},
			want: 4,
		},
		{
			name: "Target not in rotated array",
			args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3},
			want: -1,
		},
		{
			name: "Single element - target not found",
			args: args{nums: []int{1}, target: 0},
			want: -1,
		},
		{
			name: "Single element - target found",
			args: args{nums: []int{1}, target: 1},
			want: 0,
		},
		{
			name: "Target at beginning of array",
			args: args{nums: []int{6, 7, 8, 1, 2, 3, 4, 5}, target: 6},
			want: 0,
		},
		{
			name: "Target at end of array",
			args: args{nums: []int{6, 7, 8, 1, 2, 3, 4, 5}, target: 5},
			want: 7,
		},
		{
			name: "Target in middle of rotated part",
			args: args{nums: []int{5, 6, 7, 0, 1, 2, 3, 4}, target: 2},
			want: 5,
		},
		{
			name: "Array not rotated",
			args: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, target: 3},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
