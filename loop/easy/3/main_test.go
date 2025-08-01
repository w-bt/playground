package main

import "testing"

func TestFirstFactorial(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Factorial of 1",
			args: args{num: 1},
			want: 1,
		},
		{
			name: "Factorial of 2",
			args: args{num: 2},
			want: 2,
		},
		{
			name: "Factorial of 3",
			args: args{num: 3},
			want: 6,
		},
		{
			name: "Factorial of 4",
			args: args{num: 4},
			want: 24,
		},
		{
			name: "Factorial of 5",
			args: args{num: 5},
			want: 120,
		},
		{
			name: "Factorial of 10",
			args: args{num: 10},
			want: 3628800,
		},
		{
			name: "Factorial of 18 (upper bound)",
			args: args{num: 18},
			want: 6402373705728000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstFactorial(tt.args.num); got != tt.want {
				t.Errorf("FirstFactorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
