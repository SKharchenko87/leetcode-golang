package p0078

func subsets(nums []int) [][]int {
	l := len(nums)
	lRes := 1 << l
	res := make([][]int, lRes)
	for i := 0; i < lRes; i++ {
		var tmp []int
		for j := 0; j < l && i>>j > 0; j++ {
			if checkBit(i, j) {
				tmp = append(tmp, nums[j])
			}
		}
		res[i] = tmp
	}
	return res
}

func checkBit(num int, position int) bool {
	return num&(1<<position) > 0
}
