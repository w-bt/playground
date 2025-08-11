package main

func IsPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	left, right := 1, num/2
	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid
		if square == num {
			return true
		} else if square < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
