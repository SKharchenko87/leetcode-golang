package p0198

func rob(nums []int) int {
	n := len(nums)
	switch n {
	case 0:
		return -1
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	}
	nums[2] = nums[2] + nums[0]
	for i := 3; i < len(nums); i++ {
		nums[i] = nums[i] + max(nums[i-2], nums[i-3])
	}
	return max(nums[n-1], nums[n-2])
}
