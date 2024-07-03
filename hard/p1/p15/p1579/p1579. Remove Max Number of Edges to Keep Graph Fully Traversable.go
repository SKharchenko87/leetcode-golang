package p1579

import "math/rand"

func maxNumEdgesToRemove(n int, edges [][]int) int {
	dsuAlice, dsuBob := DSU{p: make([]int, n+1), numOfSets: n}, DSU{p: make([]int, n+1), numOfSets: n}
	dsuAlice.makeSet()
	dsuBob.makeSet()
	numOfBothEdge := 0
	for _, edge := range edges {
		if edge[0] == 3 {
			if dsuAlice.union(edge[1], edge[2])+dsuBob.union(edge[1], edge[2]) > 0 {
				numOfBothEdge++
			}
		}
	}
	for _, edge := range edges {
		if edge[0] == 1 {
			numOfBothEdge += dsuAlice.union(edge[1], edge[2])
		} else if edge[0] == 2 {
			numOfBothEdge += dsuBob.union(edge[1], edge[2])
		}
	}
	if dsuAlice.numOfSets == 1 && dsuBob.numOfSets == 1 {
		return len(edges) - numOfBothEdge
	}
	return -1
}

type DSU struct {
	p         []int
	numOfSets int
}

func (dsu *DSU) makeSet() {
	for i := 0; i < len(dsu.p); i++ {
		dsu.p[i] = i
	}
}

func (dsu *DSU) find(x int) int {
	for x != dsu.p[x] {
		x, dsu.p[x] = dsu.p[x], dsu.p[dsu.p[x]]
	}
	return x
}

func (dsu *DSU) union(u, v int) int {
	uf := dsu.find(u)
	vf := dsu.find(v)
	if uf == vf {
		return 0
	}
	if rand.Int()%2 == 0 {
		uf, vf = vf, uf
	}
	dsu.p[uf] = vf
	dsu.numOfSets -= 1
	return 1
}
