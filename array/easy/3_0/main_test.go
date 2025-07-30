package main

import (
	"reflect"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
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
			name: "both arrays sorted and interleaved",
			args: args{nums1: []int{1, 3, 5}, nums2: []int{2, 4, 6}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "first array empty",
			args: args{nums1: []int{}, nums2: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "second array empty",
			args: args{nums1: []int{1, 2, 3}, nums2: []int{}},
			want: []int{1, 2, 3},
		},
		{
			name: "both arrays empty",
			args: args{nums1: []int{}, nums2: []int{}},
			want: []int{},
		},
		{
			name: "arrays with duplicate values",
			args: args{nums1: []int{1, 2}, nums2: []int{2, 3}},
			want: []int{1, 2, 2, 3},
		},
		{
			name: "already merged order",
			args: args{nums1: []int{1, 4}, nums2: []int{2, 3}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "different lengths",
			args: args{nums1: []int{1, 2}, nums2: []int{3, 4, 5, 6}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "negative numbers",
			args: args{nums1: []int{-3, -1, 0}, nums2: []int{-2, 2, 3}},
			want: []int{-3, -2, -1, 0, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := MergeSortedArrays(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(gotResult, tt.want) {
				t.Errorf("MergeSortedArrays() = %v, want %v", gotResult, tt.want)
			}
		})
	}
}
