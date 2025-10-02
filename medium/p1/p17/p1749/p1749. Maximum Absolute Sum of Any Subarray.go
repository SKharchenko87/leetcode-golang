package p1749

func maxAbsoluteSum(nums []int) int {
	sum := 0
	minSum := 0
	maxSum := 0
	for _, num := range nums {
		sum += num
		minSum = min(minSum, sum)
		maxSum = max(maxSum, sum)
	}
	return maxSum - minSum
}
