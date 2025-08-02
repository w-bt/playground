package main

func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSubArr := 1
	maxProduct := -10
	rawProduct := 0
	index := 0
	for {
		iteratorArr := 0
		for j := index; j < len(nums); j++ {
			if iteratorArr == 0 {
				rawProduct = nums[j]
				maxProduct = max(maxProduct, rawProduct)
				iteratorArr++
			} else if iteratorArr > 0 && iteratorArr < maxSubArr {
				rawProduct *= nums[j]
				iteratorArr++
			} else if iteratorArr >= maxSubArr {
				maxProduct = max(maxProduct, rawProduct)
				rawProduct = 0
				iteratorArr = 0
				break
			}
		}
		if rawProduct > maxProduct {
			maxProduct = rawProduct
			rawProduct = 0
		}
		index++
		if maxSubArr <= len(nums) && index+maxSubArr > len(nums) {
			maxSubArr++
			index = 0
		} else if maxSubArr >= len(nums) {
			break
		}
	}

	return maxProduct
}
