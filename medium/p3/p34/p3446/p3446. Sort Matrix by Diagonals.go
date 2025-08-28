package p3446

import "container/heap"

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	bottomLeft := maxHeap{}
	heap.Init(&bottomLeft)
	for i := 0; i < n; i++ {
		for d := 0; d < n-i; d++ {
			heap.Push(&bottomLeft, grid[i+d][d])
		}
		for d := 0; d < n-i; d++ {
			x := heap.Pop(&bottomLeft).(int)
			grid[i+d][d] = x
		}
	}

	topRight := minHeap{}
	heap.Init(&topRight)
	for j := 1; j < n; j++ {
		for d := 0; d < n-j; d++ {
			heap.Push(&topRight, grid[d][j+d])
		}
		for d := 0; d < n-j; d++ {
			x := heap.Pop(&topRight).(int)
			grid[d][j+d] = x
		}
	}

	return grid
}
