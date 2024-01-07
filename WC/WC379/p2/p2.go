package p2

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	if a == e || b == f {
		if a == e {
			if a == c {
				if b > d && d > f {
					return 2
				}
				if b < d && d < f {
					return 2
				}
			}
		}
		if b == f {
			if b == d {
				if a > c && c > e {
					return 2
				}
				if a < c && c < e {
					return 2
				}
			}
		}
		return 1
	}
	var lu, ld, ru, rd bool
	for i := 0; i < 8; i++ {
		if c+i == a && d+i == b {
			rd = true
		}
		if c+i == a && d-i == b {
			ld = true
		}
		if c-i == a && d+i == b {
			ru = true
		}
		if c-i == a && d-i == b {
			lu = true
		}
		if c+i == e && d+i == f {
			if rd {
				return 2
			} else {
				return 1
			}
		}
		if c+i == e && d-i == f {
			if ld {
				return 2
			} else {
				return 1
			}
		}
		if c-i == e && d+i == f {
			if ru {
				return 2
			} else {
				return 1
			}
		}
		if c-i == e && d-i == f {
			if lu {
				return 2
			} else {
				return 1
			}
		}
	}
	return 2
}
