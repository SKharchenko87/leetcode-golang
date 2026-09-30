package p1111

func maxDepthAfterSplit(seq string) []int {
	n := len(seq)
	cur := 0
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if seq[i] == '(' {
			cur++
			res[i] = 1 - cur%2
		} else {
			res[i] = 1 - cur%2
			cur--
		}
	}
	return res
}
