package main

func SingleNumber(nums []int) int {
	mapHash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := mapHash[nums[i]]; ok {
			delete(mapHash, nums[i])
		} else {
			mapHash[nums[i]] = 1
		}
	}

	for key, _ := range mapHash {
		return key
	}

	return 0
}
