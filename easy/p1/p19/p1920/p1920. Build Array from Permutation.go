package p1920

func buildArray0(nums []int) []int {
	l := len(nums)
	ans := make([]int, l)
	for i := 0; i < l; i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}

func buildArray(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		nums[i] += nums[nums[i]%1000] % 1000 * 1000
	}
	for i := 0; i < l; i++ {
		nums[i] /= 1000
	}
	return nums
}
