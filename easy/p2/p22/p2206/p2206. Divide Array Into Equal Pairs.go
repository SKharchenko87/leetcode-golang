package p2206

func divideArray(nums []int) bool {
	c := [501]bool{}
	x := 0
	for i := 0; i < len(nums); i++ {
		c[nums[i]] = !c[nums[i]]
		if c[nums[i]] {
			x++
		} else {
			x--
		}
	}
	return x == 0
}
