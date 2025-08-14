package p3243

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	dp := make([]int, n)
	res := []int{}
	roads := make([][]int, n)

	for i := 1; i < n; i++ {
		dp[i] = i
	}

	for i := 0; i < n-1; i++ {
		roads[i] = append(roads[i], i+1)
	}

	for _, q := range queries {
		source, dest := q[0], q[1]
		roads[source] = append(roads[source], dest)

		if dp[dest] > dp[source]+1 {
			dp[dest] = dp[source] + 1

			// Propagate the update to all nodes that could be affected
			queue := []int{dest}
			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]
				for _, neighbor := range roads[current] {
					if dp[neighbor] > dp[current]+1 {
						dp[neighbor] = dp[current] + 1
						queue = append(queue, neighbor)
					}
				}
			}
		}

		res = append(res, dp[n-1])
	}
	return res
}
