package p0446

func numberOfArithmeticSlices(nums []int) int {
	l := len(nums)
	cnt := 0
	dp := make([]map[int]int, l)
	for i := 0; i < l; i++ {
		dp[i] = map[int]int{}
	}
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			dp[i][diff]++
			if _, ok := dp[i][diff]; ok {
				dp[i][diff] += dp[j][diff]
				cnt += dp[j][diff]
			}
		}
	}
	return cnt
}
