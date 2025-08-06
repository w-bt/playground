package main

import "testing"

func TestRemoveElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1 - basic removal",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "Example 2 - multiple occurrences",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
		{
			name: "No occurrence of val",
			args: args{
				nums: []int{1, 2, 3, 4},
				val:  5,
			},
			want: 4,
		},
		{
			name: "All elements are val",
			args: args{
				nums: []int{7, 7, 7},
				val:  7,
			},
			want: 0,
		},
		{
			name: "Empty slice",
			args: args{
				nums: []int{},
				val:  1,
			},
			want: 0,
		},
		{
			name: "Single element equals val",
			args: args{
				nums: []int{5},
				val:  5,
			},
			want: 0,
		},
		{
			name: "Single element not equals val",
			args: args{
				nums: []int{5},
				val:  9,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
