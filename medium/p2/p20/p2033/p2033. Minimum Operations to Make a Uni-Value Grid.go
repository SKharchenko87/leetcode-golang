package p2033

import "sort"

func minOperations(grid [][]int, x int) int {
	n := len(grid)
	m := len(grid[0])
	arr := make([]int, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			arr[i*m+j] = grid[i][j]
		}
	}
	sort.Ints(arr)
	avg := arr[len(arr)/2]
	count := 0
	for i := 0; i < n*m; i++ {
		if avg%x != arr[i]%x {
			return -1
		}
		diff := avg - arr[i]
		if diff < 0 {
			diff = -diff
		}
		count += diff / x
	}
	return count
}
