package p3110

func scoreOfString(s string) int {
	prevC := int32(s[0])
	res := 0
	for _, c := range s {
		//res += int(abs(prevC, c))
		res += abs(int(prevC - c))
		prevC = c
	}
	return res
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

//func abs(a, b int32) int32 {
//	if a >= b {
//		return a - b
//	}
//	return b - a
//}
