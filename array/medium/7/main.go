package main

func Search(nums []int, target int) int {
	return binSearch(nums, target, 0)
}

func binSearch(nums []int, target int, leftIndex int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return leftIndex
		}
		return -1
	}
	if nums[0] < nums[len(nums)-1] && (target > nums[len(nums)-1] || target < nums[0]) {
		return -1
	}
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]
	index := binSearch(left, target, leftIndex)
	if index > -1 {
		return index
	}
	index = binSearch(right, target, leftIndex+mid)
	if index > -1 {
		return index
	}

	return -1
}
