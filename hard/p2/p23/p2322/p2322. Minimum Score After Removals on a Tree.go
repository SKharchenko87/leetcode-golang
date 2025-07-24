package p2322

import (
	"math"
)

// ToDo
func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	sum := make([]int, n)
	in := make([]int, n)
	out := make([]int, n)
	cnt := 0

	var dfs func(int, int)
	dfs = func(x, fa int) {
		in[x] = cnt
		cnt++
		sum[x] = nums[x]
		for _, y := range adj[x] {
			if y == fa {
				continue
			}
			dfs(y, x)
			sum[x] ^= sum[y]
		}
		out[x] = cnt
	}

	dfs(0, -1)

	res := math.MaxInt32
	for u := 1; u < n; u++ {
		for v := u + 1; v < n; v++ {
			if in[v] > in[u] && in[v] < out[u] {
				res = min(res, calc(sum[0]^sum[u], sum[u]^sum[v], sum[v]))
			} else if in[u] > in[v] && in[u] < out[v] {
				res = min(res, calc(sum[0]^sum[v], sum[v]^sum[u], sum[u]))
			} else {
				res = min(res, calc(sum[0]^sum[u]^sum[v], sum[u], sum[v]))
			}
		}
	}
	return res
}

func calc(part1, part2, part3 int) int {
	return max(part1, max(part2, part3)) - min(part1, min(part2, part3))
}
