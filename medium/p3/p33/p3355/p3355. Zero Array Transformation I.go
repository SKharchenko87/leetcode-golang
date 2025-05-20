package p3355

func isZeroArray(nums []int, queries [][]int) bool {
	l := len(nums)
	diff := make([]int, l+1)
	for _, query := range queries {
		diff[query[0]]--
		diff[query[1]+1]++
	}
	x := 0
	for i := 0; i < l; i++ {
		x += diff[i]
		if nums[i]+x > 0 {
			return false
		}

	}
	return true
}

func isZeroArray0(nums []int, queries [][]int) bool {
	l := len(nums)
	diff := make([]int, l)
	for _, query := range queries {
		diff[query[0]]--
		if query[1]+1 < l {
			diff[query[1]+1]++
		}
	}
	x := 0
	for i := 0; i < l; i++ {
		x += diff[i]
		if nums[i]+x > 0 {
			return false
		}

	}
	return true
}
