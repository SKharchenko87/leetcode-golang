package p0994

import "fmt"

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var check = func(ro map[int]map[int]bool, i, j int) {
		if _, ok := ro[i]; ok {
			ro[i][j] = true
		} else {
			ro[i] = map[int]bool{j: true}
		}
	}
	var fillNext = func(ro map[int]map[int]bool, i, j int) {
		//up
		if i != 0 && grid[i-1][j] == 1 {
			check(ro, i-1, j)
		}
		//left
		if j != 0 && grid[i][j-1] == 1 {
			check(ro, i, j-1)
		}
		//down
		if i != m-1 && grid[i+1][j] == 1 {
			check(ro, i+1, j)
		}
		//right
		if j != n-1 && grid[i][j+1] == 1 {
			check(ro, i, j+1)
		}
	}
	countFreshOrange := 0
	newRottenOrange := map[int]map[int]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				countFreshOrange++
			} else if grid[i][j] == 2 {
				fillNext(newRottenOrange, i, j)
			}
		}
	}
	count := 0
	for len(newRottenOrange) > 0 {
		fmt.Println(newRottenOrange)
		count++
		tmpNewRottenOrange := map[int]map[int]bool{}
		for i, arr := range newRottenOrange {
			for j, _ := range arr {
				grid[i][j] = 2
				if tmpNewRottenOrange[i][j] {
					delete(tmpNewRottenOrange[i], j)
					if len(tmpNewRottenOrange[i]) == 0 {
						delete(tmpNewRottenOrange, i)
					}
				}
				countFreshOrange--
				fillNext(tmpNewRottenOrange, i, j)
			}
		}
		newRottenOrange = tmpNewRottenOrange
	}
	if countFreshOrange > 0 {
		return -1
	}
	return count
}
