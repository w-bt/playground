package main

// ClimbStairs calculates the number of distinct ways to climb to the top of a staircase with n steps.
// Can be solved using Fibonacci.
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
