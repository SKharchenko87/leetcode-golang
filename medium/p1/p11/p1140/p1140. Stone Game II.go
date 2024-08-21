package p1140

import "math"

func stoneGameII(piles []int) int {
	l := len(piles)
	dp := make([][]int, l+1)
	for i := 0; i < l+1; i++ {
		dp[i] = make([]int, l+1)
	}
	suffixSum := make([]int, l+1)
	for i := l - 1; i >= 0; i-- {
		suffixSum[i] += piles[i] + suffixSum[i+1]
	}
	for i := 0; i < l+1; i++ {
		dp[i][l] = suffixSum[i]
	}
	for index := l - 1; index >= 0; index-- {
		for maxTillNow := l - 1; maxTillNow > 0; maxTillNow-- {
			for x := 1; x < min(2*maxTillNow, l-index)+1; x++ {
				dp[index][maxTillNow] = max(dp[index][maxTillNow],
					suffixSum[index]-dp[index+x][max(maxTillNow, x)],
				)
			}
		}
	}
	return dp[0][1]
}

func stoneGameII0(piles []int) int {
	l := len(piles)
	memo := make([][]int, l)
	for i := 0; i < l; i++ {
		memo[i] = make([]int, l)
	}
	suffixSum := make([]int, l)
	suffixSum[l-1] = piles[l-1]
	for i := l - 2; i >= 0; i-- {
		suffixSum[i] += piles[i] + suffixSum[i+1]
	}
	var maxStone func(maxTillNow int, currIndex int) int
	maxStone = func(maxTillNow int, currIndex int) int {
		if currIndex+2*maxTillNow >= l {
			return suffixSum[currIndex]
		}
		if memo[currIndex][maxTillNow] > 0 {
			return memo[currIndex][maxTillNow]
		}
		res := math.MaxInt
		for i := 1; i < 2*maxTillNow+1; i++ {
			res = min(res, maxStone(max(i, maxTillNow), currIndex+i))
		}
		memo[currIndex][maxTillNow] = suffixSum[currIndex] - res
		return memo[currIndex][maxTillNow]
	}
	return maxStone(1, 0)
}
