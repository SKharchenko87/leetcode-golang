package p1905

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	res := 0
	for i, line := range grid2 {
		for j, v := range line {
			if v == 1 && recursion(i, j, grid1, grid2) {
				res++
			}
		}
	}
	return res
}

func recursion(i, j int, grid1, grid2 [][]int) bool {
	if i < 0 || i >= len(grid2) || j < 0 || j >= len(grid2[0]) || grid2[i][j] == 0 {
		return true
	}
	grid2[i][j] = 0
	up := recursion(i-1, j, grid1, grid2)
	right := recursion(i, j+1, grid1, grid2)
	down := recursion(i+1, j, grid1, grid2)
	left := recursion(i, j-1, grid1, grid2)

	return grid1[i][j] == 1 && up && right && down && left
}
