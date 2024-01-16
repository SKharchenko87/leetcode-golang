package p0714

// ToDo понять как это работает
func maxProfit(prices []int, fee int) int {
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(prices[i]+dp[i-1][1]-fee, dp[i-1][0])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}

	return dp[len(prices)-1][0]
}
