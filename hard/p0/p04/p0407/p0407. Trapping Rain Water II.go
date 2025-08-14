package p0407

import (
	"container/heap"
)

type Node struct {
	i, j       int
	waterLevel int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].waterLevel < pq[j].waterLevel
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

var direction = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func trapRainWater(heightMap [][]int) int {
	n, m := len(heightMap), len(heightMap[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	candidates := &PriorityQueue{}
	heap.Init(candidates)

	j := 0
	for ; j < m; j++ {
		heap.Push(candidates, &Node{0, j, heightMap[0][j]})
		visited[0][j] = true
		heap.Push(candidates, &Node{n - 1, j, heightMap[n-1][j]})
		visited[n-1][j] = true
	}
	i := 1
	for ; i < n-1; i++ {
		heap.Push(candidates, &Node{i, 0, heightMap[i][0]})
		visited[i][0] = true
		heap.Push(candidates, &Node{i, m - 1, heightMap[i][m-1]})
		visited[i][m-1] = true
	}
	res := 0
	for candidates.Len() > 0 {
		candidate := heap.Pop(candidates).(*Node)
		i, j = candidate.i, candidate.j
		visited[i][j] = true
		for _, d := range direction {
			newI, newJ := i+d[0], j+d[1]
			if 0 < newI && newI < n-1 && 0 < newJ && newJ < m-1 && !visited[newI][newJ] {
				if heightMap[newI][newJ] < candidate.waterLevel {
					res += candidate.waterLevel - heightMap[newI][newJ]
					heightMap[newI][newJ] = candidate.waterLevel
				}
				heap.Push(candidates, &Node{newI, newJ, heightMap[newI][newJ]})
			}
		}
	}
	return res
}
