package p2567

import "sort"

const maxVal = 1<<63 - 1

func minimizeSum(nums []int) int {
	minList, maxList := [3]int{maxVal, maxVal, maxVal}, [3]int{-1, -1, -1}
	for _, n := range nums {
		if n >= maxList[2] {
			maxList[0], maxList[1], maxList[2] = maxList[1], maxList[2], n
		} else if n >= maxList[1] {
			maxList[0], maxList[1] = maxList[1], n
		} else if n >= maxList[0] {
			maxList[0] = n
		}
		if n <= minList[0] {
			minList[0], minList[1], minList[2] = n, minList[0], minList[1]
		} else if n <= minList[1] {
			minList[1], minList[2] = n, minList[1]
		} else if n <= minList[2] {
			minList[2] = n
		}
	}
	return min(maxList[0]-minList[0], maxList[1]-minList[1], maxList[2]-minList[2])
}

func minimizeSum0(nums []int) int {
	l := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return min(nums[l-1]-nums[2], nums[l-2]-nums[1], nums[l-3]-nums[0])
}
