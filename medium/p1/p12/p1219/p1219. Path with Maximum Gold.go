package p1219

import "maps"

// ToDo
func getMaximumGold(grid [][]int) int {
	rows, cols, gold := len(grid), len(grid[0]), 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			gold = max(gold, dfs(i, j, rows, cols, &grid))
		}
	}
	return gold
}

func dfs(row, col, rows, cols int, grid *[][]int) (maxGold int) {
	if row < 0 || row >= rows || col < 0 || col >= cols || (*grid)[row][col] == 0 {
		return 0
	}

	gold, directions := (*grid)[row][col], []int{-1, 0, 1, 0, -1}
	(*grid)[row][col] = 0
	for k := 0; k < 4; k++ {
		maxGold = max(maxGold, gold+dfs(row+directions[k], col+directions[k+1], rows, cols, &(*grid)))
	}

	(*grid)[row][col] = gold
	return maxGold
}

type Point struct {
	i, j int
}

func getMaximumGold0(grid [][]int) int {

	m, n := len(grid), len(grid[0])
	var dfs func(mapPoint map[Point]int, curPoint Point) int
	dfs = func(mapPoint map[Point]int, curPoint Point) int {
		if curPoint.i < 0 || curPoint.j < 0 || curPoint.i >= m || curPoint.j >= n {
			return 0
		}

		if grid[curPoint.i][curPoint.j] == 0 {
			return 0
		}

		tmp := Point{curPoint.i, curPoint.j}
		if _, ok := mapPoint[tmp]; ok {
			return 0
		}
		newMapPoint := maps.Clone(mapPoint)
		newMapPoint[tmp] = 1

		up := dfs(newMapPoint, Point{curPoint.i - 1, curPoint.j})
		right := dfs(newMapPoint, Point{curPoint.i, curPoint.j + 1})
		down := dfs(newMapPoint, Point{curPoint.i + 1, curPoint.j})
		left := dfs(newMapPoint, Point{curPoint.i, curPoint.j - 1})
		return grid[curPoint.i][curPoint.j] + max(up, right, down, left)
	}
	res := 0
	mapPoint := map[Point]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp := 0
			if grid[i][j] > 0 {
				tmp = dfs(mapPoint, Point{i, j})
			}
			if res < tmp {
				res = tmp
			}
		}
	}
	return res

}
