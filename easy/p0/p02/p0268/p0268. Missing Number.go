package p0268

func missingNumber(nums []int) int {
	l := len(nums)
	sum := (l + 1) * l / 2
	for i := 0; i < l; i++ {
		sum -= nums[i]
	}
	return sum
}
