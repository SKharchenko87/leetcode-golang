package p2530

import (
	"container/heap"
	"math"
)

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] >= h[j] }
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))

}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	var maxHeap maxHeap
	heap.Init(&maxHeap)
	for i := 0; i < len(nums); i++ {
		heap.Push(&maxHeap, nums[i])
	}
	var res int64
	for i := 0; i < k; i++ {
		x := heap.Pop(&maxHeap).(int)
		res += int64(x)
		heap.Push(&maxHeap, int(math.Ceil(float64(x)/3)))
	}
	return res
}
