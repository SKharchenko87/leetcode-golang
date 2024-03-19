package p0930

func sumAtMost(nums []int, goal int) int {
	if goal < 0 {
		return 0
	}

	res, prev, cur := 0, 0, 0
	for i, val := range nums {
		cur += val
		for cur > goal {
			cur -= nums[prev]
			prev++
		}
		res += i - prev + 1
	}

	return res
}

func numSubarraysWithSum(nums []int, goal int) int {
	return sumAtMost(nums, goal) - sumAtMost(nums, goal-1)
}
