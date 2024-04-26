package p1289

func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	for i := n - 2; i >= 0; i-- {
		min0, min1 := get2Min(grid[i+1])
		for j := 0; j < n; j++ {
			if grid[i+1][j] == min0 {
				grid[i][j] += min1
			} else {
				grid[i][j] += min0
			}
		}
	}
	res := grid[0][0]
	for j := 0; j < n; j++ {
		if res > grid[0][j] {
			res = grid[0][j]
		}
	}
	return res
}

func get2Min(arr []int) (int, int) {
	min0, min1 := min(arr[0], arr[1]), max(arr[0], arr[1])
	for i := 2; i < len(arr); i++ {
		if min1 > arr[i] {
			if min0 > arr[i] {
				min0, min1 = arr[i], min0
			} else {
				min1 = arr[i]
			}
		}
	}
	return min0, min1
}
