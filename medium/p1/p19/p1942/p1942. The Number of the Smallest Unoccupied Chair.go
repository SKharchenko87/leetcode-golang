package p1942

import (
	"container/heap"
	"slices"
	"sort"
)

type Time struct {
	arrival int
	leave   int
	index   int
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type LeavingQueue []Time

func (h LeavingQueue) Len() int           { return len(h) }
func (h LeavingQueue) Less(i, j int) bool { return h[i].leave < h[j].leave }
func (h LeavingQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *LeavingQueue) Push(x interface{}) {
	*h = append(*h, x.(Time))
}

func (h *LeavingQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func smallestChair(times [][]int, targetFriend int) int {
	//targetArrival := times[targetFriend][0]
	timeList := make([]Time, len(times))

	for i, t := range times {
		timeList[i] = Time{arrival: t[0], leave: t[1], index: i}
	}

	sort.Slice(timeList, func(i, j int) bool {
		if timeList[i].arrival == timeList[j].arrival {
			return timeList[i].leave < timeList[j].leave
		}
		return timeList[i].arrival < timeList[j].arrival
	})

	nextChair := 0
	availableChairs := &MinHeap{}
	leavingQueue := &LeavingQueue{}

	for _, time := range timeList {
		arrival, leave, index := time.arrival, time.leave, time.index

		// Free up chairs based on current time
		for leavingQueue.Len() > 0 && (*leavingQueue)[0].leave <= arrival {
			leavingTime := heap.Pop(leavingQueue).(Time)
			heap.Push(availableChairs, leavingTime.index)
		}

		var currentChair int
		if availableChairs.Len() > 0 {
			currentChair = heap.Pop(availableChairs).(int)
		} else {
			currentChair = nextChair
			nextChair++
		}

		// Push current leave time and chair
		heap.Push(leavingQueue, Time{leave: leave, index: currentChair})

		// Check if it's the target friend
		if index == targetFriend {
			return currentChair
		}
	}

	return 0
}

type time struct {
	open  *[]int
	close *[]int
}

type PriorityHeap []int

func (h PriorityHeap) Len() int { return len(h) }

func (h PriorityHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h PriorityHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *PriorityHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *PriorityHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func smallestChair0(times [][]int, targetFriend int) int {
	length := len(times)
	timesLineMap := make(map[int]time)
	timesLine := make([]int, 0, length)
	for i := 0; i < length; i++ {
		if times[i][0] <= times[targetFriend][0] {
			if _, ok := timesLineMap[times[i][0]]; !ok {
				timesLineMap[times[i][0]] = time{&[]int{}, &[]int{}}
			}
			arr := timesLineMap[times[i][0]].open
			*arr = append(*arr, i)
			timesLine = append(timesLine, times[i][0])

			if _, ok := timesLineMap[times[i][1]]; !ok {
				timesLineMap[times[i][1]] = time{&[]int{}, &[]int{}}
			}

			arr = timesLineMap[times[i][1]].close
			*arr = append(*arr, i)
			timesLine = append(timesLine, times[i][1])
		}
	}
	sort.Ints(timesLine)
	friends := make([]int, length)
	chair := -1
	ph := &PriorityHeap{}
	heap.Init(ph)
	for _, tl := range timesLine {
		for _, c := range *timesLineMap[tl].close {
			heap.Push(ph, friends[c])
		}
		for _, o := range *timesLineMap[tl].open {
			if ph.Len() > 0 {
				friends[o] = heap.Pop(ph).(int)
			} else {
				chair++
				friends[o] = chair
			}
		}
	}
	return friends[targetFriend]
}

// TLE
func smallestChair1(times [][]int, targetFriend int) int {
	length := len(times)
	timesLineMap := make(map[int]time)
	timesLine := make([]int, 0, length)
	for i := 0; i < length; i++ {
		if times[i][0] <= times[targetFriend][0] {
			if _, ok := timesLineMap[times[i][0]]; !ok {
				timesLineMap[times[i][0]] = time{&[]int{}, &[]int{}}
			}
			arr := timesLineMap[times[i][0]].open
			*arr = append(*arr, i)
			timesLine = append(timesLine, times[i][0])

			if _, ok := timesLineMap[times[i][1]]; !ok {
				timesLineMap[times[i][1]] = time{&[]int{}, &[]int{}}
			}

			arr = timesLineMap[times[i][1]].close
			*arr = append(*arr, i)
			timesLine = append(timesLine, times[i][1])
		}
	}
	sort.Ints(timesLine)
	friends := make([]int, length)
	chair := -1
	ph := make([]int, 0, length)
	for _, tl := range timesLine {
		for _, c := range *timesLineMap[tl].close {
			index, _ := slices.BinarySearch(ph, friends[c])
			ph = slices.Insert(ph, index, friends[c])
		}
		for _, o := range *timesLineMap[tl].open {
			if len(ph) > 0 {
				friends[o] = ph[0]
				ph = slices.Delete(ph, 0, 1)
			} else {
				chair++
				friends[o] = chair
			}
		}
	}
	return friends[targetFriend]
}
