package main

func MoveZeroes(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	result := make([]int, len(nums))
	counter := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			result[counter] = nums[i]
			counter++
		}
	}

	return result
}
