package p2

func longestSubsequence(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}

	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = 1
	}

	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			a := abs(nums[i] - nums[j])
			if j == 0 || a <= abs(nums[j]-nums[j-1]) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 1
	for _, d := range dp {
		res = max(res, d)
	}

	return res
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
