package main

import "math"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	digit := getDigitSize(x)
	slice := splitToArray(x, digit)
	return compareArray(slice)
}

func getDigitSize(x int) int {
	counter := 1
	tempX := x
	for true {
		div := tempX / 10
		if div > 0 {
			counter += 1
			tempX = div
		} else {
			break
		}
	}

	return counter
}

func splitToArray(x, digit int) []int {
	result := []int{}
	tempX := x
	for i := digit; i > 0; i-- {
		div := tempX / int(math.Pow(10, float64(i-1)))
		remainder := tempX - (div * int(math.Pow(10, float64(i-1))))
		result = append(result, div)
		tempX = remainder
	}

	return result
}

func compareArray(slice []int) bool {
	median := len(slice) / 2

	for i := 0; i < median; i++ {
		if slice[i] != slice[len(slice)-(1+i)] {
			return false
		}
	}

	return true
}
