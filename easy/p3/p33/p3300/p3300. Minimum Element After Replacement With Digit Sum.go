package p3300

import "math"

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
