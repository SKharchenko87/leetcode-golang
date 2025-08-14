package p0238

func productExceptSelf(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	left := 1
	right := 1
	for i := 0; i < l; i++ {
		res[i] = 1
	}
	for i := 0; i < l; i++ {
		res[i], left = res[i]*left, left*nums[i]
		j := l - 1 - i
		res[j], right = res[j]*right, right*nums[j]
	}
	return res
}
