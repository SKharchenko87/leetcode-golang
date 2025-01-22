package p1

import "sort"

func minOperations(nums []int, k int) int {
	sort.Ints(nums)
	res := 0
	if k > nums[0] {
		return -1
	} else if k < nums[0] {
		res++
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			res++
		}
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
