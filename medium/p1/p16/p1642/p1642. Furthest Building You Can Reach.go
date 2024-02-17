package p1642

import (
	"container/heap"
	"fmt"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// ToDo разобраться
func furthestBuilding(heights []int, bricks, ladders int) int {
	minHeap := &MinHeap{}

	for i := 1; i < len(heights); i++ {
		diff := heights[i] - heights[i-1]
		fmt.Println("diff=", diff)
		if diff <= 0 {
			continue
		}
		heap.Push(minHeap, diff)
		fmt.Println("heap=", minHeap)
		if minHeap.Len() > ladders {
			x := heap.Pop(minHeap).(int)
			fmt.Println("x=", x)
			bricks -= x
		}
		if bricks < 0 {
			return i - 1
		}
	}

	return len(heights) - 1
}

func furthestBuilding1(heights []int, bricks int, ladders int) int {
	arr := make([]int, len(heights)-1)
	sum := 0
	cnt := 0
	for i := range arr {
		tmp := heights[i+1] - heights[i]
		if tmp > 0 {
			arr[i] = tmp
			sum += tmp
			cnt++
		}
	}
	if cnt < ladders || sum <= bricks {
		return len(arr)
	}
	//fmt.Println(arr, sum)
	i := len(arr)
	for ; i >= 0; i-- {
		if i < len(arr) {
			sum -= arr[i]
		}

		arrtmp := make([]int, i)
		copy(arrtmp, arr[:i])
		sort.Ints(arrtmp)
		x := len(arrtmp) - ladders
		if x < 0 {
			break
		}
		arrtmp = arrtmp[:x]
		coef2 := 0
		for j := x; j < x+ladders; j++ {
			coef2 += arrtmp[j-1]
		}

		sumtmp := sum - coef2
		fmt.Println(arr[:i], sumtmp, sum, arrtmp)

		if sumtmp <= bricks {
			break
		}
	}
	return i
}
