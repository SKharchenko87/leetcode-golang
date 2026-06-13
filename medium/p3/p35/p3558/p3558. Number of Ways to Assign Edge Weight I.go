package p3558

const MOD = 1e9 + 7

//[[2,3], [5,6], [4,5], [3,1], [6,1]]

func assignEdgeWeights(edges [][]int) int {
	l := len(edges)
	depth := make([]int, l+2)
	tree := make(map[int][]int, l)
	seq := make([]int, 0, l)
	newSeq := make([]int, 0, l)
	maxDepth := 0
	for i := 0; i < l; i++ {
		if _, ok := tree[edges[i][0]]; !ok {
			tree[edges[i][0]] = []int{edges[i][1]}
		} else {
			tree[edges[i][0]] = append(tree[edges[i][0]], edges[i][1])
		}
		if _, ok := tree[edges[i][1]]; !ok {
			tree[edges[i][1]] = []int{edges[i][0]}
		} else {
			tree[edges[i][1]] = append(tree[edges[i][1]], edges[i][0])
		}
	}
	seq = append(seq, 1)
	depth[1] = 1
	for len(seq) > 0 {
		newSeq = newSeq[:0]
		for _, v := range seq {
			for _, w := range tree[v] {
				if depth[w] == 0 {
					newSeq = append(newSeq, w)
					depth[w] = depth[v] + 1
					maxDepth = max(maxDepth, depth[w])
				}
			}
		}
		seq, newSeq = newSeq, seq
	}

	res := 1
	for i := 0; i < maxDepth-2; i++ {
		res = res * 2 % MOD
	}
	return res
}
