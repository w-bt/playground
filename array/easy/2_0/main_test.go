package main

import (
	"reflect"
	"testing"
)

func TestPrefixSuffixSum(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name       string
		args       args
		wantFirst  []int
		wantSecond []int
	}{
		{"empty", args{[]int{}}, nil, nil},
		{"single element", args{[]int{5}}, []int{5}, []int{5}},
		{"multiple elements", args{[]int{1, 2, 3, 4}}, []int{1, 3, 6, 10}, []int{10, 9, 7, 4}},
		{"negative numbers", args{[]int{-1, -2, -3}}, []int{-1, -3, -6}, []int{-6, -5, -3}},
		{"mixed numbers", args{[]int{1, -2, 3, -4}}, []int{1, -1, 2, -2}, []int{-2, -3, -1, -4}},
		{"all zeros", args{[]int{0, 0, 0}}, []int{0, 0, 0}, []int{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := PrefixSuffixSum(tt.args.input)
			if !reflect.DeepEqual(got, tt.wantFirst) {
				t.Errorf("PrefixSuffixSum() got = %v, wantFirst %v", got, tt.wantFirst)
			}
			if !reflect.DeepEqual(got1, tt.wantSecond) {
				t.Errorf("PrefixSuffixSum() got1 = %v, wantFirst %v", got1, tt.wantSecond)
			}
		})
	}
}
