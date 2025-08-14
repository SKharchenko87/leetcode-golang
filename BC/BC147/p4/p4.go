package p4

import (
	"math"
)

func maxSubarraySum(nums []int) int64 {
	l := len(nums)
	minus := make(map[int]struct{}, l)
	minus[0] = struct{}{}
	for i := 0; i < l; i++ {
		if nums[i] < 0 {
			minus[nums[i]] = struct{}{}
		}
	}
	bestSum := int64(math.MinInt64)
	for m := range minus {
		curSum := int64(0)
		maxSum := int64(math.MinInt64)
		for _, n := range nums {
			if n != m {
				curSum = max64(int64(n), curSum+int64(n))
				maxSum = max64(maxSum, curSum)
			}
		}
		if bestSum < maxSum {
			bestSum = maxSum
		}
	}

	return bestSum
}

func max64(x, y int64) int64 {
	if x > y {
		return int64(x)
	}
	return int64(y)
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

//-3 -1 -3 -4 -1 -3 0
//    2  0 -1  2  0 3
//-3 -1    -2  1    4
// -3, 2, -2, -1, 3, -2, 3
// 0  2    0  -1  3   1  4
// 0  2    2   2  2   2  4

// -3, 2, -1, 3, 3
// 0   2   1  4  7
// 0   2   2  4  7
