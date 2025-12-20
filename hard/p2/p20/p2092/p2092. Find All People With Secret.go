package p2092

import "container/heap"

type minHeap [][2]int

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	var ans = make([]int, n)
	graph := make([][][2]int, n)
	for _, meeting := range meetings {
		x, y, time := meeting[0], meeting[1], meeting[2]
		if graph[x] == nil {
			graph[x] = [][2]int{}
		}
		if graph[y] == nil {
			graph[y] = [][2]int{}
		}
		graph[x] = append(graph[x], [2]int{y, time})
		graph[y] = append(graph[y], [2]int{x, time})
	}
	mh := minHeap{}
	heap.Init(&mh)
	heap.Push(&mh, [2]int{firstPerson, 1})
	ans[0] = 1
	for _, ints := range graph[0] {
		heap.Push(&mh, ints)
	}

	for mh.Len() > 0 {
		x := heap.Pop(&mh).([2]int)
		if ans[x[0]] != 0 {
			continue
		}
		ans[x[0]] = x[1]
		for _, ints := range graph[x[0]] {
			if ans[ints[0]] == 0 && ints[1] >= x[1] {
				heap.Push(&mh, ints)
			}
		}

	}

	var result []int
	for i, x := range ans {
		if x > 0 {
			result = append(result, i)
		}
	}
	return result
}
