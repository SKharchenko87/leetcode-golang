package main

import (
	"fmt"
	"slices"
)

func main() {
	grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}
	fmt.Println(equalPairs(grid))

}
func equalPairs(grid [][]int) int {
	n := len(grid)
	vertical := make([][]int, n)
	counter := 0
	for i := 0; i < n; i++ {
		vertical[i] = make([]int, n)
		for j := 0; j < n; j++ {
			vertical[i][j] = grid[j][i]
		}
		for j := 0; j < n; j++ {
			if slices.Equal(vertical[i], grid[j]) {
				counter++
			}
		}
	}
	//fmt.Println(vertical)
	return counter
}
