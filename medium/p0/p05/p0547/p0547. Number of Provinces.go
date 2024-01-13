package p0547

func findCircleNum(isConnected [][]int) int {
	count := 0
	n := len(isConnected)
	visited := make(map[int]bool, n)

	var dfs func(i int)
	dfs = func(i int) {
		if !visited[i] {
			visited[i] = true
			for j := 0; j < n; j++ {
				if isConnected[i][j] == 1 {
					dfs(j)
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
			count++
		}
	}

	return count
}
