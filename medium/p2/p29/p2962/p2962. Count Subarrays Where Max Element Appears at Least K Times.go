package p2962

import "slices"

func countSubarrays(nums []int, k int) int64 {
	l := len(nums)

	currentMax := -1
	currentMaxCount := 0
	currentMaxIndexes := []int{}
	var res int64
	for rightIndex := 0; rightIndex < l; rightIndex++ {
		if nums[rightIndex] == currentMax {
			currentMaxIndexes = append(currentMaxIndexes, rightIndex)
			currentMaxCount++
		} else if nums[rightIndex] > currentMax {
			currentMaxIndexes = []int{rightIndex}
			currentMax = nums[rightIndex]
			currentMaxCount = 1
			res = 0
		}
		if currentMaxCount >= k {
			x := len(currentMaxIndexes)
			res += int64(currentMaxIndexes[x-k] + 1)
		}
	}
	return res
}

func countSubarrays1(nums []int, k int) int64 {
	l := len(nums)
	rightIndex := 0
	currentMax := slices.Max(nums)
	currentMaxCount := 0
	currentMaxIndexes := []int{}
	var res int64
	for ; rightIndex < l; rightIndex++ {
		if nums[rightIndex] == currentMax {
			currentMaxIndexes = append(currentMaxIndexes, rightIndex)
			currentMaxCount++
		}
		if currentMaxCount >= k {
			x := len(currentMaxIndexes)
			res += int64(currentMaxIndexes[x-k] + 1)
		}
	}
	return res
}

func countSubarrays0(nums []int, k int) int64 {
	l := len(nums)
	rightIndex := 1
	currentMax := nums[0]
	currentMaxCount := 1
	currentMaxIndexes := []int{0}
	var res int64
	for ; rightIndex < l; rightIndex++ {
		if nums[rightIndex] == currentMax {
			currentMaxIndexes = append(currentMaxIndexes, rightIndex)
			currentMaxCount++
		} else if nums[rightIndex] > currentMax {
			currentMaxIndexes = []int{rightIndex}
			currentMax = nums[rightIndex]
			currentMaxCount = 1
		}
		if currentMaxCount >= k {
			x := len(currentMaxIndexes)
			res += int64(currentMaxIndexes[x-k] + 1)
		}
	}
	return res
}
