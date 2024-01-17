package p2542

import (
	"container/heap"
	"sort"
)

// ToDo разобраться
func maxScore(nums1 []int, nums2 []int, k int) int64 {
	magic := make([][2]int, len(nums1))
	for i := range magic {
		magic[i] = [2]int{nums1[i], nums2[i]}
	}
	sort.Slice(magic, func(i, j int) bool {
		return magic[i][1] > magic[j][1]
	})
	minHeap := &Heap{}
	summation := 0
	for i := 0; i < k; i++ {
		heap.Push(minHeap, magic[i][0])
		summation += magic[i][0]
	}
	factor := magic[k-1][1]
	ms := summation * factor
	for i := k; i < len(magic); i++ {
		summation += magic[i][0]
		heap.Push(minHeap, magic[i][0])
		summation -= heap.Pop(minHeap).(int)
		cms := magic[i][1] * summation
		if cms > ms {
			ms = cms
		}
	}
	return int64(ms)
}

type Heap []int

func (p Heap) Len() int            { return len(p) }
func (p Heap) Less(i, j int) bool  { return p[i] < p[j] }
func (p *Heap) Swap(i, j int)      { (*p)[i], (*p)[j] = (*p)[j], (*p)[i] }
func (p *Heap) Push(i interface{}) { *p = append(*p, i.(int)) }
func (p *Heap) Pop() interface{}   { v := (*p)[len(*p)-1]; *p = (*p)[:len(*p)-1]; return v }
