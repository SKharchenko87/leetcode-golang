package p1531

func getLengthOfOptimalCompression(s string, K int) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = n
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for k := 0; k <= K; k++ {
			if k > 0 { // delete s[i]
				dp[i][k] = Min(dp[i][k], dp[i-1][k-1])
			}
			// keep s[i] concate the same, remove the different
			same := 0
			diff := 0
			for j := i; j >= 1; j-- {
				if s[j-1] == s[i-1] {
					same++
				} else {
					diff++
				}

				if diff > k {
					break
				}
				// // baaaaaaccaaaa  k=2, s[i] = a, j will extend to index 1
				dp[i][k] = Min(dp[i][k], dp[j-1][k-diff]+getLen(same))
			}
		}
	}
	return dp[n][K]
}

func getLen(length int) int {
	if length == 1 {
		return 1
	} else if length < 10 {
		return 2
	} else if length < 100 {
		return 3
	}
	return 4
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
