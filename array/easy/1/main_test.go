package main

import (
	"testing"
)

func TestFindMax(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		wantErr  bool
	}{
		{"mixed positive", []int{3, 5, 1, 7, 2}, 7, false},
		{"all negative", []int{-10, -3, -7, -1}, -1, false},
		{"single positive", []int{5}, 5, false},
		{"single value", []int{42}, 42, false},
		{"ascending", []int{1, 2, 3, 4, 5}, 5, false},
		{"negative only", []int{-10, -5, -20}, -5, false},
		{"mixed", []int{0, -1, 10, 3, -5}, 10, false},
		{"all same", []int{7, 7, 7, 7}, 7, false},
		{"empty", []int{}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindMax(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMax() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("FindMax() = %v, want %v", got, tt.expected)
			}
		})
	}
}
