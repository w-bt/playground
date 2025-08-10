package main

func MoveZeroes(nums []int) {
	nonZero := 0
	for i, _ := range nums {
		if nums[i] != 0 {
			nums[nonZero], nums[i] = nums[i], nums[nonZero]
			nonZero++
		}
	}
}
