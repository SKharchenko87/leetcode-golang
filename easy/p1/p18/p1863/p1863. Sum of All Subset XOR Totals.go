package p1863

func subsetXORSum(nums []int) (res int) {
	for _, num := range nums {
		res |= num
	}
	return res << (len(nums) - 1)
}
