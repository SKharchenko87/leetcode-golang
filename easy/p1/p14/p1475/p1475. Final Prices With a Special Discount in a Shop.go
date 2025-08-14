package p1475

func finalPrices(prices []int) []int {
	l := len(prices)
	for i := 0; i < l-1; i++ {
		j := i + 1
		for ; j < l && prices[i] < prices[j]; j++ {
		}
		if j < l && prices[i]-prices[j] >= 0 {
			prices[i] -= prices[j]
		}
	}
	return prices
}
