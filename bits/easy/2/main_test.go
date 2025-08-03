package main

import "testing"

func TestHammingWeight(t *testing.T) {
	type args struct {
		num uint32
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
			name: "One bit set",
			args: args{num: 1},
			want: 1,
		},
		{
			name: "All bits set",
			args: args{num: 0xFFFFFFFF},
			want: 32,
		},
		{
			name: "Multiple bits set",
			args: args{num: 0b101010},
			want: 3,
		},
		{
			name: "High bit set",
			args: args{num: 1 << 31},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HammingWeight(tt.args.num); got != tt.want {
				t.Errorf("HammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
