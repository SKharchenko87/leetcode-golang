package p3498

func reverseDegree(s string) int {
	res := 0
	for i, ch := range s {
		res += (i + 1) * int('z'-ch+1)
	}
	return res
}
