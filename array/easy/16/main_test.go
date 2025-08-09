package main

import "testing"

func TestMerge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Example 1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
		},
		{
			name: "Example 2 - nums2 empty",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
		},
		{
			name: "Example 3 - nums1 empty initialized part",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
		},
		{
			name: "Interleaving merge",
			args: args{
				nums1: []int{1, 3, 5, 0, 0, 0},
				m:     3,
				nums2: []int{2, 4, 6},
				n:     3,
			},
		},
		{
			name: "nums2 smaller elements",
			args: args{
				nums1: []int{4, 5, 6, 0, 0, 0},
				m:     3,
				nums2: []int{1, 2, 3},
				n:     3,
			},
		},
		{
			name: "nums1 smaller elements",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{4, 5, 6},
				n:     3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		})
	}
}
