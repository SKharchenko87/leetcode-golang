package p3026

import "math"

func maximumSubarraySum(nums []int, k int) int64 {
	l := len(nums)
	m := make(map[int][]int, l)
	dp := make([]int64, l+1)
	var res int64 = math.MinInt64
	for i, v := range nums {
		dp[i+1] = dp[i] + int64(v)
		if _, ok := m[v]; ok {
			m[v] = append(m[v], i)
		} else {
			m[v] = []int{i}
		}
	}
	for key, arr0 := range m {
		if arr1, ok := m[key-k]; ok {
			for _, index1 := range arr1 {
				for _, index0 := range arr0 {
					start, end := min(index0, index1), max(index0, index1)+1
					if res < dp[end]-dp[start] {
						res = dp[end] - dp[start]
					}
				}
			}
		}
	}
	if res == math.MinInt64 {
		return 0
	}
	return res
}

func abs(x int64) int64 {
	if x < 0 {
		return -1 * x
	}
	return x
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
