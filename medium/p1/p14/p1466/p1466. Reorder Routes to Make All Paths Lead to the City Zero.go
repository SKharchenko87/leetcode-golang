package p1466

func minReorder(n int, connections [][]int) int {
	path := make(map[int][]int, n)
	var fillPath func(i, j int)
	fillPath = func(i, j int) {
		if _, ok := path[i]; ok {
			path[i] = append(path[i], j)
		} else {
			path[i] = []int{j}
		}
	}
	for i := 0; i < n-1; i++ {
		arr := connections[i]
		fillPath(arr[0], i)
		fillPath(arr[1], i)
	}
	visited := make(map[int]bool, n)
	count := 0
	var dfs func(targetCity int)
	dfs = func(targetCity int) {
		visited[targetCity] = true
		for _, k := range path[targetCity] {
			if connections[k][0] == targetCity {
				if !visited[connections[k][1]] {
					count++
					dfs(connections[k][1])
				}
			} else {
				if !visited[connections[k][0]] {
					dfs(connections[k][0])
				}
			}
		}
	}
	dfs(0)
	return count
}
