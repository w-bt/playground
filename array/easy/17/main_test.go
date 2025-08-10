package main

import "testing"

func TestSingleNumber(t *testing.T) {
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
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "Example case 1",
			args: args{
				nums: []int{2, 2, 1},
			},
			want: 1,
		},
		{
			name: "Example case 2",
			args: args{
				nums: []int{4, 1, 2, 1, 2},
			},
			want: 4,
		},
		{
			name: "Negative numbers",
			args: args{
				nums: []int{-1, -1, -2},
			},
			want: -2,
		},
		{
			name: "Mixed positive and negative",
			args: args{
				nums: []int{-3, 1, 1, -3, 5},
			},
			want: 5,
		},
		{
			name: "Large array",
			args: args{
				nums: []int{10, 14, 10, 18, 14, 20, 20},
			},
			want: 18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumber(tt.args.nums); got != tt.want {
				t.Errorf("SingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
