package p3208

func numberOfAlternatingGroups(colors []int, k int) int {
	l := len(colors)
	res := 0
	d := 1
	for i := 1; i < l+k-1; i++ {
		if colors[i%l] == colors[(i-1)%l] {
			d = 1
		} else {
			d++
		}
		if d >= k {
			res++
		}
	}
	return res
}
