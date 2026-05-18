package p2770

func maximumJumps(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[n-1] = -1
	for j := 1; j < n; j++ {
		if diff := nums[j] - nums[0]; -target <= diff && diff <= target {
			dp[j] = 1
		}
	}
	for i := 1; i < n-1; i++ {
		if dp[i] > 0 {
			for j := i + 1; j < n; j++ {
				if diff := nums[j] - nums[i]; -target <= diff && diff <= target {
					dp[j] = max(dp[j], dp[i]+1)
				}
			}
		}
	}
	return dp[n-1]
}
