package p3736

func minMoves(nums []int) int {
	l := len(nums)
	sum := 0
	maxElement := 0
	for i := 0; i < l; i++ {
		maxElement = max(maxElement, nums[i])
		sum += nums[i]
	}
	return maxElement*l - sum
}
