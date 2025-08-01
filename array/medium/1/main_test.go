package main

import "testing"

func TestMaxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
		{
			name: "Example 2",
			args: args{height: []int{1, 1}},
			want: 1,
		},
		{
			name: "Decreasing heights",
			args: args{height: []int{6, 5, 4, 3, 2, 1}},
			want: 9, // between index 0 and 3 (min height 3 * width 3)
		},
		{
			name: "All same height",
			args: args{height: []int{5, 5, 5, 5, 5}},
			want: 20, // between index 0 and 4
		},
		{
			name: "Peak in the middle",
			args: args{height: []int{1, 3, 5, 9, 5, 3, 1}},
			want: 12, // between index 1 and 5 (min height 3 * width 4)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxArea(tt.args.height); got != tt.want {
				t.Errorf("MaxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
