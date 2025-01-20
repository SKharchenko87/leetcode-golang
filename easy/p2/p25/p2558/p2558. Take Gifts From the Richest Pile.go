package p2558

import (
	"container/heap"
	"math"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func pickGifts(gifts []int, k int) int64 {
	queue := MaxHeap{}
	heap.Init(&queue)
	var sum int64
	for _, gift := range gifts {
		sum += int64(gift)
		heap.Push(&queue, gift)
	}
	for i := 0; i < k && queue.Len() > 0; i++ {
		x := heap.Pop(&queue).(int)
		sum -= int64(x)
		y := int(math.Sqrt(float64(x)))
		sum += int64(y)
		heap.Push(&queue, y)
	}
	return sum
}
