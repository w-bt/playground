package main

func ContainsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 2 || k <= 0 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums) && j-i <= k; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}
