package main

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "interleaved zeroes",
			args: args{nums: []int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "trailing zeroes",
			args: args{nums: []int{1, 2, 3, 0, 0}},
			want: []int{1, 2, 3, 0, 0},
		},
		{
			name: "leading zeroes",
			args: args{nums: []int{0, 0, 1, 2}},
			want: []int{1, 2, 0, 0},
		},
		{
			name: "all zeroes",
			args: args{nums: []int{0, 0, 0}},
			want: []int{0, 0, 0},
		},
		{
			name: "no zeroes",
			args: args{nums: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "alternating zeroes",
			args: args{nums: []int{0, 1, 0, 2, 0, 3}},
			want: []int{1, 2, 3, 0, 0, 0},
		},
		{
			name: "single zero",
			args: args{nums: []int{0}},
			want: []int{0},
		},
		{
			name: "single non-zero",
			args: args{nums: []int{5}},
			want: []int{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoveZeroes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
