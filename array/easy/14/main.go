package main

func SearchInsert(nums []int, target int) int {
	return searchRecursive(nums, target, 0, len(nums))
}

func searchRecursive(nums []int, target int, left int, right int) int {
	if left > right {
		return left
	}

	mid := (left + right) / 2
	if mid == len(nums) {
		return mid
	}

	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return searchRecursive(nums, target, mid+1, right)
	} else {
		return searchRecursive(nums, target, left, mid-1)
	}
}
