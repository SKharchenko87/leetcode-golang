package p2149

func rearrangeArray(nums []int) (res []int) {
	res = make([]int, len(nums))
	ip, in := 0, 1
	for _, v := range nums {
		if v < 0 {
			res[in] = v
			in += 2
		} else {
			res[ip] = v
			ip += 2
		}
	}
	return res
}
