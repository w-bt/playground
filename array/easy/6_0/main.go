package main

func MaxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		minPrice := 0
		if profit > 0 {
			minPrice = prices[i] + profit
		}
		maxNumber := 0
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				break
			}
			if prices[j] > maxNumber {
				maxNumber = prices[j]
			}
			if (maxNumber > 0 && prices[j] < maxNumber) || (minPrice > 0 && prices[j] < minPrice) {
				continue
			}
			prof := prices[j] - prices[i]
			if prof > profit {
				profit = prof
			}
		}
	}
	return profit
}
