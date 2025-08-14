package p2016

func maximumDifference(nums []int) int {
	maxDiff := -1
	preMin := nums[0]
	for i := 1; i < len(nums); i++ {
		if preMin >= nums[i] {
			preMin = nums[i]
		} else {
			maxDiff = max(maxDiff, nums[i]-preMin)
		}
	}
	return maxDiff
}

func maximumDifference0(nums []int) int {
	maxDiff := -1
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				maxDiff = max(maxDiff, nums[j]-nums[i])
			}
		}
	}
	return maxDiff
}
