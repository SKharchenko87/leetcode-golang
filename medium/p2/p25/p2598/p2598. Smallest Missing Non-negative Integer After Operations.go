package p2598

import (
	"container/heap"
)

func findSmallestInteger(nums []int, value int) int {
	l := len(nums)
	mex := make([]int, value)
	for i := 0; i < l; i++ {
		x := (nums[i]%value + value) % value
		mex[x]++
	}
	c := 0
	for {
		for i := 0; i < value; i++ {
			if mex[i] == 0 {
				return value*c + i
			}
			mex[i]--
		}
		c++
	}
}

func findSmallestInteger1(nums []int, value int) int {
	l := len(nums)
	mh := map[int]int{}
	for i := 0; i < l; i++ {
		x := (nums[i]%value + value) % value
		mh[x]++
	}

	index := 0
	for {
		if v, ok := mh[index]; ok {
			if v > 1 {
				mh[index+value] += v - 1
			}
		} else {
			return index
		}
		index++
	}
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findSmallestInteger0(nums []int, value int) int {
	l := len(nums)
	mh := minHeap(make([]int, 0, l))
	heap.Init(&mh)
	for i := 0; i < l; i++ {
		x := (nums[i]%value + value) % value
		heap.Push(&mh, x)
	}
	prev := 0
	cnt := 1
	x := heap.Pop(&mh).(int)
	if x != 0 {
		return 0
	}
	prev = x
	for mh.Len() > 0 {
		x = heap.Pop(&mh).(int)
		if x > prev && x != prev+1 {
			return prev + 1
		} else if x == prev {
			y := value*cnt + prev
			heap.Push(&mh, y)
			cnt++
		} else {
			cnt = 1
		}
		prev = x
	}
	return x + 1
}
