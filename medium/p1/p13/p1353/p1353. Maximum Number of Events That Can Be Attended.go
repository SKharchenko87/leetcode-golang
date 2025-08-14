package p1353

import (
	"container/heap"
	"slices"
	"sort"
)

type minHeap [][]int

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	if h[i][1] == h[j][1] {
		return h[i][0] < h[j][0]
	}
	return h[i][1] < h[j][1]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]

}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h minHeap) Peek() interface{} {
	if h.Len() == 0 {
		return nil
	}
	return h[0]
}

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	maxEndDayEvent := slices.MaxFunc(events, func(a, b []int) int {
		return max(a[1], b[1])
	})

	maxEndDay := maxEndDayEvent[1]

	attendedEventsCount := 0
	currentDay := events[0][0]
	hp := &minHeap{}
	heap.Init(hp)

	currentIndex := 0
	for currentDay <= maxEndDay || hp.Len() > 0 {
		for currentIndex < len(events) && events[currentIndex][0] <= currentDay {
			heap.Push(hp, events[currentIndex])
			currentIndex++
		}

		for hp.Len() > 0 && hp.Peek().([]int)[1] < currentDay {
			heap.Pop(hp)
		}

		if hp.Len() > 0 {
			_ = heap.Pop(hp).([]int)
			attendedEventsCount++
			currentDay++
		} else {
			if currentIndex < len(events) {
				currentDay = events[currentIndex][0]
			} else {
				// Если нет доступных событий и больше нет будущих событий, то завершаем.
				break
			}
		}
	}

	return attendedEventsCount
}

// [[1,4],[4,4],[2,2],[3,4],[1,1]]
// 0 1 2 3 4
//   _
//     _
//   _______
//       ___
//         _

// [[1,2],[1,2],[3,3],[1,5],[1,5]]
// 0 1 2 3 4 5
//   ___
//   ___
//       _
//   _________
//   _________
