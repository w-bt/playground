package main

func RemoveDuplicates(nums []int) int {
	uniqMap := make(map[int]bool)
	var index int
	for _, item := range nums {
		if _, ok := uniqMap[item]; !ok {
			nums[index] = item
			uniqMap[item] = true
			index++
		}
	}

	return index
}
