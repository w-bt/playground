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
			name: "With Duplicates",
			args: args{[]int{1, 2, 3, 1}},
			want: true,
		},
		{
			name: "All Unique",
			args: args{[]int{1, 2, 3, 4}},
			want: false,
		},
		{
			name: "Single Element",
			args: args{[]int{1}},
			want: false,
		},
		{
			name: "All Identical",
			args: args{[]int{5, 5, 5, 5}},
			want: true,
		},
		{
			name: "Empty Array",
			args: args{[]int{}},
			want: false,
		},
		{
			name: "With Negative Duplicate",
			args: args{[]int{1, -1, 2, -2, 1}},
			want: true,
		},
		{
			name: "All Unique Larger",
			args: args{[]int{10, 20, 30, 40, 50}},
			want: false,
		},
		{
			name: "Large Extreme Unique",
			args: args{[]int{999999999, -999999999}},
			want: false,
		},
		{
			name: "Zero Duplicate",
			args: args{[]int{0, 0}},
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
