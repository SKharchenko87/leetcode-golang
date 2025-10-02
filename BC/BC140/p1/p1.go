package p1

import "math"

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

func minElement(nums []int) int {
	res := math.MaxInt
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		sum := 0
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		if sum < res {
			res = sum
		}
	}
	return res
}
