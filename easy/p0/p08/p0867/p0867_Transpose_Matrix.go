package main

import "fmt"

func transpose(matrix [][]int) [][]int {

	dy := len(matrix)
	dx := len(matrix[0])

	res := make([][]int, dx)
	for i := range res {
		res[i] = make([]int, dy)
		for j := 0; j < dy; j++ {
			res[i][j] = matrix[j][i]
		}
	}
	return res
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//fmt.Println(matrix[1][2])

	fmt.Println(transpose(matrix))
}
