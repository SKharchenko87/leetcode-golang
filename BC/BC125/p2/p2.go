package p2

import (
	"fmt"
	"sort"
)

func minOperations(nums []int, k int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	sum := 0
	cnt := 0
	for _, v := range nums {
		if v < k {
			sum += v
			cnt++
		} else {
			break
		}
	}
	if sum*2 < k {
		return cnt / 2
	} else {
		return cnt / 2
	}
}

func minOperations1(nums []int, k int) int {
	cnt := 0
	for nums[0] < k && len(nums) >= 2 {
		fmt.Println(nums)
		nums[1] = 2*min(nums[0], nums[1]) + max(nums[0], nums[1])
		nums = nums[1:]
		sort.Ints(nums)
		cnt++
	}
	return cnt
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
