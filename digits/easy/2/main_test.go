package main

import "testing"

func TestAddDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Zero",
			args: args{num: 0},
			want: 0,
		},
		{
			name: "Single digit",
			args: args{num: 7},
			want: 7,
		},
		{
			name: "Two digits sum once",
			args: args{num: 10},
			want: 1, // 1+0=1
		},
		{
			name: "Two digits sum twice",
			args: args{num: 38},
			want: 2, // 3+8=11 -> 1+1=2
		},
		{
			name: "Multiple nines",
			args: args{num: 99},
			want: 9, // 9+9=18 -> 1+8=9
		},
		{
			name: "Large number",
			args: args{num: 2147483647},
			want: 1, // Digital root formula
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddDigits(tt.args.num); got != tt.want {
				t.Errorf("AddDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
