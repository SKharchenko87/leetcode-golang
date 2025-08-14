package p2874

func maximumTripletValue(nums []int) int64 {
	l := len(nums)
	leftMax := max(nums[0], nums[1])
	diffMax := int64(leftMax - nums[1])
	var res int64
	for i := 2; i < l; i++ {
		c := diffMax * int64(nums[i])
		if res < c {
			res = c
		}
		diffMax = max(diffMax, int64(leftMax-nums[i]))
		leftMax = max(leftMax, nums[i])
	}
	return res
}

func maximumTripletValue0(nums []int) int64 {
	l := len(nums)
	leftMax, rightMax := make([]int, l), make([]int, l)
	currentMax := nums[0]
	for i := 1; i < l; i++ {
		leftMax[i] = currentMax
		currentMax = max(currentMax, nums[i])
	}
	currentMax = nums[l-1]
	for i := l - 2; i >= 0; i-- {
		rightMax[i] = currentMax
		currentMax = max(currentMax, nums[i])
	}
	var res int64
	for i := 1; i < l-1; i++ {
		c := int64(leftMax[i]-nums[i]) * int64(rightMax[i])
		if res < c {
			res = c
		}
	}
	return res
}
