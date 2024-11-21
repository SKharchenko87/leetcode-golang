package p2257

var directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

const (
	UNGUARD                         = 0
	GUARDED_HORIZONTAL              = 1
	GUARDED_VERTICAL                = 2
	GUARDED_HORIZONTAL_AND_VERTICAL = 3
	GUARD                           = 4
	WALL                            = 5
)

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	tab := make([][]byte, m)
	for i := range tab {
		tab[i] = make([]byte, n)
	}
	res := m * n
	for _, wall := range walls {
		tab[wall[0]][wall[1]] = WALL
		res--
	}
	for _, guard := range guards {
		if tab[guard[0]][guard[1]] == UNGUARD {
			res--
		}
		tab[guard[0]][guard[1]] = GUARD
		for b, direction := range directions {
			x := byte(b%2 + 1)
			i, j := guard[0]+direction[0], guard[1]+direction[1]
			pV, pH := 0, 0
			for ; 0 <= i && i < m && 0 <= j && j < n && tab[i][j] < GUARD; i, j = i+direction[0], j+direction[1] {
				if tab[i][j] == GUARDED_HORIZONTAL_AND_VERTICAL {
					pV++
					pH++
				} else if tab[i][j] == GUARDED_VERTICAL {
					pV++
				} else if tab[i][j] == GUARDED_HORIZONTAL {
					pH++
				} else if tab[i][j] == UNGUARD {
					tab[i][j] = x
					res--
				} else if tab[i][j]%2 != x {
					tab[i][j] = GUARDED_HORIZONTAL_AND_VERTICAL
				}
				if pV >= 2 && b%2 == 1 {
					break
				}
				if pH >= 2 && b%2 == 0 {
					break
				}
			}
		}
	}
	//fmt.Println(tab)
	return res
}

func countUnguarded1(m int, n int, guards [][]int, walls [][]int) int {
	tab := make([][]int, m)
	for i := range tab {
		tab[i] = make([]int, n)
	}
	res := m * n
	for _, wall := range walls {
		tab[wall[0]][wall[1]] = 5
		res--
	}
	for _, guard := range guards {
		tab[guard[0]][guard[1]] = 4
		res--
	}
	for _, guard := range guards {
		for _, direction := range directions {
			i, j := guard[0]+direction[0], guard[1]+direction[1]
			for ; 0 <= i && i < m && 0 <= j && j < n; i, j = i+direction[0], j+direction[1] {
				if tab[i][j] == 0 {
					tab[i][j] = 1
					res--
				} else if tab[i][j] >= 4 {
					break
				}
			}
		}
	}
	return res
}

// TLE
func countUnguarded0(m int, n int, guards [][]int, walls [][]int) int {
	tab := make([][]int, m)
	for i := range tab {
		tab[i] = make([]int, n)
	}
	res := m * n
	for _, wall := range walls {
		tab[wall[0]][wall[1]] = 2
		res--
	}
	for _, guard := range guards {
		for _, direction := range directions {
			i, j := guard[0], guard[1]
			for ; 0 <= i && i < m && 0 <= j && j < n && tab[i][j] != 2; i, j = i+direction[0], j+direction[1] {
				if tab[i][j] == 0 {
					tab[i][j] = 1
					res--
				}
			}
		}
	}
	return res
}
