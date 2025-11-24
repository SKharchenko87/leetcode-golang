package p1018

func prefixesDivBy5(nums []int) []bool {
	l := len(nums)
	res := make([]bool, l)
	cur := 0
	for i := 0; i < l; i++ {
		cur = (cur*2 + nums[i]) % 5
		res[i] = cur == 0
	}
	return res
}
