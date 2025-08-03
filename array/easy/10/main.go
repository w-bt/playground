package main

func ArraySign(nums []int) int {
	negativeCounter := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			negativeCounter += 1
		}
	}
	if negativeCounter%2 == 0 {
		return 1
	}
	return -1
}
