package main

import "testing"

func TestMajorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Single element array",
			args: args{
				nums: []int{5},
			},
			want: 5,
		},
		{
			name: "Example case 1",
			args: args{
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
		{
			name: "Example case 2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
		{
			name: "Majority at the start",
			args: args{
				nums: []int{4, 4, 4, 2, 2},
			},
			want: 4,
		},
		{
			name: "Majority at the end",
			args: args{
				nums: []int{1, 3, 3, 3, 3, 2, 3},
			},
			want: 3,
		},
		{
			name: "Negative numbers majority",
			args: args{
				nums: []int{-1, -1, -1, 2, 3, -1, 4},
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElement(tt.args.nums); got != tt.want {
				t.Errorf("MajorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
