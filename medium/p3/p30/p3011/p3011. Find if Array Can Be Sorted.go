package p3011

import "math/bits"

func canSortArray(nums []int) bool {
	prev := nums[0]
	prevMax := -1
	currMax := nums[0]
	prevGroup := countOneBits(prev)
	currGroup := prevGroup
	for i := 1; i < len(nums); i++ {

		if currGroup == countOneBits(nums[i]) {
			if prevMax > nums[i] {
				return false
			}
			if currMax < nums[i] {
				currMax = nums[i]
			}
		} else {
			prevMax = currMax
			if prevMax > nums[i] {
				return false
			}
			currMax = nums[i]
			currGroup = countOneBits(nums[i])
		}
	}
	return true
}

func countOneBits(num int) (res int) {
	for num > 0 {
		res += num & 1
		num >>= 1
	}
	return res
}

func canSortArray0(nums []int) bool {
	prev := nums[0]
	prevMax := -1
	currMax := nums[0]
	prevGroup := bits.OnesCount16(uint16(prev))
	currGroup := prevGroup
	for i := 1; i < len(nums); i++ {
		if bits.OnesCount16(uint16(nums[i])) == currGroup {
			if prevMax > nums[i] {
				return false
			}
			if currMax < nums[i] {
				currMax = nums[i]
			}
		} else {
			prevMax = currMax
			if prevMax > nums[i] {
				return false
			}
			currMax = nums[i]
			currGroup = bits.OnesCount16(uint16(nums[i]))
		}
	}
	return true
}
