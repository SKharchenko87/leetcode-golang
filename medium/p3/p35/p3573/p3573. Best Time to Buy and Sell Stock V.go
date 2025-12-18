package p3573

// TLE
func maximumProfit(prices []int, k int) int64 {
	l := len(prices)
	dp := make([][]int64, l+1)
	for i := 0; i < l+1; i++ {
		dp[i] = make([]int64, k)
	}
	var ans int64
	for i := 0; i < l; i++ {
		index := i + 1
		for j := i + 1; j < l; j++ {
			jindex := j + 1
			d := int64(abs(prices[i] - prices[j]))
			dp[jindex][0] = max(d, dp[jindex][0], dp[index][0])
			ans = max(ans, dp[jindex][0])
			for m := 1; m < k; m++ {
				dp[jindex][m] = max(d+dp[index-1][m-1], dp[jindex][m], dp[index][m])
				ans = max(ans, dp[jindex][m])
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}
