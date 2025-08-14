package main

import "fmt"

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}
	if numRows == 1 {
		return result
	}
	result[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i/2+1; j++ {
			val := result[i-1][j-1] + result[i-1][j]
			result[i][j] = val
			result[i][i-j] = val
		}
	}
	return result
}

const N = 30

var triangle = make([][]int, N)

func init() {

	triangle[0] = []int{1}
	triangle[1] = []int{1, 1}
	for i := 2; i < N; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j <= (i+1)/2; j++ {
			val := triangle[i-1][j-1] + triangle[i-1][j]
			triangle[i][j] = val
			triangle[i][i-j] = val
		}
	}
}

func generate1(numRows int) [][]int {
	return triangle[:numRows]
}

func generate0(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		if i == 0 {
			res[0][0] = 1
		} else if i == 1 {
			res[1][0] = 1
			res[1][1] = 1
		} else {
			for j := 0; j <= i; j++ {
				if j == 0 || j == i {
					res[i][j] = 1
				} else {
					res[i][j] = res[i-1][j-1] + res[i-1][j]
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(generate(5))
}
