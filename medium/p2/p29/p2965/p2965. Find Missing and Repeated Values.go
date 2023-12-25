package main

import "fmt"

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	m := make(map[int]int, n*n)
	counter := 0
	twice := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			counter = counter + grid[i][j] - (j*n + i + 1)
			m[grid[i][j]]++
			if m[grid[i][j]] == 2 {
				twice = grid[i][j]
			}
		}
	}
	return []int{twice, -1 * (counter - twice)}
}

func main() {
	//grid := [][]int{{1, 3}, {2, 2}}
	grid := [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}
	fmt.Println(findMissingAndRepeatedValues(grid))
}

//0 1 2
//3 4 5
//6 7 8
//
//1 2 3
//4 5 6
//7 8 9

//1 2
//3 4
