package main

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1 - merge one interval",
			args: args{
				intervals:   [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 5},
			},
			want: [][]int{{1, 5}, {6, 9}},
		},
		{
			name: "Example 2 - merge multiple intervals",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 8},
			},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name: "Insert with no overlap at beginning",
			args: args{
				intervals:   [][]int{{3, 5}, {6, 9}},
				newInterval: []int{0, 1},
			},
			want: [][]int{{0, 1}, {3, 5}, {6, 9}},
		},
		{
			name: "Insert with no overlap at end",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}},
				newInterval: []int{6, 7},
			},
			want: [][]int{{1, 2}, {3, 5}, {6, 7}},
		},
		{
			name: "Insert between intervals",
			args: args{
				intervals:   [][]int{{1, 2}, {6, 9}},
				newInterval: []int{3, 5},
			},
			want: [][]int{{1, 2}, {3, 5}, {6, 9}},
		},
		{
			name: "Insert into empty intervals",
			args: args{
				intervals:   [][]int{},
				newInterval: []int{4, 8},
			},
			want: [][]int{{4, 8}},
		},
		{
			name: "Fully overlap all intervals",
			args: args{
				intervals:   [][]int{{2, 3}, {4, 5}, {6, 7}},
				newInterval: []int{1, 8},
			},
			want: [][]int{{1, 8}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
