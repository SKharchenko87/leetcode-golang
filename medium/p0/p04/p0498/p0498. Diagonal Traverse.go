package p0498

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	result := make([]int, 0, m*n)
	i, j, index, dir := 0, 0, 0, 1
	for index < m*n {
		if -1 == i || i == m || -1 == j || j == n {
			if dir == 1 {
				i++
			} else {
				j++
			}
			dir *= -1
		}

		if 0 <= i && i < m && 0 <= j && j < n {
			result = append(result, mat[i][j])
			index++
		}
		i, j = i-dir, j+dir
	}
	return result
}

var directions = [][2]int{{-1, 1}, {1, -1}}

func findDiagonalOrder0(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	dirIndex := 0
	result := make([]int, 0, m*n)
	i, j, index := 0, 0, 0
	for index < m*n {
		if !(0 <= i && i < m && 0 <= j && j < n) {
			j += dirIndex
			dirIndex ^= 1
			i += dirIndex
		}

		if 0 <= i && i < m && 0 <= j && j < n {
			result = append(result, mat[i][j])
			index++
		}
		i, j = i+directions[dirIndex][0], j+directions[dirIndex][1]
	}
	return result
}

/*
0,0 (-1,1) => -1, 1 => 0,1 (1,-1)
0,1 (1,-1)
1,0 (1,-1) =>  2,-1 => 2,0 (-1,1)
2,0 (-1,1)
1,1 (-1,1)
0,2 (-1,1) => -1, 3 => 0,3 => 1,2 (1,-1)
1,2 (1,-1)
2,1 (1,-1) =>  3, 0 => 2,0 => 2,1 (-1,1)
2,2
*/
