package p3505

import (
	"container/heap"
	"math"
)

type MaxHeap []int

func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Push(x any)         { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h *MaxHeap) peek() int { return (*h)[0] }

type MinHeap []int

func (h *MinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap) Push(x any)         { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h *MinHeap) peek() int { return (*h)[0] }

func minOperations(nums []int, x, k int) int64 {
	n := len(nums)
	upper, lower := &MinHeap{}, &MaxHeap{}
	heap.Init(upper)
	heap.Init(lower)

	// Отдельные мапы удалений для каждой кучи!
	delLower := make(map[int]int)
	delUpper := make(map[int]int)

	var sumLeft, sumRight int64
	sizeLeft, sizeRight := 0, 0

	pruneLower := func() {
		for lower.Len() > 0 && delLower[lower.peek()] > 0 {
			delLower[lower.peek()]--
			heap.Pop(lower)
		}
	}
	pruneUpper := func() {
		for upper.Len() > 0 && delUpper[upper.peek()] > 0 {
			delUpper[upper.peek()]--
			heap.Pop(upper)
		}
	}

	rebalance := func() {
		for sizeLeft > (x+1)/2 {
			pruneLower()
			val := heap.Pop(lower).(int)
			sumLeft -= int64(val)
			sizeLeft--
			heap.Push(upper, val)
			sumRight += int64(val)
			sizeRight++
		}
		for sizeLeft < (x+1)/2 && sizeRight > 0 {
			pruneUpper()
			val := heap.Pop(upper).(int)
			sumRight -= int64(val)
			sizeRight--
			heap.Push(lower, val)
			sumLeft += int64(val)
			sizeLeft++
		}
		pruneLower()
		pruneUpper()
	}

	// Считаем первоначальное окно
	for i := 0; i < x; i++ {
		heap.Push(lower, nums[i])
		sumLeft += int64(nums[i])
		sizeLeft++
	}
	rebalance()

	// Проходим по всем окнам и сохраняем результат
	xWin := make([]int64, n-x+1)
	median := int64(lower.peek())
	xWin[0] = int64((x+1)/2)*median - sumLeft + sumRight - int64(x/2)*median

	for i := x; i < n; i++ {
		in, out := nums[i], nums[i-x]

		// 1. Добавляем новый элемент
		if in <= lower.peek() {
			heap.Push(lower, in)
			sumLeft += int64(in)
			sizeLeft++
		} else {
			heap.Push(upper, in)
			sumRight += int64(in)
			sizeRight++
		}

		// 2. Удаляем старый элемент из ТОЙ кучи, где он лежит
		if out <= lower.peek() {
			delLower[out]++
			sumLeft -= int64(out)
			sizeLeft--
		} else {
			delUpper[out]++
			sumRight -= int64(out)
			sizeRight--
		}

		rebalance()
		median = int64(lower.peek())
		xWin[i-x+1] = int64((x+1)/2)*median - sumLeft + sumRight - int64(x/2)*median
	}

	// ДП для выбора k непересекающихся окон длины x
	const inf = math.MaxInt64 / 4
	dp := make([]int64, n+1)
	nextDp := make([]int64, n+1)

	for j := 1; j <= k; j++ {
		for i := 0; i <= n; i++ {
			nextDp[i] = inf
		}
		for i := j * x; i <= n; i++ {
			nextDp[i] = nextDp[i-1]
			if prev := dp[i-x] + xWin[i-x]; prev < nextDp[i] {
				nextDp[i] = prev
			}
		}
		copy(dp, nextDp)
	}

	return dp[n]
}
