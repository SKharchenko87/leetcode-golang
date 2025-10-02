package p2

// abcdefghijklmnopqrstuvwxyz
func shiftDistance(s string, t string, nextCost []int, previousCost []int) int64 {
	var res int64

	for i := 0; i < len(s); i++ {
		var cntNext, cntPrev int
		sb, tb := int(s[i]-'a'), int(t[i]-'a')
		if sb == tb {

		} else if sb <= tb {
			cntNext = tb - sb
			cntPrev = sb + 26 - tb

		} else if sb > tb {
			cntNext = 26 - sb + tb
			cntPrev = sb - tb
		}

		tmpNext := 0
		for j := 0; j < cntNext; j++ {
			tmpNext += nextCost[(sb+j)%26]
		}
		tmpPrev := 0
		for j := 0; j < cntPrev; j++ {
			tmpPrev += previousCost[(26+sb-j)%26]
		}

		res += int64(min(tmpNext, tmpPrev))
	}
	return res
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
