package p3551

import "slices"

func minSwaps(nums []int) int {
	n := len(nums)
	indexes0 := make([]int, n)
	for i := 0; i < n; i++ {
		indexes0[i] = i
	}
	slices.SortFunc(indexes0, func(a, b int) int {
		sa, sb := sum(nums[a]), sum(nums[b])
		if sa == sb {
			return nums[a] - nums[b]
		}
		return sa - sb
	})
	res := 0
	for i := 0; i < n; i++ {
		for indexes0[i] != i {
			res++
			indexes0[indexes0[i]], indexes0[i] = indexes0[i], indexes0[indexes0[i]]
		}
	}
	return res
}

func sum(n int) int {
	res := 0
	for n > 0 {
		res += n % 10
		n /= 10
	}
	return res
}
