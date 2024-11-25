package p1072

func maxEqualRowsAfterFlips(matrix [][]int) int {
	res := 0
	m := len(matrix)
	c := map[string]int{}
	for i := 0; i < m; i++ {
		str, reversStr := convertBinarySliceToString(matrix[i])
		c[str]++
		if res < c[str] {
			res = c[str]
		}
		c[reversStr]++
		if res < c[reversStr] {
			res = c[reversStr]
		}
	}
	return res
}

func convertBinarySliceToString(slice []int) (string, string) {
	str, reversStr := make([]byte, len(slice)), make([]byte, len(slice))
	for i, c := range slice {
		str[i] += byte(c)
		reversStr[i] += byte(c ^ 1)
	}
	return string(str), string(reversStr)
}

func convertBinarySliceToString0(slice []int) (string, string) {
	str, reversStr := make([]byte, len(slice)), make([]byte, len(slice))
	for i, c := range slice {
		str[i] += byte(c) + '0'
		reversStr[i] += byte(c^1) + '0'
	}
	return string(str), string(reversStr)
}

func maxEqualRowsAfterFlips1(matrix [][]int) int {
	res := 1
	m := len(matrix)
	n := len(matrix[0])
	for i0 := 0; i0 < m-1; i0++ {
		cur := 0
	LOOP:
		for i1 := i0; i1 < m; i1++ {
			eq, neq := 0, 0
			for j := 0; j < n; j++ {
				if matrix[i0][j] == matrix[i1][j] {
					eq++
				} else {
					neq++
				}
				if (eq-1) != j && (neq-1) != j {
					continue LOOP
				}
			}
			cur++
		}
		if res < cur {
			res = cur
		}
	}
	return res
}

func maxEqualRowsAfterFlips0(matrix [][]int) int {
	visited := map[int]bool{}
	res := 1
	m := len(matrix)
	n := len(matrix[0])
	for i0 := 0; i0 < m-1; i0++ {
		visited[i0] = true
		cur := 1
	LOOP:
		for i1 := i0; i1 < m; i1++ {
			if !visited[i1] {
				eq, neq := 0, 0
				for j := 0; j < n; j++ {
					if matrix[i0][j] == matrix[i1][j] {
						eq++
					} else {
						neq++
					}
					if (eq-1) != j && (neq-1) != j {
						continue LOOP
					}
				}
				cur++
				visited[i1] = true
			}
		}
		if res < cur {
			res = cur
		}
	}
	return res
}
