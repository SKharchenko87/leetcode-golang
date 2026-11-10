package p3693

func climbStairs(n int, costs []int) int {
	dp := make([]int, n+1)
	dp[0] = costs[0] + 1
	for j := 1; j <= 3 && j <= n; j++ {
		dp[j-1] = costs[j-1] + j*j
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= 3 && i-j >= 0; j++ {
			v := dp[i-j] + costs[i] + j*j
			if dp[i] == 0 || dp[i] > v {
				dp[i] = v
			}
		}
	}
	return dp[n-1]
}
