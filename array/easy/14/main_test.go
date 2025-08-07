package main

import "testing"

func TestSearchInsert(t *testing.T) {
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
			name: "target found in middle",
			args: args{nums: []int{1, 3, 5, 6}, target: 5},
			want: 2,
		},
		{
			name: "target should be inserted at beginning",
			args: args{nums: []int{1, 3, 5, 6}, target: 0},
			want: 0,
		},
		{
			name: "target should be inserted in middle",
			args: args{nums: []int{1, 3, 5, 6}, target: 2},
			want: 1,
		},
		{
			name: "target should be inserted at end",
			args: args{nums: []int{1, 3, 5, 6}, target: 7},
			want: 4,
		},
		{
			name: "target equals first element",
			args: args{nums: []int{2, 4, 6, 8}, target: 2},
			want: 0,
		},
		{
			name: "target equals last element",
			args: args{nums: []int{2, 4, 6, 8}, target: 8},
			want: 3,
		},
		{
			name: "single element equal target",
			args: args{nums: []int{3}, target: 3},
			want: 0,
		},
		{
			name: "single element less than target",
			args: args{nums: []int{3}, target: 5},
			want: 1,
		},
		{
			name: "single element greater than target",
			args: args{nums: []int{3}, target: 1},
			want: 0,
		},
		{
			name: "empty array",
			args: args{nums: []int{}, target: 1},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
