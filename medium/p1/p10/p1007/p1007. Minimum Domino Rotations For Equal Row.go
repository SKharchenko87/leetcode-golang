package p1007

func minDominoRotations(tops []int, bottoms []int) int {
	l := len(tops)

	cntA := 0
	cntB := 0
	cntC := 1
	cntD := 1
	for i := 1; i < l; i++ {
		calc(tops[0], l+1, tops[i], bottoms[i], &cntA)
		calc(bottoms[0], l+1, bottoms[i], tops[i], &cntB)
		calc(bottoms[0], l+1, tops[i], bottoms[i], &cntC)
		calc(tops[0], l+1, bottoms[i], tops[i], &cntD)
	}
	res := min(l-cntA%(l+1), l-cntB%(l+1), l-cntC%(l+1), l-cntD%(l+1), cntA, cntB, cntC, cntD)
	if res == l {
		return -1
	}
	return res
}

func calc(val, maxVal int, a, b int, cnt *int) {
	if *cnt != maxVal {
		if val == a {

		} else if val != a && val == b {
			*cnt++
		} else {
			*cnt = maxVal
		}
	}
}
