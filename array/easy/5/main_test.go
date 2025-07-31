package main

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "all ones",
			args: args{nums: []int{1, 1, 1, 1}},
			want: 4,
		},
		{
			name: "all zeroes",
			args: args{nums: []int{0, 0, 0, 0}},
			want: 0,
		},
		{
			name: "alternating ones and zeroes",
			args: args{nums: []int{1, 0, 1, 0, 1}},
			want: 1,
		},
		{
			name: "ending with longest streak",
			args: args{nums: []int{0, 1, 1, 1}},
			want: 3,
		},
		{
			name: "starting with longest streak",
			args: args{nums: []int{1, 1, 1, 0, 0}},
			want: 3,
		},
		{
			name: "middle longest streak",
			args: args{nums: []int{0, 1, 1, 0, 1, 1, 1, 0}},
			want: 3,
		},
		{
			name: "single one",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "single zero",
			args: args{nums: []int{0}},
			want: 0,
		},
		{
			name: "mixed with one longest",
			args: args{nums: []int{1, 0, 1, 1, 0, 1}},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("FindMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
