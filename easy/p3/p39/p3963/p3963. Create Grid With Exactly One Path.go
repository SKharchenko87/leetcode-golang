package p3963

func createGrid(m int, n int) []string {
	firstRow := make([]byte, n)
	for i := range firstRow {
		firstRow[i] = '.'
	}
	anotherRow := make([]byte, n)
	for i := 0; i < n; i++ {
		anotherRow[i] = '#'
	}
	anotherRow[n-1] = '.'
	res := make([]string, m)
	res[0] = string(firstRow)
	anotherStr := string(anotherRow)
	for i := 1; i < m; i++ {
		res[i] = anotherStr
	}
	return res
}
