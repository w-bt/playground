package main

func ProductExceptSelf(nums []int) []int {
	product := 1
	productNonZero := 1
	zeroCounter := 0
	for i := 0; i < len(nums); i++ {
		product *= nums[i]
		if nums[i] == 0 && zeroCounter == 0 {
			zeroCounter += 1
			continue
		}
		productNonZero *= nums[i]
	}

	result := []int{}
	zeroCounter = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if zeroCounter == 0 {
				result = append(result, productNonZero)
				zeroCounter += 1
			} else {
				result = append(result, 0)
			}
			continue
		}
		result = append(result, product/nums[i])
	}
	return result
}
