package p1477

import (
	"math"
)

func minSumOfLengths(arr []int, target int) int {
	dp := make([]int, len(arr)+1)
	dp[0] = math.MaxInt

	curSum := 0
	res := math.MaxInt
	for l, r := 0, 0; r < len(arr); r++ {
		curSum += arr[r]
		for curSum > target {
			curSum -= arr[l]
			l++
		}

		dp[r+1] = dp[r]
		if curSum == target {
			if dp[l] != math.MaxInt {
				res = min(res, dp[l]+r-l+1)
			}
			dp[r+1] = min(dp[r+1], r-l+1)
		}
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}

func minSumOfLengths0(arr []int, target int) int {
	n := len(arr)
	prevSum := make(map[int]int, n)
	prevSum[0] = -1
	dp := make([]int, n+1)
	dp[0] = math.MaxInt
	curSum := 0
	res := math.MaxInt
	for i := 0; i < n; i++ {
		curSum += arr[i]
		prevSum[curSum] = i
		dp[i+1] = dp[i]
		if v, ok := prevSum[curSum-target]; ok {
			if dp[v+1] != math.MaxInt {
				res = min(res, i-v+dp[v+1])
			}
			dp[i+1] = min(dp[i+1], i-v)
		}
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}
