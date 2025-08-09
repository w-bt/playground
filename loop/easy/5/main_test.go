package main

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"Zero", 0, 0},
		{"One", 1, 1},
		{"Perfect square", 4, 2},
		{"Non-perfect square", 8, 2},
		{"Large square", 10000, 100},
		{"Between squares", 15, 3},
		{"Max int", 2147395600, 46340},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MySqrt(tt.x)
			if got != tt.want {
				t.Errorf("mySqrt(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}
