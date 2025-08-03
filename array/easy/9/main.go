package main

func MissingNumber(nums []int) int {
	flags := make([]bool, len(nums)+1)
	for _, num := range nums {
		flags[num] = true
	}
	for index, _ := range flags {
		if !flags[index] {
			return index
		}
	}

	return 0
}
