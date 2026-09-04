package p3903

func firstStableIndex(nums []int, k int) int {
	n := len(nums)
	mins := make([]int, n)
	currMin := 1<<32 - 1
	for i := n - 1; i >= 0; i-- {
		currMin = min(currMin, nums[i])
		mins[i] = currMin
	}
	currMax := 0
	res := 1 << 32
	for i := 0; i < n; i++ {
		currMax = max(currMax, nums[i])
		if v := currMax - mins[i]; res > v && v <= k {
			return i
		}
	}
	return -1
}
