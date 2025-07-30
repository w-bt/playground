package main

import (
	"sort"
	"strconv"
)

func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	originInt := make(map[int]int)
	for _, num := range nums {
		originInt[num]++
	}

	uniqueHash := make(map[string]struct{})
	for i := 0; i < len(nums); i++ {
		copyMap := make(map[int]int)
		for k, v := range originInt {
			copyMap[k] = v
		}
		var triplet []int
		counter := 0

		if len(triplet) == 0 {
			counter = nums[i]
			copyMap[nums[i]]--
			triplet = append(triplet, nums[i])
		}

		for j := i + 1; j < len(nums); j++ {
			if len(triplet) == 3 {
				break
			}

			if len(triplet) == 1 {
				pair := counter + nums[j]
				missingValue := 0 - pair
				copyMap[nums[j]]--
				if v, ok := copyMap[missingValue]; ok && v > 0 && missingValue >= nums[j] {
					triplet = append(triplet, nums[j], missingValue)
					copyMap[missingValue]--

					sort.Ints(triplet)
					strHash := ""
					for _, v := range triplet {
						strHash += strconv.Itoa(v)
					}

					if _, ok := uniqueHash[strHash]; !ok {
						var validTriplet []int
						validTriplet = append(validTriplet, triplet...)
						result = append(result, validTriplet)
						uniqueHash[strHash] = struct{}{}
					}
					triplet = triplet[:1]
				} else {
					copyMap[nums[j]]++
				}
			}
		}
	}

	return result
}
