package p1888

func minFlips(s string) int {
	if len(s) <= 1 {
		return 0
	}
	zeroEvenIndexCount := 0
	zeroOddIndexCount := 0
	oneEvenIndexCount := 0
	oneOddIndexCount := 0
	for i, ch := range s {
		if ch == '0' {
			if i%2 == 0 {
				zeroEvenIndexCount++
			} else {
				zeroOddIndexCount++
			}
		} else {
			if i%2 == 0 {
				oneEvenIndexCount++
			} else {
				oneOddIndexCount++
			}
		}
	}
	res := min(zeroEvenIndexCount+oneOddIndexCount, zeroOddIndexCount+oneEvenIndexCount)
	if len(s)%2 == 0 {
		return res
	} else {
		if res == 0 {
			return res
		}
		for i := 0; i < len(s); i++ {
			if s[i] == '0' {
				zeroEvenIndexCount--
			} else {
				oneEvenIndexCount--
			}

			zeroEvenIndexCount, zeroOddIndexCount = zeroOddIndexCount, zeroEvenIndexCount
			oneEvenIndexCount, oneOddIndexCount = oneOddIndexCount, oneEvenIndexCount

			if s[i] == '0' {
				zeroEvenIndexCount++
			} else {
				oneEvenIndexCount++
			}

			res = min(res, zeroEvenIndexCount+oneOddIndexCount, zeroOddIndexCount+oneEvenIndexCount)
		}
		return res
	}
}
