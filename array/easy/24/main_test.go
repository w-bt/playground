package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "both-empty",
			args: args{nums1: []int{}, nums2: []int{}},
			want: []int{},
		},
		{
			name: "one-empty",
			args: args{nums1: []int{1, 2}, nums2: []int{}},
			want: []int{},
		},
		{
			name: "no-intersection",
			args: args{nums1: []int{1, 3, 5}, nums2: []int{2, 4, 6}},
			want: []int{},
		},
		{
			name: "example-duplicates",
			args: args{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}},
			want: []int{2},
		},
		{
			name: "example-two",
			args: args{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}},
			want: []int{4, 9},
		},
		{
			name: "negatives",
			args: args{nums1: []int{-1, -2, -3, -3}, nums2: []int{-3, 0, 2, -2}},
			want: []int{-3, -2},
		},
		{
			name: "all-same-arrays",
			args: args{nums1: []int{7, 7, 7}, nums2: []int{7, 7}},
			want: []int{7},
		},
		{
			name: "mixed-order",
			args: args{nums1: []int{3, 1, 2, 3}, nums2: []int{2, 3, 4, 3, 2}},
			want: []int{2, 3},
		},
		{
			name: "with-zero",
			args: args{nums1: []int{0, 0, 1}, nums2: []int{0, 2, 2}},
			want: []int{0},
		},
		{
			name: "large-dupes",
			args: args{nums1: []int{5, 5, 5, 6, 6, 7}, nums2: []int{5, 6, 6, 6, 8}},
			want: []int{5, 6},
		},
		{
			name: "single-match",
			args: args{nums1: []int{10, 11, 12}, nums2: []int{9, 10, 13}},
			want: []int{10},
		},
		{
			name: "overlap-three",
			args: args{nums1: []int{1, 1, 2, 3, 5, 8}, nums2: []int{0, 2, 3, 4, 5}},
			want: []int{2, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersection(tt.args.nums1, tt.args.nums2)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
