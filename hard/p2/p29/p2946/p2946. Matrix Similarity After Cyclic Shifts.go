package p2946

func areSimilar(mat [][]int, k int) bool {
	m, n := len(mat), len(mat[0])
	if k%n == 0 {
		return true
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != mat[i][(j+k)%n] {
				return false
			}
		}
	}
	return true
}
