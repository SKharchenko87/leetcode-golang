package p3578

import "math"

/*ToDo*/
func countPartitions(nums []int, k int) int {
	l := len(nums)
	rl := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		mn := math.MaxInt
		mx := math.MinInt
		rl[i]++
		for j := i - 1; j > 0; j-- {
			mn = min(mn, nums[j])
			mx = max(mx, nums[j])
			if mx-mn <= k {
				rl[j]++
			} else {
				break
			}
		}
		rl[i]--
	}
	res := 0
	for i := 0; i < l; i++ {
		res += rl[i]
	}
	//for i := 0; i < l; i++ {
	//	mn := math.MaxInt
	//	mx := math.MinInt
	//
	//	for j := 0; j < l-1; j++ {
	//		mn = min(mn, nums[j])
	//		mx = max(mx, nums[j])
	//		if mx-mn >= k {
	//			res += rl[j+1]
	//		}
	//	}
	//}
	return res
}
