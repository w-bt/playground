package main

import "fmt"

// FindMax returns the maximum value in the slice.
// If the slice is empty, it returns an error.
func FindMax(input []int) (int, error) {
	if len(input) == 0 {
		return 0, fmt.Errorf("empty array")
	}
	max := input[0]
	for _, i := range input {
		if i > max {
			max = i
		}
	}
	return max, nil
}
