package p2078

func maxDistance(colors []int) int {
	n := len(colors)
	if colors[n-1] != colors[0] {
		return n - 1
	}
	index2 := 1
	for ; index2 < n && colors[index2] == colors[0]; index2++ {
	}
	res := max(index2, n-index2) - 1
	index := n - 2
	for ; index > index2 && index > res && colors[index] == colors[0]; index-- {
	}
	return max(res, index)
}

func maxDistance0(colors []int) int {
	n := len(colors)
	res := 0
	for i := 0; res+i < n; i++ {
		for j := n - 1; j > i; j-- {
			if colors[i] != colors[j] {
				res = max(res, j-i)
				break
			}
		}
	}
	return res
}
