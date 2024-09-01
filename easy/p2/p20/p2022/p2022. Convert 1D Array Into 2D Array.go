package p2022

func construct2DArray(original []int, m int, n int) [][]int {
	l := len(original)
	if m*n != l {
		return [][]int{}
	}
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = original[i*n : (i+1)*n]
	}
	return result
}

func construct2DArray1(original []int, m int, n int) [][]int {
	l := len(original)
	if m*n != l {
		return [][]int{}
	}
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	for i := 0; i < l; i++ {
		result[i/n][i%n] = original[i]
	}
	return result
}

func construct2DArray0(original []int, m int, n int) [][]int {
	l := len(original)
	if m*n != l {
		return [][]int{}
	}
	result := make([][]int, m)
	for i := 0; i < l; i++ {
		if i%n == 0 {
			result[i/n] = make([]int, n)
		}
		result[i/n][i%n] = original[i]
	}
	return result
}
