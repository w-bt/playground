package main

func Intersection(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		hashMap[nums1[i]]++
	}
	result := make([]int, 0)
	resultMap := make(map[int]bool)
	for i := 0; i < len(nums2); i++ {
		if _, ok := hashMap[nums2[i]]; ok && hashMap[nums2[i]] > 0 {
			hashMap[nums2[i]]--
			if _, ok := resultMap[nums2[i]]; !ok {
				result = append(result, nums2[i])
				resultMap[nums2[i]] = true
			}
		}
	}
	return result
}
