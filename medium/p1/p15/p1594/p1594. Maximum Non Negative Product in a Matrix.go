package p1594

const MOD = 1e9 + 7

/*
ToDo
можно использовать одну строку positive, negative
или вообще негативы хранить в grid
*/
func maxProductPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	positive, negative := make([][]*int, m), make([][]*int, m)
	zeroHas := false
	for i := 0; i < m; i++ {
		positive[i], negative[i] = make([]*int, n), make([]*int, n)
	}
	if grid[0][0] >= 0 {
		zeroHas = zeroHas || grid[0][0] == 0
		positive[0][0] = &grid[0][0]
	} else {
		negative[0][0] = &grid[0][0]
	}
	for j := 1; j < n; j++ {
		if grid[0][j] >= 0 {
			zeroHas = zeroHas || grid[0][j] == 0
			if positive[0][j-1] != nil {
				x := grid[0][j] * *positive[0][j-1]
				positive[0][j] = &x
			}
			if negative[0][j-1] != nil {
				x := grid[0][j] * *negative[0][j-1]
				if x == 0 {
					positive[0][j] = &x
				} else {
					if x == 0 {
						positive[0][j] = &x
					} else {
						negative[0][j] = &x
					}
				}
			}
		} else {
			if positive[0][j-1] != nil {
				x := grid[0][j] * *positive[0][j-1]
				if x == 0 {
					positive[0][j] = &x
				} else {
					if x == 0 {
						positive[0][j] = &x
					} else {
						negative[0][j] = &x
					}
				}
			}
			if negative[0][j-1] != nil {
				x := grid[0][j] * *negative[0][j-1]
				positive[0][j] = &x
			}
		}
	}
	for i := 1; i < m; i++ {
		if grid[i][0] >= 0 {
			zeroHas = zeroHas || grid[i][0] == 0
			if positive[i-1][0] != nil {
				x := grid[i][0] * *positive[i-1][0]
				positive[i][0] = &x
			}
			if negative[i-1][0] != nil {
				x := grid[i][0] * *negative[i-1][0]
				if x == 0 {
					positive[i][0] = &x
				} else {
					negative[i][0] = &x
				}
			}
		} else {
			if positive[i-1][0] != nil {
				x := grid[i][0] * *positive[i-1][0]
				if x == 0 {
					positive[i][0] = &x
				} else {
					negative[i][0] = &x
				}
			}
			if negative[i-1][0] != nil {
				x := grid[i][0] * *negative[i-1][0]
				positive[i][0] = &x
			}
		}

		for j := 1; j < n; j++ {
			if grid[i][j] >= 0 {
				zeroHas = zeroHas || grid[i][j] == 0
				if positive[i-1][j] != nil {
					x := grid[i][j] * *positive[i-1][j]
					positive[i][j] = &x
				}
				if negative[i-1][j] != nil {
					x := grid[i][j] * *negative[i-1][j]
					if x == 0 {
						positive[i][j] = &x
					} else {
						negative[i][j] = &x
					}
				}

				if positive[i][j-1] != nil {
					x := grid[i][j] * *positive[i][j-1]
					if positive[i][j] == nil {
						positive[i][j] = &x
					} else {
						if abs(*positive[i][j]) < abs(x) {
							positive[i][j] = &x
						}
					}
				}
				if negative[i][j-1] != nil {
					x := grid[i][j] * *negative[i][j-1]
					if negative[i][j] == nil {
						if x == 0 {
							positive[i][j] = &x
						} else {
							negative[i][j] = &x
						}
					} else {
						if abs(*negative[i][j]) < abs(x) {
							if x == 0 {
								positive[i][j] = &x
							} else {
								negative[i][j] = &x
							}
						}
					}
				}

			} else {
				if positive[i-1][j] != nil {
					x := grid[i][j] * *positive[i-1][j]
					if x == 0 {
						positive[i][j] = &x
					} else {
						negative[i][j] = &x
					}
				}
				if negative[i-1][j] != nil {
					x := grid[i][j] * *negative[i-1][j]
					positive[i][j] = &x
				}

				if negative[i][j-1] != nil {
					x := grid[i][j] * *negative[i][j-1]
					if positive[i][j] == nil {
						positive[i][j] = &x
					} else {
						if abs(*positive[i][j]) < abs(x) {
							positive[i][j] = &x
						}
					}
				}
				if positive[i][j-1] != nil {
					x := grid[i][j] * *positive[i][j-1]
					if x == 0 {
						if positive[i][j] == nil {
							positive[i][j] = &x
						} else {
							if abs(*positive[i][j]) < abs(x) {
								positive[i][j] = &x
							}
						}
					} else if negative[i][j] == nil {
						if x == 0 {
							positive[i][j] = &x
						} else {
							negative[i][j] = &x
						}
						negative[i][j] = &x
					} else {
						if abs(*negative[i][j]) < abs(x) {
							if x == 0 {
								positive[i][j] = &x
							} else {
								negative[i][j] = &x
							}
							negative[i][j] = &x
						}
					}
				}
			}
		}
	}
	if positive[m-1][n-1] == nil {
		if zeroHas {
			return 0
		}
		return -1
	}
	return *positive[m-1][n-1] % MOD
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}
