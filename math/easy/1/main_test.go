package main

import "testing"

func TestCanWinNim(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "n=1 (win)",
			args: args{n: 1},
			want: true,
		},
		{
			name: "n=2 (win)",
			args: args{n: 2},
			want: true,
		},
		{
			name: "n=3 (win)",
			args: args{n: 3},
			want: true,
		},
		{
			name: "n=4 (lose)",
			args: args{n: 4},
			want: false,
		},
		{
			name: "n=5 (win)",
			args: args{n: 5},
			want: true,
		},
		{
			name: "n=7 (win)",
			args: args{n: 7},
			want: true,
		},
		{
			name: "n=8 (lose)",
			args: args{n: 8},
			want: false,
		},
		{
			name: "n=12 (lose)",
			args: args{n: 12},
			want: false,
		},
		{
			name: "n=15 (win)",
			args: args{n: 15},
			want: true,
		},
		{
			name: "n=16 (lose)",
			args: args{n: 16},
			want: false,
		},
		{
			name: "n=21 (win)",
			args: args{n: 21},
			want: true,
		},
		{
			name: "n=2147483644 (lose, multiple of 4)",
			args: args{n: 2147483644},
			want: false,
		},
		{
			name: "n=2147483647 (win, max int32)",
			args: args{n: 2147483647},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanWinNim(tt.args.n); got != tt.want {
				t.Errorf("CanWinNim() = %v, want %v", got, tt.want)
			}
		})
	}
}
