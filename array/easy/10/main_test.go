package main

import "testing"

func TestArraySign(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "All negative with even count",
			args: args{nums: []int{-1, -2, -3, -4}},
			want: 1, // product is positive
		},
		{
			name: "All negative with odd count",
			args: args{nums: []int{-1, -2, -3}},
			want: -1, // product is negative
		},
		{
			name: "Contains zero",
			args: args{nums: []int{1, -2, 0, 3}},
			want: 0, // product is zero
		},
		{
			name: "All positive",
			args: args{nums: []int{2, 3, 4}},
			want: 1, // product is positive
		},
		{
			name: "Single zero",
			args: args{nums: []int{0}},
			want: 0,
		},
		{
			name: "Single positive",
			args: args{nums: []int{5}},
			want: 1,
		},
		{
			name: "Single negative",
			args: args{nums: []int{-7}},
			want: -1,
		},
		{
			name: "Mixed signs, odd negatives",
			args: args{nums: []int{1, -1, 2}},
			want: -1,
		},
		{
			name: "Mixed signs, even negatives",
			args: args{nums: []int{1, -1, -1, 2}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArraySign(tt.args.nums); got != tt.want {
				t.Errorf("ArraySign() = %v, want %v", got, tt.want)
			}
		})
	}
}
