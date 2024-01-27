package p0629

const mod int = 1e9 + 7

func kInversePairs(n int, k int) int {
	dp := [1001][1001]int{}
	for i := 1; i <= n; i++ {
		dp[i][0] = 1
		dp[i][1] = int(i - 1)
		j := 2
		for ; j < min(k+1, i); j++ {
			dp[i][j] = (dp[i-1][j] + dp[i][j-1] + mod) % mod
		}
		for ; j <= k; j++ {
			dp[i][j] = (dp[i-1][j] + dp[i][j-1] - dp[i-1][j-i] + mod) % mod
			if dp[i][j] == 0 {
				break
			}
		}
	}
	return dp[n][k]
}
