package p0121

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return maxProfit
}

// 7, 1,5,3,6,4
// 0 -6 4 2 5 3
// 7  1 1 1 1 1
