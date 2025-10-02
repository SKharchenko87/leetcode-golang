package main

import (
	"container/heap"
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	res := make([][]int, len(nums)/3)
	for i := 0; i < len(nums); i = i + 3 {
		if nums[i+1]-nums[i] > k || nums[i+2]-nums[i] > k {
			return [][]int{}
		}
		res[i/3] = []int{nums[i], nums[i+1], nums[i+2]}
	}
	return res
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
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

func divideArray2(nums []int, k int) [][]int {
	h := &minHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}

	res := make([][]int, 0, len(nums)/3)
	for h.Len() > 0 {
		x0 := heap.Pop(h).(int)
		x1 := heap.Pop(h).(int)
		x2 := heap.Pop(h).(int)
		if x2-x0 > k {
			return [][]int{}
		}
		res = append(res, []int{x0, x1, x2})
	}
	return res

}

func divideArray1(nums []int, k int) [][]int {
	l := len(nums)
	sort.Ints(nums)
	res := make([][]int, l/3)
	for i := 0; i < l/3; i++ {
		res[i] = make([]int, 3)
		res[i][0] = nums[i*3+0]
		res[i][1] = nums[i*3+1]
		res[i][2] = nums[i*3+2]
		if res[i][2]-res[i][0] > k {
			return [][]int{}
		}
	}
	return res
}

func divideArray0(nums []int, k int) [][]int {
	freq := make([]int, 10001)
	minVal := nums[0]
	maxVal := nums[0]
	for _, num := range nums {
		freq[num]++
		minVal = min(minVal, num)
		maxVal = max(maxVal, num)
	}

	l := len(nums)
	res := make([][]int, l/3)
	index := -1
	for i := minVal; i <= maxVal; i++ {
		cnt := freq[i]
		//for i, cnt := range freq {
		for cnt > 0 {
			index++
			if index%3 == 0 {
				res[index/3] = make([]int, 3)
			}
			res[index/3][index%3] = i
			cnt--
			if index%3 == 2 && res[index/3][2]-res[index/3][0] > k {
				return [][]int{}
			}
		}
	}
	return res
}
