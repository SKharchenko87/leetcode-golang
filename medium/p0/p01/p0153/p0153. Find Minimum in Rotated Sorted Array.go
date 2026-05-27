package p0153

func findMin(nums []int) int {
	l := len(nums) - 1
	if l == 0 || nums[0] < nums[l] {
		return nums[0]
	}
	if nums[0] > nums[l/2] {
		return findMin(nums[1 : l/2+1])
	}
	return findMin(nums[l/2+1:])
}
