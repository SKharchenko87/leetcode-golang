package p0070

func climbStairs(n int) int {
	cur, prev := 1, 1
	for i := 1; i < n; i++ {
		cur, prev = cur+prev, cur
	}
	return cur
}

func climbStairs0(n int) int {
	switch n {
	case 0, 1, 2, 3:
		return n
	}
	cur, prev := 3, 2
	for i := 3; i < n; i++ {
		cur, prev = cur+prev, cur
	}
	return cur
}

func climbStairs1(n int) int {
	dp := [46]int{}
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
