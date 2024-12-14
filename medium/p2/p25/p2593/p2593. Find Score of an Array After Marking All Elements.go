package p2593

import (
	"container/heap"
	"sort"
)

type Node struct {
	val, index int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].val == pq[j].val {
		return pq[i].index < pq[j].index
	}
	return pq[i].val < pq[j].val
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func findScore(nums []int) int64 {
	nodes := make([]*Node, len(nums))
	for i, n := range nums {
		nodes[i] = &Node{val: n, index: i}
	}
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].val == nodes[j].val {
			return nodes[i].index < nodes[j].index
		}
		return nodes[i].val < nodes[j].val
	})
	res := int64(0)
	marks := make(map[int]struct{}, len(nums)+2)
	for _, node := range nodes {
		if _, ok := marks[node.index]; !ok {
			res += int64(node.val)
			marks[node.index-1] = struct{}{}
			marks[node.index] = struct{}{}
			marks[node.index+1] = struct{}{}
		}
	}
	return res
}

func findScore1(nums []int) int64 {
	nodes := make([]*Node, len(nums))
	for i, n := range nums {
		nodes[i] = &Node{val: n, index: i}
	}
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].val == nodes[j].val {
			return nodes[i].index < nodes[j].index
		}
		return nodes[i].val < nodes[j].val
	})
	res := int64(0)
	marks := make(map[int]bool, len(nums)+2)
	//marks := map[int]bool{}
	for _, node := range nodes {
		if !marks[node.index] {
			res += int64(node.val)
			marks[node.index-1] = true
			marks[node.index] = true
			marks[node.index+1] = true
		}
	}
	return res
}

func findScore0(nums []int) int64 {
	q := PriorityQueue{}
	heap.Init(&q)
	for i, n := range nums {
		heap.Push(&q, &Node{val: n, index: i})
	}
	res := int64(0)
	marks := map[int]struct{}{}
	for q.Len() > 0 {
		cur := heap.Pop(&q).(*Node)
		if _, ok := marks[cur.index]; !ok {
			res += int64(cur.val)
			marks[cur.index-1] = struct{}{}
			marks[cur.index] = struct{}{}
			marks[cur.index+1] = struct{}{}
		}
	}
	return res
}
