package p3652

func maxProfit(prices []int, strategy []int, k int) int64 {
	l := len(prices)
	var sum int64
	for i := 0; i < l; i++ {
		sum += int64(strategy[i] * prices[i])
	}
	var window int64
	i := 0
	for ; i < k/2; i++ {
		window += int64(-strategy[i] * prices[i])
	}
	for ; i < k; i++ {
		window += int64((1 - strategy[i]) * prices[i])
	}
	var result int64
	result = max(sum, sum+window)
	for ; i < l; i++ {
		window += int64(strategy[i-k] * prices[i-k])
		window += int64(-1 * prices[i-k/2])
		window += int64((1 - strategy[i]) * prices[i])
		result = max(result, sum+window)
	}
	return result
}
