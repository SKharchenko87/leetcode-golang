package p2661

type Point struct {
	i, j int
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	indexes := make([]Point, m*n)
	cols := make([]int, n)
	rows := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			indexes[mat[i][j]-1] = Point{i, j}
		}
	}
	for res, a := range arr {
		rows[indexes[a-1].i]++
		if rows[indexes[a-1].i] == n {
			return res
		}
		cols[indexes[a-1].j]++
		if cols[indexes[a-1].j] == m {
			return res
		}
	}
	return m * n
}

func firstCompleteIndex0(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	indexesMap := make(map[int]Point, m*n)
	cols := make([]int, n)
	rows := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			indexesMap[mat[i][j]] = Point{i, j}
		}
	}
	for res, a := range arr {
		rows[indexesMap[a].i]++
		if rows[indexesMap[a].i] == n {
			return res
		}
		cols[indexesMap[a].j]++
		if cols[indexesMap[a].j] == m {
			return res
		}
	}
	return m * n
}
