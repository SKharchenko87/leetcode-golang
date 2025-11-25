package p1792

import (
	"container/heap"
	"math"
)

type Class struct {
	pass          float64
	total         float64
	increaseRatio float64
}

type maxHeap []Class

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	return h[i].increaseRatio > h[j].increaseRatio
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(Class))
}

func (h *maxHeap) Pop() any {
	l := h.Len()
	x := (*h)[l-1]
	*h = (*h)[:l-1]
	return x
}

const actual = 10e4

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	classSize := len(classes)
	arr := make([]Class, 0, classSize)
	mh := maxHeap(arr)
	heap.Init(&mh)
	for _, class := range classes {
		pass := float64(class[0])
		total := float64(class[1])
		heap.Push(&mh, Class{pass, total, (pass+1)/(total+1) - pass/total})
	}
	for i := 0; i < extraStudents; i++ {
		x := heap.Pop(&mh).(Class)
		pass := x.pass + 1
		total := x.total + 1
		heap.Push(&mh, Class{pass, total, (pass+1)/(total+1) - pass/total})
	}
	res := float64(0)
	for mh.Len() > 0 {
		x := heap.Pop(&mh).(Class)
		res += x.pass / x.total
	}
	return math.Round(res/float64(classSize)*actual) / actual
}
