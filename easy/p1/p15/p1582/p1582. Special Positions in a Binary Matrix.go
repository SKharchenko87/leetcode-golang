package main

import "fmt"

func numSpecial(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	sumI := make([]int, m)
	sumJ := make([]int, n)
	counter := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sumI[i] += mat[i][j]
			sumJ[j] += mat[i][j]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if sumI[i] == 1 && sumJ[j] == 1 && mat[i][j] == 1 {
				counter++
			}
		}
	}
	return counter
}

func main() {
	//mat := [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}
	//mat := [][]int{{0, 0, 0, 1}, {1, 0, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	//mat := [][]int{{0, 0}, {0, 0}, {1, 0}}
	mat := [][]int{{0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 1, 0, 0, 1}, {0, 0, 0, 0, 1, 0, 0, 0}, {1, 0, 0, 0, 1, 0, 0, 0}, {0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(numSpecial(mat))
}
