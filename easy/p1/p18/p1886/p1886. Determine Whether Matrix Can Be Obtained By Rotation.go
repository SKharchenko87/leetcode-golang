package p1886

/*ToDo*/
func check(m, tg [][]int, typeVal int) bool {
	n := len(m)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val := 0
			if typeVal == 0 {
				val = m[i][j]
			} else if typeVal == 1 {
				val = m[n-1-j][i]
			} else if typeVal == 2 {
				val = m[n-1-i][n-1-j]
			} else {
				val = m[j][n-1-i]
			}

			if val != tg[i][j] {
				return false
			}
		}
	}
	return true
}

func findRotation(m [][]int, tg [][]int) bool {
	for k := 0; k < 4; k++ {
		if check(m, tg, k) {
			return true
		}
	}
	return false
}
