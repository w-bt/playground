package main

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "num=1-true",
			args: args{num: 1},
			want: true,
		},
		{
			name: "num=2-false",
			args: args{num: 2},
			want: false,
		},
		{
			name: "num=4-true",
			args: args{num: 4},
			want: true,
		},
		{
			name: "just-below-9",
			args: args{num: 8},
			want: false,
		},
		{
			name: "num=9-true",
			args: args{num: 9},
			want: true,
		},
		{
			name: "just-above-9",
			args: args{num: 10},
			want: false,
		},
		{
			name: "num=16-true",
			args: args{num: 16},
			want: true,
		},
		{
			name: "just-below-25",
			args: args{num: 24},
			want: false,
		},
		{
			name: "num=25-true",
			args: args{num: 25},
			want: true,
		},
		{
			name: "large-true-808201",
			args: args{num: 808201}, // 899^2
			want: true,
		},
		{
			name: "large-false-808200",
			args: args{num: 808200},
			want: false,
		},
		{
			name: "edge-true-46340sq",
			args: args{num: 2147395600}, // 46340^2, batas aman int32
			want: true,
		},
		{
			name: "edge-below-false",
			args: args{num: 2147395599}, // 46340^2 - 1
			want: false,
		},
		{
			name: "edge-above-false",
			args: args{num: 2147395601}, // 46340^2 + 1
			want: false,
		},
		{
			name: "near-int32-max-false",
			args: args{num: 2147483647}, // bukan kuadrat sempurna
			want: false,
		},
		{
			name: "big-square-1e8-true",
			args: args{num: 100000000}, // 10000^2
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("IsPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
