package main

func MaxArea(height []int) int {
	maxArea := 0
	for i := 0; i < len(height); i++ {
		if height[i] == 0 {
			continue
		}
		h := height[i]
		minLen := 1
		if maxArea > 0 {
			minLen = maxArea / height[i]
			if minLen == 0 {
				minLen = 1
			}
		}
		for j := i + minLen; j < len(height); j++ {
			if height[j] == 0 {
				continue
			}
			if height[j] <= height[i] {
				h = height[j]
			}
			area := h * (j - i)
			if area > maxArea {
				maxArea = area
			}
			h = height[i]
		}
	}
	return maxArea
}
