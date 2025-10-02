package p1514

import (
	"container/heap"
	"math"
)

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	maxProb := make([]float64, n)
	maxProb[start_node] = 1.0
	for i := 0; i < n; i++ {
		flgExit := true
		for j, edge := range edges {
			if t := maxProb[edge[0]] * succProb[j]; t > maxProb[edge[1]] {
				maxProb[edge[1]] = t
				flgExit = false
			}
			if t := maxProb[edge[1]] * succProb[j]; t > maxProb[edge[0]] {
				maxProb[edge[0]] = t
				flgExit = false
			}
		}
		if flgExit {
			break
		}
	}
	return math.Round(maxProb[end_node]*100_000) / 100_000
}

type way struct {
	index  int
	weight float64
}

type MaxHeap []way

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool { return h[i].weight > h[j].weight }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(way))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxProbability3(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	graph := make(map[int][]way, n)
	for i, edge := range edges {
		if _, ok := graph[edge[0]]; !ok {
			graph[edge[0]] = make([]way, 0)
		}
		graph[edge[0]] = append(graph[edge[0]], way{edge[1], succProb[i]})
		if _, ok := graph[edge[1]]; !ok {
			graph[edge[1]] = make([]way, 0)
		}
		graph[edge[1]] = append(graph[edge[1]], way{edge[0], succProb[i]})
	}
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	maxHeap.Push(way{start_node, 1.0})
	maxProb := make([]float64, n)
	maxProb[start_node] = 1.0
	for maxHeap.Len() > 0 {
		x := heap.Pop(maxHeap).(way)
		if x.index == end_node {
			return math.Round(x.weight*100_000) / 100_000
		}
		for _, node := range graph[x.index] {
			if tmp := x.weight * node.weight; tmp > maxProb[node.index] {
				maxProb[node.index] = tmp
				heap.Push(maxHeap, way{node.index, tmp})
			}
		}
	}
	return 0
}

func maxProbability2(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	graph := make(map[int][]way, n)
	for i, edge := range edges {
		if _, ok := graph[edge[0]]; !ok {
			graph[edge[0]] = make([]way, 0)
		}
		graph[edge[0]] = append(graph[edge[0]], way{edge[1], succProb[i]})
		if _, ok := graph[edge[1]]; !ok {
			graph[edge[1]] = make([]way, 0)
		}
		graph[edge[1]] = append(graph[edge[1]], way{edge[0], succProb[i]})
	}
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	maxHeap.Push(way{start_node, 1.0})
	visited := make(map[int]bool, n)
	for maxHeap.Len() > 0 {
		x := heap.Pop(maxHeap).(way)
		if x.index == end_node {
			return math.Round(x.weight*100_000) / 100_000
		}
		visited[x.index] = true
		for _, node := range graph[x.index] {
			if !visited[node.index] {
				heap.Push(maxHeap, way{node.index, x.weight * node.weight})
			}
		}
	}
	return 0
}

func maxProbability1(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	graph := make(map[int]map[int]float64, n)
	for i := 0; i < len(edges); i++ {
		if _, ok := graph[edges[i][0]]; !ok {
			graph[edges[i][0]] = make(map[int]float64)
		}
		graph[edges[i][0]][edges[i][1]] = succProb[i]

		if _, ok := graph[edges[i][1]]; !ok {
			graph[edges[i][1]] = make(map[int]float64)
		}
		graph[edges[i][1]][edges[i][0]] = succProb[i]
	}
	candidate := make(map[int]float64, n)
	candidate = graph[start_node]
	visited := make(map[int]float64, n)
	visited = graph[start_node]
	for len(candidate) != 0 {
		newCandidate := make(map[int]float64, n)
		for k, v := range candidate {
			for i, f := range graph[k] {
				if f*v > visited[i] {
					visited[i] = f * v
					if i != end_node {
						newCandidate[i] = f * v
					}
				}
			}
		}
		candidate = newCandidate
	}
	return math.Round(visited[end_node]*100_000) / 100_000
}

func maxProbability0(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	visited := make(map[int]float64, n)
	candidate := make(map[int]float64, n)
	for i := 0; i < len(edges); i++ {
		if edges[i][0] == start_node {
			visited[edges[i][1]] = succProb[i]
			if edges[i][1] != end_node {
				candidate[edges[i][1]] = succProb[i]
			}
		}
		if edges[i][1] == start_node {
			visited[edges[i][0]] = succProb[i]
			if edges[i][0] != end_node {
				candidate[edges[i][0]] = succProb[i]
			}
		}
	}
	for len(candidate) != 0 {
		newCandidate := make(map[int]float64, n)
		for k, v := range candidate {
			for i := 0; i < len(edges); i++ {
				if edges[i][0] == k && succProb[i]*v > visited[edges[i][1]] {
					visited[edges[i][1]] = succProb[i] * v
					if edges[i][1] != end_node {
						newCandidate[edges[i][1]] = succProb[i] * v
					}
				}
				if edges[i][1] == k && succProb[i]*v > visited[edges[i][0]] {
					visited[edges[i][0]] = succProb[i] * v
					if edges[i][0] != end_node {
						newCandidate[edges[i][0]] = succProb[i] * v
					}
				}
			}
		}
		candidate = newCandidate
	}
	return math.Round(visited[end_node]*100_000) / 100_000
}
