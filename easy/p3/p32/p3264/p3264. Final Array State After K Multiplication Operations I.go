package p3264

import "container/heap"

type Node struct {
	value int
	index int
}

type MinHeap []Node

func (mh MinHeap) Len() int {
	return len(mh)
}
func (mh MinHeap) Less(i, j int) bool {
	if mh[i].value == mh[j].value {
		return mh[i].index < mh[j].index
	}
	return mh[i].value < mh[j].value
}
func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}
func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.(Node))
}
func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}

func getFinalState(nums []int, k int, multiplier int) []int {
	mh := MinHeap{}
	heap.Init(&mh)
	for index, value := range nums {
		heap.Push(&mh, Node{value, index})
	}
	for i := 0; i < k; i++ {
		node := heap.Pop(&mh).(Node)
		nums[node.index] = node.value * multiplier
		heap.Push(&mh, Node{node.value * multiplier, node.index})
	}
	return nums
}
