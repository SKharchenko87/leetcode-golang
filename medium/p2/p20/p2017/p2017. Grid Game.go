package p2017

import "math"

func gridGame(grid [][]int) int64 {
	n := len(grid[0])
	upSum := 0
	for i := 0; i < n; i++ {
		upSum += grid[0][i]
	}

	downSum := 0
	res := math.MaxInt
	for i := 0; i < n; i++ {
		upSum -= grid[0][i]
		res = min(res, max(upSum, downSum))
		downSum += grid[1][i]
	}

	return int64(res)
}

func gridGame0(grid [][]int) int64 {
	n := len(grid[0])
	upSum := 0
	for i := 0; i < n; i++ {
		upSum += grid[0][i]
	}

	tmpUp := 0
	tmpDown := 0
	res := math.MaxInt
	for i := 0; i < n; i++ {
		tmpUp += grid[0][i]
		res = min(res, max(upSum-tmpUp, tmpDown))
		tmpDown += grid[1][i]
	}

	return int64(res)
}
