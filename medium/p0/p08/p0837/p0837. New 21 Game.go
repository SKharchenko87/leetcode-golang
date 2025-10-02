package p0837

import (
	"math"
)

const Precision = 1e5

func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 {
		return 1.0
	}
	dp := make([]float64, n+1)
	dp[0] = 1.0
	windowSum := dp[0]
	for i := 1; i <= n; i++ {
		dp[i] = windowSum / float64(maxPts)

		if i < k {
			windowSum += dp[i]
		}
		leftIndex := i - maxPts
		if leftIndex >= 0 && leftIndex < k {
			windowSum -= dp[leftIndex]
		}
	}
	var res float64
	for i := k; i <= n; i++ {
		res += dp[i]
	}
	return math.Round(res*Precision) / Precision
}

func new21Game0(n int, k int, maxPts int) float64 {
	dp := make([]float64, n+1)
	dp[0] = 1.0

	for i := 1; i <= n; i++ {
		for j := max(0, i-maxPts); j <= min(i-1, k-1); j++ {
			dp[i] += dp[j] / float64(maxPts)
		}
	}
	var res float64
	for i := k; i <= n; i++ {
		res += dp[i]
	}
	return math.Round(res*Precision) / Precision
}
