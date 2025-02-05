package p0827

func largestIsland(grid [][]int) int {
	n := len(grid)
	islandArea := make([]int, 0, n*n/2+2)
	islandArea = append(islandArea, 0, 0)
	var dfs func(i, j int, index int) int
	dfs = func(i, j int, index int) int {
		if i < 0 || i >= n || j < 0 || j >= n {
			return 0
		}
		if grid[i][j] != 1 {
			return 0
		}
		grid[i][j] = index
		return 1 + dfs(i-1, j, index) + dfs(i+1, j, index) + dfs(i, j-1, index) + dfs(i, j+1, index)

	}
	curIndex := 2
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				x := dfs(i, j, curIndex)
				if res < x {
					res = x
				}
				islandArea = append(islandArea, x)
				curIndex++
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				x := 1
				m := map[int]bool{}
				if i+1 < n {
					m[grid[i+1][j]] = true
				}
				if 0 <= i-1 {
					m[grid[i-1][j]] = true
				}
				if j+1 < n {
					m[grid[i][j+1]] = true
				}
				if 0 <= j-1 {
					m[grid[i][j-1]] = true
				}
				for k := range m {
					x += islandArea[k]
				}
				if res < x {
					res = x
				}
			}
		}
	}
	return res
}
