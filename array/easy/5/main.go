package main

func FindMaxConsecutiveOnes(nums []int) int {
	latestOnes := 0
	flagOnes := 0
	for _, num := range nums {
		if num == 1 {
			flagOnes++
			if flagOnes > latestOnes {
				latestOnes = flagOnes
			}
		} else {
			flagOnes = 0
		}
	}

	return latestOnes
}
