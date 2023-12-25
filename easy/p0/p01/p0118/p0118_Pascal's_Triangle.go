package main

import "fmt"

func generate(numRows int) [][]int {
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
