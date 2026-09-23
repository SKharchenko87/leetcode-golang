package p1658

import (
	"math"
	"slices"
)

func minOperations(nums []int, x int) int {
	n := len(nums)
	sumR := 0
	j := n - 1
	for ; j >= 0 && sumR < x; j-- {
		sumR += nums[j]
	}
	j++
	if sumR < x {
		return -1
	}
	res := math.MaxInt
	sumL := 0
	for i := 0; i < n; i++ {
		for j < n && sumL+sumR > x {
			sumR -= nums[j]
			j++
		}
		if sumL+sumR == x {
			res = min(res, i+(n-j))
		}
		sumL += nums[i]
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}

func minOperations0(nums []int, x int) int {
	n := len(nums)
	rl := make([]int, n)
	sum := 0
	i := 0
	for _, num := range slices.Backward(nums) {
		sum += num
		rl[i] = sum
		i++
	}
	j, ok := slices.BinarySearch(rl, x)
	res := math.MaxInt
	if ok {
		res = j + 1
	}
	sum = 0
	i = 0
	j = min(j, n-1)
	j--
	for ; i < n-1 && sum < x; i++ {
		sum += nums[i]
		if sum == x {
			res = min(res, i+1)
		}
		for j > 0 && (sum+rl[j] > x || i+j >= n-1) {
			j--
		}
		if j >= 0 && sum+rl[j] == x {
			res = min(res, i+j+2)
		}
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}
