package p1504

import "fmt"

func numSubmat(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	res := 0

	for j := 0; j < n; j++ {
		cur := 0
		for i := 0; i < m; i++ {
			if mat[i][j] == 0 {
				cur = mat[i][j]
			}
			mat[i][j] += cur
			cur = mat[i][j]

			if mat[i][j] != 0 {
				minVal := mat[i][j]
				curVal := 0
				for k := j; k >= 0 && mat[i][k] != 0; k-- {
					minVal = min(minVal, mat[i][k])
					curVal += minVal
				}
				res += curVal
			}
		}
	}

	return res
}

func numSubmat0(mat [][]int) int {
	m, n := len(mat), len(mat[0])

	for j := 0; j < n; j++ {
		cur := 0
		for i := 0; i < m; i++ {
			if mat[i][j] == 0 {
				cur = mat[i][j]
			}
			mat[i][j] += cur
			cur = mat[i][j]
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] != 0 {
				minVal := mat[i][j]
				curVal := 0
				for k := j; k >= 0 && mat[i][k] != 0; k-- {
					minVal = min(minVal, mat[i][k])
					curVal += minVal
				}
				mat[i][j] = curVal
				res += mat[i][j]
			}
		}
	}

	fmt.Println(mat)

	return res
}
