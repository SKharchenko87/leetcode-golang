package p3066

import (
	"container/heap"
	"sort"
)

func minOperations(a []int, k int) int {
	sort.Ints(a)
	var b []int

	for i, j, count, x, y := 0, 0, 0, 0, 0; ; count++ {
		if i < len(a) && (j >= len(b) || a[i] <= b[j]) {
			x = a[i]
			i++
		} else {
			x = b[j]
			j++
		}
		if !(k > x) {
			return count
		}
		if i < len(a) && (j >= len(b) || a[i] <= b[j]) {
			y = a[i]
			i++
		} else {
			y = b[j]
			j++
		}
		t := 2*x + y
		if t < k {
			b = append(b, t)
		} else {
			b = append(b, k)
		}
	}
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minOperations0(nums []int, k int) int {
	//q := &MinHeap{}
	//heap.Init(q)
	//for _, num := range nums {
	//	if num < k {
	//		heap.Push(q, num)
	//	}
	//}

	t := MinHeap(nums)
	q := &t
	heap.Init(q)

	res := 0
	for q.Len() > 1 {
		x := heap.Pop(q).(int)
		if x >= k {
			break
		}
		y := heap.Pop(q).(int)
		num := min(x, y)*2 + max(x, y)
		if num < k {
			heap.Push(q, num)
		}
		res++
	}
	if q.Len() > 0 {
		num := heap.Pop(q).(int)
		if num < k {
			res++
		}
	}
	return res
}
