package p2940

import (
	"container/heap"
	"slices"
)

type Ints interface {
	int | int8 | int16 | int32 | int64
}
type Floats interface {
	float32 | float64
}
type Numeric interface{ Ints | Floats }

func max[T Numeric](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// TLE
func leftmostBuildingQueries0(heights []int, queries [][]int) []int {
	r := make([]int, len(heights))
	for i := range r {
		j := i
		for ; j < len(heights) && heights[j] <= heights[i]; j++ {
		}
		if j == len(heights) {
			r[i] = -1
		} else {
			r[i] = j
		}
	}

	result := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		if queries[i][0] == queries[i][1] {
			result[i] = queries[i][0]
		} else if queries[i][0] < queries[i][1] && heights[queries[i][0]] < heights[queries[i][1]] {
			result[i] = queries[i][1]
		} else if queries[i][1] < queries[i][0] && heights[queries[i][1]] < heights[queries[i][0]] {
			result[i] = queries[i][0]
		} else {
			maxIndex := max(queries[i][0], queries[i][1])
			maxValue := max(heights[queries[i][0]], heights[queries[i][1]])
			for maxIndex != -1 && heights[maxIndex] < maxValue+1 {
				maxIndex = r[maxIndex]
			}
			result[i] = maxIndex
		}
	}

	return result
}

type MaxQuery struct {
	value    int
	index    int
	queryIdx int
}

type MaxQueryHeap []MaxQuery

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	results := make([]int, len(queries))
	deferredQueries := make([]MaxQuery, 0)

	// Prepare queries
	for idx, q := range queries {
		l, r := q[0], q[1]
		if l > r {
			l, r = r, l
		}
		if l == r || heights[r] > heights[l] {
			results[idx] = r
		} else {
			results[idx] = -1
			deferredQueries = append(deferredQueries, MaxQuery{heights[l] + 1, r, idx})
		}
	}

	// Sort deferred queries by index
	slices.SortFunc(deferredQueries, func(a, b MaxQuery) int {
		return a.index - b.index
	})

	// Process queries using a heap
	h := &MaxQueryHeap{}
	for currentBuilding := range heights {
		// Add all queries up to the current position
		for len(deferredQueries) > 0 && deferredQueries[0].index <= currentBuilding {
			heap.Push(h, deferredQueries[0])
			deferredQueries = deferredQueries[1:]
		}

		// Resolve queries in the heap
		for h.Len() > 0 && (*h)[0].value <= heights[currentBuilding] {
			topQuery := heap.Pop(h).(MaxQuery)
			results[topQuery.queryIdx] = currentBuilding
		}
	}
	return results
}

func (h MaxQueryHeap) Len() int           { return len(h) }
func (h MaxQueryHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h MaxQueryHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxQueryHeap) Push(x any) {
	*h = append(*h, x.(MaxQuery))
}

func (h *MaxQueryHeap) Pop() any {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}
