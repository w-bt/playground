package main

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "n = 0",
			args: args{n: 0},
			want: []int{0},
		},
		{
			name: "n = 1",
			args: args{n: 1},
			want: []int{0, 1},
		},
		{
			name: "n = 2",
			args: args{n: 2},
			want: []int{0, 1, 1},
		},
		{
			name: "n = 5",
			args: args{n: 5},
			want: []int{0, 1, 1, 2, 1, 2},
		},
		{
			name: "n = 10",
			args: args{n: 10},
			want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
