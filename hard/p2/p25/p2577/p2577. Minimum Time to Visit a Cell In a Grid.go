package p2577

import (
	"container/heap"
)

var directions = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

type Node struct {
	i, j, cost int
}

type MinHeap []Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minimumTime(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if !(m > 1 && grid[0][1] <= 1 || n > 1 && grid[1][0] <= 1) {
		return -1
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, Node{0, 0, 0})
	visited[0][0] = true
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Node)

		for _, direction := range directions {
			i := cur.i + direction[0]
			j := cur.j + direction[1]
			if i < 0 || i > m-1 || j < 0 || j > n-1 || visited[i][j] {
				continue
			}
			var node Node
			if cur.cost+1 >= grid[i][j] {
				node = Node{i, j, cur.cost + 1}
			} else {
				node = Node{i, j, grid[i][j]}
				if (grid[i][j]-cur.cost)%2 == 0 {
					node.cost++
				}
			}
			if i == m-1 && j == n-1 {
				return node.cost
			}
			heap.Push(pq, node)
			visited[i][j] = true
		}
	}
	return -1
}
