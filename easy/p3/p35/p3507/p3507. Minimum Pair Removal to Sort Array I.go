package p3507

import (
	"math"
	"slices"
)

func minimumPairRemoval(nums []int) int {
	steps := 0
	for !slices.IsSorted(nums) {
		minSum := math.MaxInt
		minSumIndex := -1
		for i := 1; i < len(nums); i++ {
			if minSum > nums[i]+nums[i-1] {
				minSum = nums[i] + nums[i-1]
				minSumIndex = i
			}
		}
		nums[minSumIndex-1] += nums[minSumIndex]
		nums = slices.Delete(nums, minSumIndex, minSumIndex+1)
		steps++
	}
	return steps
}
