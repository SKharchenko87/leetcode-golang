package p3379

func constructTransformedArray(nums []int) []int {
	size := len(nums)
	result := make([]int, size)
	bigOffset := 100 * size
	for i := 0; i < size; i++ {
		result[i] = nums[(i+nums[i]+bigOffset)%size]
	}
	return result
}
