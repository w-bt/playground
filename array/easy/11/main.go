package main

func TwoSum(nums []int, target int) []int {
	for i := range nums {
		stackIndex := []int{i}
		total := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if total+nums[j] == target {
				stackIndex = append(stackIndex, j)
				return stackIndex
			}
		}
	}
	return []int{}
}
