package p2270

func waysToSplitArray(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	n--
	res := 0
	for i := 0; i < n; i++ {
		if 2*nums[i] >= nums[n] {
			res++
		}
	}
	return res
}
