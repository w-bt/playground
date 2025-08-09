package main

func PlusOne(digits []int) []int {
	remainder := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+remainder > 9 {
			digits[i] = 0
			remainder = 1
			continue
		}

		digits[i] += remainder
		remainder = 0
	}

	if remainder > 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
