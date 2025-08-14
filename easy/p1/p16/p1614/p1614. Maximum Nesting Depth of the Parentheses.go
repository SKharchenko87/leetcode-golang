package p1614

func maxDepth(s string) int {
	res, cur := 0, 0
	for _, c := range s {
		if c == '(' {
			cur++
		} else if c == ')' {
			cur--
		}
		if res < cur {
			res = cur
		}
	}
	return res
}
