package p0135

import (
	"container/heap"
	"sort"
)

func candy(ratings []int) int {
	l := len(ratings)
	if l < 2 {
		return l
	}
	candies := make([]int, l)
	candies[0] = 1
	candies[l-1] = 1
	for i := 1; i < l; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}
	for i := l - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i+1]+1, candies[i])
		}
	}
	res := 0
	for i := 0; i < l; i++ {
		res += candies[i]
	}
	return res
}

type rank struct {
	index  int
	weight int
}

type minHeap []rank

func (mh minHeap) Len() int {
	return len(mh)
}

func (mh minHeap) Less(i, j int) bool {
	return mh[i].weight < mh[j].weight
}

func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *minHeap) Push(x any) {
	*mh = append(*mh, x.(rank))
}

func (mh *minHeap) Pop() any {
	l := mh.Len()
	res := (*mh)[l-1]
	*mh = (*mh)[:l-1]
	return res
}

func candy2(ratings []int) int {
	l := len(ratings)
	if l < 2 {
		return l
	}
	mh := &minHeap{}
	heap.Init(mh)
	for i, rating := range ratings {
		heap.Push(mh, rank{i, rating})
	}
	candies := make([]int, l)
	res := 0
	for mh.Len() > 0 {
		x := heap.Pop(mh).(rank)
		left, right := -1, -1
		if x.index > 0 {
			if ratings[x.index-1] == x.weight {
				left = 1
			} else {
				left = candies[x.index-1] + 1
			}
		}
		if x.index < l-1 {
			if ratings[x.index+1] == x.weight {
				right = 1
			} else {
				right = candies[x.index+1] + 1
			}
		}
		candies[x.index] = max(left, right)
		res += candies[x.index]
	}
	return res
}

type twoSlice struct {
	indexes []int
	data    []int
}

func (t twoSlice) Len() int {
	return len(t.data)
}

func (t twoSlice) Swap(i, j int) {
	t.indexes[i], t.indexes[j] = t.indexes[j], t.indexes[i]
	t.data[i], t.data[j] = t.data[j], t.data[i]
}

func (t twoSlice) Less(i, j int) bool {
	return t.data[i] < t.data[j]
}

func candy1(ratings []int) int {
	l := len(ratings)
	if l == 0 {
		return 0
	} else if l == 1 {
		return 1
	}

	d := make([]int, l)
	copy(d, ratings)
	indexes := make([]int, l)
	for i := 0; i < l; i++ {
		indexes[i] = i
	}
	sort.Sort(twoSlice{indexes, ratings})
	candies := make([]int, l)
	for _, index := range indexes {
		left, right := -1, -1
		if index > 0 {
			if d[index-1] == d[index] {
				left = 1
			} else {
				left = candies[index-1] + 1
			}
		}
		if index < l-1 {
			if d[index+1] == d[index] {
				right = 1
			} else {
				right = candies[index+1] + 1
			}
		}
		candies[index] = max(left, right)
	}
	res := 0
	for i := 0; i < l; i++ {
		res += candies[i]
	}
	return res
}

func candy0(ratings []int) int {
	l := len(ratings)
	if l == 0 {
		return 0
	} else if l == 1 {
		return 1
	}

	d := make([]int, l)
	copy(d, ratings)
	indexes := make([]int, l)
	for i := 0; i < l; i++ {
		indexes[i] = i
	}
	sort.Sort(twoSlice{indexes, ratings})
	candies := make([]int, l)
	for _, index := range indexes {
		if index == 0 {
			if d[index] == d[index+1] {
				candies[index] = 1
			} else {
				candies[index] = candies[index+1] + 1
			}
		} else if index == l-1 {
			if d[index] == d[index-1] {
				candies[index] = 1
			} else {
				candies[index] = candies[index-1] + 1
			}
		} else {
			left, right := candies[index-1]+1, candies[index+1]+1
			if d[index-1] == d[index] {
				left = 1
			}
			if d[index+1] == d[index] {
				right = 1
			}
			if d[index-1] == d[index] && d[index] == d[index+1] {
				right = 1
				left = 1
			}
			candies[index] = max(left, right)
		}
	}
	res := 0
	for i := 0; i < l; i++ {
		res += candies[i]
	}
	return res
}
