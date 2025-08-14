package p2981

func maximumLength(s string) int {
	l := len(s)
	m := make(map[byte]*[3]int, 26)
	prev := s[0]
	cnt := 1
	maxCnt := 1
	var paste = func() {
		if cnt > 2 && cnt > maxCnt {
			maxCnt = cnt
		}
		if _, ok := m[prev]; !ok {
			m[prev] = &[3]int{}
		}
		if m[prev][0] >= cnt {
			return
		}
		if m[prev][2] <= cnt {
			m[prev][0], m[prev][1], m[prev][2] = m[prev][1], m[prev][2], cnt
		} else if m[prev][1] <= cnt {
			m[prev][0], m[prev][1] = m[prev][1], cnt
		} else if m[prev][0] < cnt {
			m[prev][0] = cnt
		}
	}
	for i := 1; i < l; i++ {
		if s[i] == prev {
			cnt++
		} else {
			paste()
			cnt = 1
			prev = s[i]
		}
	}
	paste()
	maxCnt -= 2

	for _, ints := range m {
		if ints[0] == ints[1] && ints[0] == ints[2] && ints[0] > maxCnt && ints[0] != 0 {
			maxCnt = ints[0]
		} else if ints[1] < ints[2] && ints[1] != 0 {
			if ints[1] > maxCnt {
				maxCnt = ints[1]
			}
		} else if ints[1] == ints[2] {
			if ints[1]-1 > maxCnt && ints[1] > 2 {
				maxCnt = ints[1] - 1
			}
		}
	}

	return maxCnt
}

//func paste(arr *[3]int, x int) {
//	if arr[0] >= x {
//		return
//	}
//	if arr[2] <= x {
//		arr[0], arr[1], arr[2] = arr[1], arr[2], x
//	} else if arr[1] <= x {
//		arr[0], arr[1] = arr[1], x
//	} else if arr[0] < x {
//		arr[0] = x
//	}
//}
