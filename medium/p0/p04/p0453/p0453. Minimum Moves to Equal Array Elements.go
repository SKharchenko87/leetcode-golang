package p0453

func minMoves(nums []int) int {
	sum := nums[0]
	minNums := nums[0]
	for _, num := range nums[1:] {
		sum += num
		minNums = min(minNums, num)
	}
	return sum - len(nums)*minNums
}
