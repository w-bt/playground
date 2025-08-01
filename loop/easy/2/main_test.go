package main

import (
	"reflect"
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	type args struct {
		candies    int
		num_people int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1 - candies = 7, num_people = 4",
			args: args{candies: 7, num_people: 4},
			want: []int{1, 2, 3, 1},
		},
		{
			name: "Example 2 - candies = 10, num_people = 3",
			args: args{candies: 10, num_people: 3},
			want: []int{5, 2, 3},
		},
		{
			name: "Single person",
			args: args{candies: 5, num_people: 1},
			want: []int{5},
		},
		{
			name: "Not enough for all in first round",
			args: args{candies: 2, num_people: 3},
			want: []int{1, 1, 0},
		},
		{
			name: "Even distribution",
			args: args{candies: 6, num_people: 3},
			want: []int{1, 2, 3},
		},
		{
			name: "One candy only",
			args: args{candies: 1, num_people: 4},
			want: []int{1, 0, 0, 0},
		},
		{
			name: "Edge case - candies = 0",
			args: args{candies: 0, num_people: 4},
			want: []int{0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistributeCandies(tt.args.candies, tt.args.num_people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
