package main

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Simple increment without carry",
			args: args{
				digits: []int{1, 2, 3},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "Single digit without carry",
			args: args{
				digits: []int{4},
			},
			want: []int{5},
		},
		{
			name: "Single digit with carry",
			args: args{
				digits: []int{9},
			},
			want: []int{1, 0},
		},
		{
			name: "Multiple digits with carry at last",
			args: args{
				digits: []int{1, 2, 9},
			},
			want: []int{1, 3, 0},
		},
		{
			name: "Multiple digits all 9",
			args: args{
				digits: []int{9, 9, 9},
			},
			want: []int{1, 0, 0, 0},
		},
		{
			name: "Carry in middle",
			args: args{
				digits: []int{2, 9, 9},
			},
			want: []int{3, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
