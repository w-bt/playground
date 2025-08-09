package main

import "testing"

func TestClimbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n=1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "n=2",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "n=3",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "n=4",
			args: args{n: 4},
			want: 5,
		},
		{
			name: "n=5",
			args: args{n: 5},
			want: 8,
		},
		{
			name: "n=6",
			args: args{n: 6},
			want: 13,
		},
		{
			name: "n=10",
			args: args{n: 10},
			want: 89,
		},
		{
			name: "n=45",
			args: args{n: 45},
			want: 1836311903,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairs(tt.args.n); got != tt.want {
				t.Errorf("ClimbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
