package p1937

import (
	"math"
	"slices"
)

func maxPoints(points [][]int) int64 {
	m, n := len(points), len(points[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		dp[0] = points[i][0] + dp[0]
		for j := 1; j < n; j++ {
			dp[j] = max(points[i][j]+dp[j], dp[j-1]-1)
		}
		for j := n - 2; j >= 0; j-- {
			dp[j] = max(dp[j], dp[j+1]-1)
		}
	}
	return int64(slices.Max(dp))
}

func maxPoints3(points [][]int) int64 {
	m, n := len(points), len(points[0])
	dp := make([]int64, n)
	for i := 0; i < m; i++ {
		dp[0] = int64(points[i][0]) + dp[0]
		for j := 1; j < n; j++ {
			dp[j] = max64(int64(points[i][j])+dp[j], dp[j-1]-1)
		}
		for j := n - 2; j >= 0; j-- {
			dp[j] = max64(dp[j], dp[j+1]-1)
		}
	}
	return slices.Max(dp)
}

func maxPoints2(points [][]int) int64 {
	m, n := len(points), len(points[0])
	dp := make([][]int64, m)
	getUp := func(i, j int) int64 {
		if i > 0 {
			return dp[i-1][j]
		} else {
			return 0
		}
	}
	for i := 0; i < m; i++ {
		dp[i] = make([]int64, n)
		up := getUp(i, 0)
		dp[i][0] = int64(points[i][0]) + up
		for j := 1; j < n; j++ {
			up = getUp(i, j)
			dp[i][j] = max64(int64(points[i][j])+up, dp[i][j-1]-1)
		}
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = max64(dp[i][j], dp[i][j+1]-1)
		}
	}
	return slices.Max(dp[m-1])
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func maxPoints1(points [][]int) int64 {
	m, n := len(points), len(points[0])
	dp := make([][]int64, m)
	dp[0] = make([]int64, n)
	for j := 0; j < n; j++ {
		dp[0][j] = int64(points[0][j])
	}
	for i := 1; i < m; i++ {
		dp[i] = make([]int64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = int64(points[i][j])
			for k := 0; k < n; k++ {
				if int64(points[i][j])+dp[i-1][k]-int64(abs(j-k)) > dp[i][j] {
					dp[i][j] = int64(points[i][j]) + dp[i-1][k] - int64(abs(j-k))
				}
			}
		}
	}
	return slices.Max(dp[m-1])
}

func maxPoints0(points [][]int) int64 {
	m, n := len(points), len(points[0])
	maxTmp := points[0][0]
	for j := 1; j < n; j++ {
		if maxTmp < points[0][j] {
			maxTmp = points[0][j]
		}
	}
	for i := 1; i < m; i++ {
		maxTmp = math.MinInt
		for j := 0; j < n; j++ {
			tmp := points[i-1][0] - j
			for k := 1; k < n; k++ {
				x := points[i-1][k] - abs(j-k)
				if x > tmp {
					tmp = x
				}
			}
			points[i][j] += tmp
			if points[i][j] > maxTmp {
				maxTmp = points[i][j]
			}
		}
	}
	return int64(maxTmp)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}
