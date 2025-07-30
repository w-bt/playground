package main

func countBits(n int) []int {
	result := []int{}
	for i := 0; i <= n; i++ {
		result = append(result, getCounterBit(i))
	}
	return result
}

func getCounterBit(n int) int {
	counter := 0
	for true {
		if n == 0 {
			break
		}
		r := n % 2
		if r == 1 {
			counter += 1
		}
		n /= 2
	}
	return counter
}
