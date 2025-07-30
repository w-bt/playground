package main

func MergeSortedArrays(nums1 []int, nums2 []int) (result []int) {
	if len(nums1) == 0 && len(nums2) == 0 {
		return []int{}
	} else if len(nums1) == 0 {
		return nums2
	} else if len(nums2) == 0 {
		return nums1
	}

	pointerFirst := 0
	pointerSecond := 0
	for {
		if len(nums2) <= pointerSecond && len(nums1) <= pointerFirst {
			break
		} else if len(nums2) <= pointerSecond && pointerFirst < len(nums1) {
			result = append(result, nums1[pointerFirst])
			pointerFirst++
		} else if len(nums1) <= pointerFirst && pointerSecond < len(nums2) {
			result = append(result, nums2[pointerSecond])
			pointerSecond++
		} else {
			if nums2[pointerSecond] > nums1[pointerFirst] {
				result = append(result, nums1[pointerFirst])
				pointerFirst++
			} else if nums2[pointerSecond] <= nums1[pointerFirst] {
				result = append(result, nums2[pointerSecond])
				pointerSecond++
			}
		}
	}

	return result
}
