package main

func Merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 || len(nums2) == 0 {
		return
	}

	var dup []int
	dup = append(dup, nums1...)
	startM := 0
	startN := 0
	for {
		if startM < m {
			if startN < n {
				if dup[startM] <= nums2[startN] {
					nums1[startM+startN] = dup[startM]
					startM++
				} else {
					nums1[startM+startN] = nums2[startN]
					startN++
				}
			} else {
				nums1[startM+startN] = dup[startM]
				startM++
			}
		} else {
			if startN < n {
				nums1[startM+startN] = nums2[startN]
				startN++
			} else {
				break
			}
		}
	}
}
