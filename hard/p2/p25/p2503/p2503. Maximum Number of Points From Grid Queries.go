package p2503

import (
	"container/heap"
	"slices"
)

var DIRECTIONS = [4][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

type point struct {
	i, j int
}

type cell struct {
	point point
	val   int
}

type MinHeap []cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(cell))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxPoints(grid [][]int, queries []int) []int {

	n, m := len(grid), len(grid[0])
	fullAnswer := make([]int, n*m+1)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	mh := MinHeap{}
	heap.Init(&mh)
	startPoint := point{0, 0}
	startCell := cell{startPoint, grid[0][0]}
	mh.Push(startCell)

	visited[startPoint.i][startPoint.j] = true

	currentMax := -1
	cnt := 0
	for mh.Len() > 0 {
		cnt++
		currentCell := heap.Pop(&mh).(cell)
		currentMax = max(currentMax, currentCell.val)
		fullAnswer[cnt] = currentMax

		for _, direction := range DIRECTIONS {
			candidateI := currentCell.point.i + direction[0]
			candidateJ := currentCell.point.j + direction[1]
			if candidateI >= 0 && candidateI < n && candidateJ >= 0 && candidateJ < m && !visited[candidateI][candidateJ] {
				candidatePoint := point{candidateI, candidateJ}
				candidateVal := grid[candidateI][candidateJ]
				heap.Push(&mh, cell{candidatePoint, candidateVal})
				visited[candidateI][candidateJ] = true
			}
		}
	}

	l := len(queries)
	res := make([]int, l)
	for i := 0; i < l; i++ {
		index, _ := slices.BinarySearch(fullAnswer, queries[i])
		res[i] = index - 1
	}

	return res
}

type cell0 struct {
	i, j int
	val  int
}

type MinHeap0 []cell0

func (h MinHeap0) Len() int           { return len(h) }
func (h MinHeap0) Less(i, j int) bool { return h[i].val < h[j].val }
func (h MinHeap0) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap0) Push(x interface{}) {
	*h = append(*h, x.(cell0))
}

func (h *MinHeap0) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxPoints0(grid [][]int, queries []int) []int {
	n, m := len(grid), len(grid[0])
	mh := MinHeap0{}
	heap.Init(&mh)
	fullAnswer := make([]int, n*m+1)
	mh.Push(cell0{0, 0, grid[0][0]})
	currentMax := grid[0][0]
	fullAnswer[1] = currentMax

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	cnt := 1
	visited[0][0] = true
	for mh.Len() > 0 {
		current := heap.Pop(&mh).(cell0)
		currentMax = max(currentMax, current.val)

		for _, direction := range DIRECTIONS {
			if current.i+direction[0] >= 0 && current.i+direction[0] < n && current.j+direction[1] >= 0 && current.j+direction[1] < m && !visited[current.i+direction[0]][current.j+direction[1]] {
				heap.Push(&mh, cell0{current.i + direction[0], current.j + direction[1], grid[current.i+direction[0]][current.j+direction[1]]})
				visited[current.i+direction[0]][current.j+direction[1]] = true
			}
		}
		fullAnswer[cnt] = currentMax
		cnt++
	}
	l := len(queries)
	res := make([]int, l)
	for i := 0; i < l; i++ {

		index, _ := slices.BinarySearch(fullAnswer, queries[i])
		res[i] = index - 1
	}
	return res
}
