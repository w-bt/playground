package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "n = 1",
			n:    1,
			want: []string{"1"},
		},
		{
			name: "n = 3",
			n:    3,
			want: []string{"1", "2", "Fizz"},
		},
		{
			name: "n = 5",
			n:    5,
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name: "n = 15",
			n:    15,
			want: []string{
				"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
				"11", "Fizz", "13", "14", "FizzBuzz",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := FizzBuzz(tt.n); !reflect.DeepEqual(gotResult, tt.want) {
				t.Errorf("FizzBuzz() = %v, want %v", gotResult, tt.want)
			}
		})
	}
}
