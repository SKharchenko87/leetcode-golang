package p3203

func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	md10, md11 := minimumDiameter(edges1)
	md20, md21 := minimumDiameter(edges2)
	return max(max(md10+md11, md20+md21), md10+md20+1)
}

func minimumDiameter(edges [][]int) (int, int) {
	n := len(edges)
	if n <= 1 {
		return n, 0
	}
	graph := make(map[int]map[int]struct{}, n)
	for _, edge := range edges {
		if _, exists := graph[edge[0]]; !exists {
			graph[edge[0]] = map[int]struct{}{}
		}
		graph[edge[0]][edge[1]] = struct{}{}

		if _, exists := graph[edge[1]]; !exists {
			graph[edge[1]] = map[int]struct{}{}
		}
		graph[edge[1]][edge[0]] = struct{}{}
	}

	m0 := make(map[int]int, n)
	m1 := make(map[int]int, n)
	starts := make([]int, 0, n)
	parents := make([]int, 0, n)
	for key, ints := range graph {
		if len(ints) == 1 {
			starts = append(starts, key)
			parents = append(parents, -1)
			m0[key] = 0
			m1[key] = 0
		}
	}

	for i := 0; i < len(starts); i++ {
		cur := starts[i]
		parent := parents[i]
		for node := range graph[cur] {
			if parent != node {
				if m0[cur]+1 > m0[node] {
					m0[node], m1[node] = m0[cur]+1, m0[node]
				} else if m0[cur]+1 > m1[node] {
					m1[node] = m0[cur] + 1
				}
				delete(graph[node], cur)
				if len(graph[node]) == 1 {
					starts = append(starts, node)
					parents = append(parents, cur)
				}
			}
		}
	}
	center := starts[len(starts)-1]
	return m0[center], m1[center]
}
