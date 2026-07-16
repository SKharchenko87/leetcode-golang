package p3867

import "sort"

func gcdSum(nums []int) int64 {
	n := len(nums)
	prefixGCD := make([]int, n)
	prefixGCD[0] = nums[0]
	mx := nums[0]
	for i := 1; i < n; i++ {
		mx = max(mx, nums[i])
		prefixGCD[i] = gcd(mx, nums[i])
	}
	sort.Ints(prefixGCD)
	var res int64
	for i := 0; i < n/2; i++ {
		res += int64(gcd(prefixGCD[i], prefixGCD[n-1-i]))
	}
	return res
}

func gcd(a, b int) int {
	a, b = max(a, b), min(a, b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
