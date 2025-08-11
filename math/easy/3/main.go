package main

func IsPowerOfThree(n int) bool {
	const maxPow3 = 1162261467
	return n > 0 && maxPow3%n == 0
}
