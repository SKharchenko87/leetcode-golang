package p3191

func minOperations(nums []int) int {
	l := len(nums)
	result := 0
	for i := 0; i < l-2; i++ {
		if nums[i] == 0 {
			nums[i] = 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			result++
		}
	}
	if nums[l-1] == 0 || nums[l-2] == 0 {
		return -1
	}
	return result
}

func minOperations0(nums []int) int {
	l := len(nums)
	result := 0
	for i := 0; i < l-2; i++ {
		if nums[i] == 0 {
			nums[i] = 1
			nums[i+1] = 1 - nums[i+1]
			nums[i+2] = 1 - nums[i+2]
			result++
		}
	}
	if nums[l-1] == 0 || nums[l-2] == 0 {
		return -1
	}
	return result
}
