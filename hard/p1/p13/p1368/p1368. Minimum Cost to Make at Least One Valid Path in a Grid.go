package p1368

const (
	RIGHT = 1
	LEFT  = 2
	DOWN  = 3
	UP    = 4
)

func minCost(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	candidate := make(map[[2]int]struct{}, 1)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = 10001
		}
	}
	matrix[0][0] = 0
	candidate[[2]int{0, 0}] = struct{}{}
	newCandidate := make(map[[2]int]struct{})
	newCandidate0 := make(map[[2]int]struct{})
	var f = func(i, j int, direction int, di, dj int) {
		if grid[i][j] == direction {
			if matrix[i+di][j+dj] > matrix[i][j] {
				newCandidate0[[2]int{i + di, j + dj}] = struct{}{}
				matrix[i+di][j+dj] = matrix[i][j]
			}
		} else {
			if matrix[i+di][j+dj] > matrix[i][j]+1 {
				newCandidate[[2]int{i + di, j + dj}] = struct{}{}
				matrix[i+di][j+dj] = matrix[i][j] + 1
			}
		}
	}
	for len(candidate) > 0 {
		for c, _ := range candidate {
			i, j := c[0], c[1]
			if i-1 >= 0 {
				f(i, j, UP, -1, 0)
			}
			if j+1 < m {
				f(i, j, RIGHT, 0, 1)
			}
			if i+1 < n {
				f(i, j, DOWN, 1, 0)
			}
			if j-1 >= 0 {
				f(i, j, LEFT, 0, -1)
			}
		}
		if len(newCandidate0) > 0 {
			candidate = newCandidate0
			newCandidate0 = make(map[[2]int]struct{})
		} else {
			candidate = newCandidate
			newCandidate = make(map[[2]int]struct{})
		}
	}
	return matrix[n-1][m-1]
}

func minCost0(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	candidate := make(map[[2]int]int, 1)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = 10001
		}
	}
	matrix[0][0] = 0
	candidate[[2]int{0, 0}] = 0
	newCandidate := make(map[[2]int]int)
	newCandidate0 := make(map[[2]int]int)
	for len(candidate) > 0 {
		for c, v := range candidate {
			i, j := c[0], c[1]
			if i-1 >= 0 {
				if grid[i][j] == UP {
					if matrix[i-1][j] > v {
						newCandidate0[[2]int{i - 1, j}] = v
						matrix[i-1][j] = v
					}
				} else {
					if matrix[i-1][j] > v+1 {
						newCandidate[[2]int{i - 1, j}] = v + 1
						matrix[i-1][j] = v + 1
					}
				}
			}
			if j+1 < m {
				if grid[i][j] == RIGHT {
					if matrix[i][j+1] > v {
						newCandidate0[[2]int{i, j + 1}] = v
						matrix[i][j+1] = v
					}
				} else {
					if matrix[i][j+1] > v {
						newCandidate[[2]int{i, j + 1}] = v + 1
						matrix[i][j+1] = v + 1
					}
				}
			}
			if i+1 < n {
				if grid[i][j] == DOWN {
					if matrix[i+1][j] > v {
						newCandidate0[[2]int{i + 1, j}] = v
						matrix[i+1][j] = v
					}
				} else {
					if matrix[i+1][j] > v {
						newCandidate[[2]int{i + 1, j}] = v + 1
						matrix[i+1][j] = v + 1
					}
				}
			}
			if j-1 >= 0 {
				if grid[i][j] == LEFT {
					if matrix[i][j-1] > v {
						newCandidate0[[2]int{i, j - 1}] = v
						matrix[i][j-1] = v
					}
				} else {
					if matrix[i][j-1] > v {
						newCandidate[[2]int{i, j - 1}] = v + 1
						matrix[i][j-1] = v + 1
					}
				}
			}
		}
		if len(newCandidate0) > 0 {
			candidate = newCandidate0
			newCandidate0 = make(map[[2]int]int)
		} else {
			candidate = newCandidate
			newCandidate = make(map[[2]int]int)
		}
	}
	return matrix[n-1][m-1]
}
