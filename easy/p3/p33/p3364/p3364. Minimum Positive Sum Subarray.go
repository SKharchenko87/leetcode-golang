package p3364

import "math"

func minimumSumSubarray(nums []int, l, r int) int {
	n := len(nums)
	prevSum := make([]int, n+1)
	for i, num := range nums {
		prevSum[i+1] = prevSum[i] + num
	}
	res := math.MaxInt
	for size := l - 1; size < r; size++ {
		for i := size; i < n; i++ {
			x := prevSum[i+1] - prevSum[i-size]
			if x > 0 && x < res {
				res = x
			}
		}
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}
