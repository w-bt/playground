package main

func MajorityElement(nums []int) int {
	mapHash := make(map[int]int)
	maxCount := 0
	maxNum := 0
	for i := 0; i < len(nums); i++ {
		mapHash[nums[i]]++
		if mapHash[nums[i]] > maxCount {
			maxCount = mapHash[nums[i]]
			maxNum = nums[i]
		}
	}

	return maxNum
}
