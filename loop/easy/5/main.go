package main

func MySqrt(x int) int {
	if x < 2 {
		return x
	}
	left, right := 1, x/2
	result := 0
	for left <= right {
		mid := left + (right-left)/2
		if mid <= x/mid {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
