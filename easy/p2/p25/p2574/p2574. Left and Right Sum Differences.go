package p2574

func leftRightDifference(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	leftSum := 0
	for i := 0; i < n; i++ {
		leftSum += nums[i]
	}

	rightSum := 0
	for i := n - 1; i >= 0; i-- {
		leftSum -= nums[i]
		res[i] = abs(leftSum - rightSum)
		rightSum += nums[i]
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
