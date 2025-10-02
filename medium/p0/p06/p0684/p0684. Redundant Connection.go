package p0684

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(u int) int {
		if parent[u] != u {
			parent[u] = find(parent[u])
		}
		return parent[u]
	}

	union := func(u, v int) bool {
		pu, pv := find(u), find(v)
		if pu == pv {
			return false
		}
		parent[pv] = pu
		return true
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if !union(u, v) {
			return edge
		}
	}

	return nil
}
