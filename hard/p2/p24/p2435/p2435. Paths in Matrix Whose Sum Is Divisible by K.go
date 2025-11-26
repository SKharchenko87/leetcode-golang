package p2435

const MOD = 1e9 + 7

func numberOfPaths(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, m+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, k)
		}
	}

	dp[1][1][grid[0][0]%k] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for d := 0; d < k; d++ {
				p := (d + grid[i][j]) % k
				dp[i+1][j+1][p] += (dp[i][j+1][d] + dp[i+1][j][d]) % MOD
			}
		}
	}
	return dp[m][n][0]
}
