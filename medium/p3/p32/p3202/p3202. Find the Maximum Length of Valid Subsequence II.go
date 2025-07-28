package p3202

func maximumLength(nums []int, k int) int {
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, k)
	}
	res := 0
	for _, num := range nums {
		num %= k
		for prev := 0; prev < k; prev++ {
			dp[prev][num] = dp[num][prev] + 1
			res = max(res, dp[prev][num])
		}
	}
	return res
}
