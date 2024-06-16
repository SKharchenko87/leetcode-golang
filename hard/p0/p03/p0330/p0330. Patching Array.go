package p0330

import "fmt"

func minPatches(nums []int, n int) int {
	var res, i int
	sum := 1
	for sum < n {
		if i < len(nums) && nums[i] <= sum {
			sum += nums[i]
			i++
		} else {
			sum += sum
			res++
		}
	}
	return res
}

func minPatches1(nums []int, n int) int {
	var res, i int
	if nums[0] != 1 {
		res++
	} else {
		i = 1
	}
	sum := 1
	prevEnd := 1
	currStart := 1
	for ; sum < n; i++ {
		if i < len(nums) && prevEnd+1 >= nums[i] {
			currStart = nums[i]
		} else {
			if i < len(nums) {
				i--
			}
			res++
			currStart = prevEnd + 1
		}
		prevEnd = prevEnd + currStart
		sum += currStart
	}
	return res
}

func minPatches0(nums []int, n int) int {
	var res int
	if nums[0] != 1 {
		nums = append([]int{1}, nums...)
		res++
	}
	sum := nums[0]
	i := 1
	prevStart := 1
	prevEnd := 1
	currStart, currEnd := 1, 1
	for ; i < len(nums) && sum < n; i++ {
		if prevEnd+1 < nums[i] {
			tmp := append([]int{prevEnd + 1}, nums[i:]...)
			nums = append(nums[:i], tmp...)
			res++
		}
		currStart = prevEnd + 1
		currEnd = prevEnd + min(currStart, nums[i])
		sum += nums[i]
		fmt.Printf("%d   %d-%d:%d-%d\n", nums[i], prevStart, prevEnd, currStart, currEnd)
		prevStart, prevEnd = currStart, currEnd
		prevEnd = currEnd
	}
	for ; sum < n; i++ {
		currStart = prevEnd + 1
		currEnd = prevEnd + currStart
		sum += currStart
		fmt.Printf("%d   %d-%d:%d-%d\n", currStart, prevStart, prevEnd, currStart, currEnd)
		prevStart, prevEnd = currStart, currEnd
		prevEnd = currEnd
		res++
	}
	return res
}
