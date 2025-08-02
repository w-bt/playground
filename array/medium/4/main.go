package main

func MaxSubArray(nums []int) int {
	max := 0
	max_sum := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			max = nums[i]
			max_sum = nums[i]
			continue
		}
		if max+nums[i] > nums[i] {
			max = max + nums[i]
		} else {
			max = nums[i]
		}
		if max > max_sum {
			max_sum = max
		}
	}
	return max_sum
}
