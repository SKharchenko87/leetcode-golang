package p3607

import "sort"

type DSU struct {
	root  []int
	next  []int
	del   []bool
	count int
	union map[int]int
}

func NewDSU(n int) *DSU {
	root := make([]int, n)
	next := make([]int, n)
	del := make([]bool, n)

	for i := 0; i < n; i++ {
		root[i] = i
		next[i] = -1
	}
	return &DSU{root: root, next: next, del: del, count: n}
}

func (ds *DSU) Find(i int) int {
	if ds.root[i] == i {
		return i
	}
	ds.root[i] = ds.Find(ds.root[i])
	return ds.root[i]
}

func (ds *DSU) Union(i int, j int) bool {
	rootI, rootJ := ds.Find(i), ds.Find(j)
	if rootI == rootJ {
		return false
	}
	minRoot := min(rootI, rootJ)
	ds.root[rootJ] = minRoot
	ds.root[rootI] = minRoot
	ds.count--
	return true
}

func (ds *DSU) BuildNext() {
	prev := make(map[int]int, ds.count)
	union := make(map[int]int, ds.count)
	for i := 0; i < len(ds.root); i++ {
		ds.Find(i)
		if ds.root[i] == i {
			prev[i] = ds.root[i]
			union[i] = i
		} else {
			ds.next[prev[ds.root[i]]] = i
			prev[ds.root[i]] = i
		}
	}
	ds.union = union
}

func (ds *DSU) Delete(i int) bool {
	if ds.del[i] {
		return false
	}

	if ds.union[ds.root[i]] == i {
		next := ds.next[i]
		for next != -1 && ds.del[next] {
			next = ds.next[next]
		}
		ds.union[ds.root[i]] = next
	}
	ds.del[i] = true
	return true
}

const (
	CheckMaintenance = 1
	Offline          = 2
)

func (ds *DSU) processQuery(kind, index int) int {
	switch kind {
	case CheckMaintenance:
		var x int
		if !ds.del[index] {
			x = index
		} else {
			x = ds.union[ds.root[index]]
		}
		if x == -1 {
			return -1
		} else {
			return x + 1
		}
	case Offline:
		ds.Delete(index)
		return 0
	}
	return 0
}

func (ds *DSU) GetCount() int {
	return ds.count
}

func processQueries(c int, connections [][]int, queries [][]int) []int {
	dsu := NewDSU(c)
	for _, conn := range connections {
		u, v := conn[0]-1, conn[1]-1
		dsu.Union(u, v)
	}
	dsu.BuildNext()
	result := make([]int, 0, len(queries))
	for _, q := range queries {
		r := dsu.processQuery(q[0], q[1]-1)
		if r != 0 {
			result = append(result, r)
		}
	}
	return result
}

/*
=================================================================================
*/

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type X struct {
	del     bool
	unionId int
	node    *Node
}

func processQueries0(c int, connections [][]int, queries [][]int) []int {
	graph := make(map[int][]int)
	for _, conn := range connections {
		u, v := conn[0]-1, conn[1]-1

		if _, ok := graph[u]; !ok {
			graph[u] = make([]int, 0, 1)
		}
		graph[u] = append(graph[u], v)

		if _, ok := graph[v]; !ok {
			graph[v] = make([]int, 0, 1)
		}
		graph[v] = append(graph[v], u)
	}
	visited := make([]bool, c)
	unionsMap := make(map[int]int)
	unionId := -1
	unions := make([][]int, 0)
	for i := 0; i < c; i++ {
		if !visited[i] {
			unionId++
			unions = append(unions, []int{})
			unions[unionId] = append(unions[unionId], i)
			visited[i] = true
			candidates := graph[i]
			for j := 0; j < len(candidates); j++ {
				if !visited[candidates[j]] {
					visited[candidates[j]] = true
					unions[unionId] = append(unions[unionId], candidates[j])
					unionsMap[candidates[j]] = unionId
					candidates = append(candidates, graph[candidates[j]]...)
				}
			}
		}
	}
	unionsList := make([]*Node, len(unions))
	unionsX := make([]X, c)
	for unionIndex, union := range unions {
		sort.Ints(union)
		unions[unionIndex] = union
		root := &Node{value: union[0]}
		unionsX[union[0]] = X{false, unionIndex, root}
		prev := root
		unionsList[unionIndex] = root
		for i := 1; i < len(union); i++ {
			curr := &Node{value: union[i], prev: prev}
			unionsX[union[i]] = X{false, unionIndex, curr}
			prev.next = curr
			prev = curr
		}
	}

	res := []int{}
	for _, query := range queries {
		y := query[1] - 1
		if query[0] == 2 {
			if !unionsX[y].del {
				unionsX[y].del = true
				if unionsX[y].node.prev != nil {
					unionsX[y].node.prev.next = unionsX[y].node.next
				} else {
					unionsList[unionsX[y].unionId] = unionsX[y].node.next
				}
				if unionsX[y].node.next != nil {
					unionsX[y].node.next.prev = unionsX[y].node.prev
				}
				unionsX[y].node.next = nil
				unionsX[y].node.prev = nil
			}
		} else if query[0] == 1 {
			if unionsX[y].del {

				curr := unionsList[unionsX[y].unionId]
				if curr == nil {
					res = append(res, -1)
				} else {
					res = append(res, curr.value+1)
				}

			} else {
				res = append(res, y+1)
			}

		}
	}
	return res
}
