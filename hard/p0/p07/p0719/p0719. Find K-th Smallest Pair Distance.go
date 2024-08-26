package p0719

import (
	"container/heap"
	"fmt"
	"slices"
	"sort"
)

type MinHeapK struct {
	heap []int
	k    int
}

func Constructor(k int) MinHeapK {
	return MinHeapK{heap: make([]int, 0, k), k: k}
}

func (h MinHeapK) Len() int            { return len(h.heap) }
func (h MinHeapK) Less(i, j int) bool  { return h.heap[i] > h.heap[j] }
func (h MinHeapK) Swap(i, j int)       { h.heap[i], h.heap[j] = h.heap[j], h.heap[i] }
func (h *MinHeapK) Push(x interface{}) { (*h).heap = append((*h).heap, x.(int)) }

func (h *MinHeapK) Pop() interface{} {
	old := *h
	n := len(old.heap)
	x := old.heap[n-1]
	(*h).heap = old.heap[0 : n-1]
	return x
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, nums[len(nums)-1]

	for left <= right {
		mid := (left + right) / 2
		cc := 0

		j := 1
		for i := 0; i < len(nums)-1; i++ {
			j = max(j, i+1)
			for j < len(nums) && nums[j]-nums[i] <= mid {
				j++
			}
			cc += j - 1 - i
		}

		if cc < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func smallestDistancePair3(nums []int, k int) int {
	sort.Ints(nums)
	l := len(nums)
	arr := make([]int, 0, k)
	for i := 0; i < l; i++ {
		for j := i + 1; j < i+1+k && j < l; j++ {
			v := nums[j] - nums[i]
			index, _ := slices.BinarySearch(arr, v)
			arr = slices.Insert(arr, index, v)
		}
	}
	fmt.Println(arr)
	return arr[k-1]
}

func smallestDistancePair2(nums []int, k int) int {
	sort.Ints(nums)
	l := len(nums)
	arr := make([]int, 0, k)
	for i := 0; i < l; i++ {
		for j := i + 1; j < i+1+k && j < l; j++ {
			v := abs(nums[i] - nums[j])
			index, _ := slices.BinarySearch(arr, v)
			arr = slices.Insert(arr, index, v)
		}
	}
	fmt.Println(arr)
	return arr[k-1]
}

func smallestDistancePair1(nums []int, k int) int {
	sort.Ints(nums)
	l := len(nums)

	//arr := make([]int, 0, 3*l*(l-1)/4)
	//arr := make([]int, 0, 5*l*(l-1)/8)
	//arr := make([]int, 0, 9*l*(l-1)/16)
	arr := make([]int, 0, l*(l-1)/2)
	//for i := 0; i < l*(l-1)/2; i++ {
	//	arr[i] = 10000
	//}
	for i := 0; i < l; i++ {
		for j := i + 1; j < i+1+k && j < l; j++ {
			v := abs(nums[i] - nums[j])
			index, _ := slices.BinarySearch(arr, v)
			arr = slices.Insert(arr, index, v)
		}
	}
	//fmt.Println(arr)
	return arr[k-1]
}

func smallestDistancePair0(nums []int, k int) int {
	mhk := Constructor(k)
	heap.Init(&mhk)
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			heap.Push(&mhk, abs(nums[i]-nums[j]))
			if mhk.Len() > k {
				heap.Pop(&mhk)
			}
		}
	}
	return heap.Pop(&mhk).(int)
}
