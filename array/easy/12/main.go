package main

func RemoveElement(nums []int, val int) int {
	var index int
	for _, item := range nums {
		if item != val {
			nums[index] = item
			index++
		}
	}

	return index
}
