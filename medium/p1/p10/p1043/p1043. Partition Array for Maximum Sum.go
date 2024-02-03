package p1043

import "fmt"

/*
    1, 15,  7,  9,  2,  5, 10
0,  1, 30, 45, 54, 63, 72, 84 - лучший варинат
       16, 37, 46, 57, 64, 83
           31, 48, 48, 68, 82
*/

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		maxVal := 0
		for j := 1; j <= min(i, k); j++ {
			maxVal = max(maxVal, arr[i-j])
			dp[i] = max(dp[i], dp[i-j]+maxVal*j)
		}
	}
	fmt.Println(dp)
	return dp[n]
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
