package main

import "testing"

func TestFindLucky(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1 - only 2 is lucky",
			args: args{arr: []int{2, 2, 3, 4}},
			want: 2,
		},
		{
			name: "Example 2 - multiple lucky, return largest",
			args: args{arr: []int{1, 2, 2, 3, 3, 3}},
			want: 3,
		},
		{
			name: "Example 3 - no lucky number",
			args: args{arr: []int{2, 2, 2, 3, 3}},
			want: -1,
		},
		{
			name: "Single element, lucky",
			args: args{arr: []int{1}},
			want: 1,
		},
		{
			name: "Single element, not lucky",
			args: args{arr: []int{3}},
			want: -1,
		},
		{
			name: "All elements same and lucky",
			args: args{arr: []int{4, 4, 4, 4}},
			want: 4,
		},
		{
			name: "Large value not matching frequency",
			args: args{arr: []int{500, 500, 500}},
			want: -1,
		},
		{
			name: "Two lucky numbers, pick largest",
			args: args{arr: []int{2, 2, 3, 3, 3}},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLucky(tt.args.arr); got != tt.want {
				t.Errorf("FindLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}
