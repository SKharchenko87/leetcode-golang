package p3423

func maxAdjacentDistance(nums []int) int {
	l := len(nums)
	res := abs(nums[0] - nums[l-1])
	for i := 0; i < l-1; i++ {
		res = max(res, abs(nums[i]-nums[i+1]))
	}
	return res
}

func abs[T int](x T) T {
	if x >= 0 {
		return x
	}
	return -1 * x
}
