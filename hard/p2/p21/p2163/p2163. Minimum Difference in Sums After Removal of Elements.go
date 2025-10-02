package p2163

import (
	"cmp"
	"container/heap"
	"slices"
)

type MinHeap []int64

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() int64        { return h[0] }

func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int64)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int64

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() int64        { return h[0] }

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int64)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Pair struct {
	firstVal, secondVal int64
}

func minimumDifference(nums []int) int64 {
	var sumFirst, sumSecond int64
	n := len(nums) / 3
	dp := make([]Pair, n+1)

	maxHeap := make(MaxHeap, 0, n)
	heap.Init(&maxHeap)
	for i := 0; i < n; i++ {
		sumFirst += int64(nums[i])
		heap.Push(&maxHeap, int64(nums[i]))
	}
	dp[0].firstVal = sumFirst
	for i := 0; i < n; i++ {
		if maxHeap.Peek() >= int64(nums[n+i]) {
			sumFirst += int64(nums[n+i])
			x := heap.Pop(&maxHeap).(int64)
			sumFirst -= x
			heap.Push(&maxHeap, int64(nums[n+i]))
		}
		dp[i+1].firstVal = sumFirst
	}

	minHeap := make(MinHeap, 0, n)
	for i := 0; i < n; i++ {
		index := 3*n - 1 - i
		sumSecond += int64(nums[index])
		heap.Push(&minHeap, int64(nums[index]))
	}
	dp[n].secondVal = sumSecond
	for i := 0; i < n; i++ {
		index := 2*n - 1 - i
		if minHeap.Peek() <= int64(nums[index]) {
			sumSecond += int64(nums[index])
			x := heap.Pop(&minHeap).(int64)
			sumSecond -= x
			heap.Push(&minHeap, int64(nums[index]))
		}
		dp[n-1-i].secondVal = sumSecond
	}

	res := slices.MinFunc(dp, func(a, b Pair) int {
		return cmp.Compare(a.firstVal-a.secondVal, b.firstVal-b.secondVal)
	})
	return res.firstVal - res.secondVal
}

// 16, 46, 43, 41, 42, 14, 36, 49, 50, 28, 38, 25, 17, 5, 18, 11, 14, 21, 23, 39, 23
// 14, 16, 36, 41, 42, 43, 46                                                         = 238
//                                                        11, 14, 18, 21, 23, 23, 39  = 149
//                           218  218 232
