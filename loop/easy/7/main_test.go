package main

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1 row",
			args: args{
				numRows: 1,
			},
			want: [][]int{
				{1},
			},
		},
		{
			name: "2 rows",
			args: args{
				numRows: 2,
			},
			want: [][]int{
				{1},
				{1, 1},
			},
		},
		{
			name: "3 rows",
			args: args{
				numRows: 3,
			},
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
			},
		},
		{
			name: "5 rows",
			args: args{
				numRows: 5,
			},
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			name: "6 rows",
			args: args{
				numRows: 6,
			},
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
			},
		},
		{
			name: "8 rows",
			args: args{
				numRows: 8,
			},
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
				{1, 6, 15, 20, 15, 6, 1},
				{1, 7, 21, 35, 35, 21, 7, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Generate(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
