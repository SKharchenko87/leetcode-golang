package p1800

func maxAscendingSum(nums []int) int {
	curSum := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			curSum += nums[i]
		} else {
			if curSum > sum {
				sum = curSum
			}
			curSum = nums[i]
		}
	}
	if curSum > sum {
		sum = curSum
	}
	return sum
}
