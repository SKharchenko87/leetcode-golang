package p2872

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	if n == 1 {
		return 1
	}
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = []int{}
	}
	degree := make([]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		degree[u]++
		degree[v]++
	}

	queue := make([]int, 0, n)
	for u, cnt := range degree {
		if cnt == 1 {
			queue = append(queue, u)
		}
	}

	res := 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		d := values[u] % k
		if d == 0 {
			res++
		}

		v := -1
		for _, v = range adj[u] {
			if degree[v] > 0 {
				degree[v]--
				values[v] = (values[v] + d) % k
				if degree[v] == 1 {
					queue = append(queue, v)
				}
				break
			}
		}
		degree[u]--
	}
	return res
}
