package main

func HammingWeight(num uint32) int {
	counter := 0
	for true {
		r := num % 2
		if r == 1 {
			counter += 1
		}
		num /= 2
		if num == 0 {
			break
		}
	}
	return counter
}
