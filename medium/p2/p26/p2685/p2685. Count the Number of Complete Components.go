package p2685

func countCompleteComponents(n int, edges [][]int) int {
	matrix := make([][]bool, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]bool, n)
		matrix[i][i] = true
	}
	for _, edge := range edges {
		matrix[edge[0]][edge[1]] = true
		matrix[edge[1]][edge[0]] = true
	}
	visited := make(map[int]bool, n)
	var res int
LOOP:
	for i := 0; i < n; i++ {
		if !visited[i] {
			currentConnectedComponent := make(map[int]bool, n)
			for j := 0; j < n; j++ {
				if matrix[i][j] {
					currentConnectedComponent[j] = true
				}
			}
			for vertices := range currentConnectedComponent {
				if !visited[vertices] {
					visited[vertices] = true
					for k := 0; k < n; k++ {
						if matrix[i][k] != matrix[vertices][k] {
							continue LOOP
						}
					}
				} else {
					continue LOOP
				}
			}
			res++
		}
	}
	return res
}
