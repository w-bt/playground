package main

import "testing"

func TestMaxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1 - Profitable",
			args: args{prices: []int{7, 1, 5, 3, 6, 4}},
			want: 5,
		},
		{
			name: "Example 2 - No Profit",
			args: args{prices: []int{7, 6, 4, 3, 1}},
			want: 0,
		},
		{
			name: "Single Day - No Transaction",
			args: args{prices: []int{5}},
			want: 0,
		},
		{
			name: "Two Days - Profit",
			args: args{prices: []int{2, 4}},
			want: 2,
		},
		{
			name: "Two Days - No Profit",
			args: args{prices: []int{4, 2}},
			want: 0,
		},
		{
			name: "Multiple Dips and Peaks",
			args: args{prices: []int{3, 2, 6, 1, 3}},
			want: 4,
		},
		{
			name: "Constant Prices",
			args: args{prices: []int{3, 3, 3, 3}},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.args.prices); got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
