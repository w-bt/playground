package main

import "fmt"

func SummaryRanges(nums []int) []string {
	result := []string{}
	allRange := [][]int{}
	resultRange := []int{}
	for i := 0; i < len(nums); i++ {
		if len(resultRange) == 0 {
			resultRange = append(resultRange, nums[i])
		} else {
			lastNumber := resultRange[len(resultRange)-1]
			if nums[i] == lastNumber+1 {
				if len(resultRange) == 1 {
					resultRange = append(resultRange, nums[i])
				} else {
					resultRange[len(resultRange)-1] = nums[i]
				}
			} else {
				allRange = append(allRange, resultRange)
				resultRange = []int{nums[i]}
			}
		}
	}

	if len(resultRange) > 0 {
		allRange = append(allRange, resultRange)
	}

	for i := 0; i < len(allRange); i++ {
		if len(allRange[i]) == 1 {
			result = append(result, fmt.Sprintf("%d", allRange[i][0]))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", allRange[i][0], allRange[i][1]))
		}
	}

	return result
}
