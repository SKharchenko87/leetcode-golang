package p3650

func minCost(n int, edges [][]int) int {

	graph := make([]map[int]int, n)
	for _, edge := range edges {
		if graph[edge[0]] == nil {
			graph[edge[0]] = map[int]int{}
		}
		if _, ok := graph[edge[0]][edge[1]]; graph[edge[0]][edge[1]] > edge[2] || !ok {
			graph[edge[0]][edge[1]] = edge[2]
		}

		if graph[edge[1]] == nil {
			graph[edge[1]] = map[int]int{}
		}
		if _, ok := graph[edge[1]][edge[0]]; graph[edge[1]][edge[0]] > edge[2]*2 || !ok {
			graph[edge[1]][edge[0]] = edge[2] * 2
		}
	}

	cost := make([]int, n)
	cost[n-1] = -1
	currentU := make([]int, 0, n/2)
	currentV := make([]int, 0, n/2)
	for u, w := range graph[0] {
		cost[u] = w
		currentU = append(currentU, u)
	}
	cost[0] = 1

	for {
		for _, u := range currentU {
			for v, w := range graph[u] {
				if w+cost[u] < cost[v] || cost[v] < 1 {
					cost[v] = w + cost[u]
					currentV = append(currentV, v)
				}
			}
		}
		currentU, currentV = currentV, currentU
		currentV = currentV[:0]
		if len(currentU) == 0 {
			break
		}
	}

	return cost[n-1]
}

func minCost0(n int, edges [][]int) int {
	graph := make([]map[int]int, n)
	for _, edge := range edges {
		if graph[edge[0]] == nil {
			graph[edge[0]] = map[int]int{}
		}
		if _, ok := graph[edge[0]][edge[1]]; graph[edge[0]][edge[1]] > edge[2] || !ok {
			graph[edge[0]][edge[1]] = edge[2]
		}

		if graph[edge[1]] == nil {
			graph[edge[1]] = map[int]int{}
		}
		if _, ok := graph[edge[1]][edge[0]]; graph[edge[1]][edge[0]] > edge[2]*2 || !ok {
			graph[edge[1]][edge[0]] = edge[2] * 2
		}
	}

	cost := make([]int, n)
	currentU := make([]int, 0, n/2)

	for k, v := range graph[0] {
		cost[k] = v
		currentU = append(currentU, k)
	}
	cost[0] = 1

	for i := 0; i < len(currentU); i++ {
		u := currentU[i]
		for k, v := range graph[u] {
			if v+cost[u] < cost[k] || cost[k] == 0 {
				cost[k] = v + cost[u]
				currentU = append(currentU, k)
			}
		}
	}
	if cost[n-1] == 0 {
		return -1
	}
	return cost[n-1]
}
