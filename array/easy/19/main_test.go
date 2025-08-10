package main

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example case 1",
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			want: true,
		},
		{
			name: "Example case 2",
			args: args{
				nums: []int{1, 0, 1, 1},
				k:    1,
			},
			want: true,
		},
		{
			name: "Example case 3",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3},
				k:    2,
			},
			want: false,
		},
		{
			name: "No duplicates",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    3,
			},
			want: false,
		},
		{
			name: "Duplicates but too far apart",
			args: args{
				nums: []int{1, 2, 3, 4, 1},
				k:    2,
			},
			want: false,
		},
		{
			name: "Duplicates exactly k apart",
			args: args{
				nums: []int{99, 1, 2, 3, 99},
				k:    4,
			},
			want: true,
		},
		{
			name: "Single element array",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: false,
		},
		{
			name: "Multiple duplicates but only one within range",
			args: args{
				nums: []int{1, 2, 3, 2, 5, 1, 2},
				k:    2,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("ContainsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
