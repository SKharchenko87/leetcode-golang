package p3440

import "sort"

type free struct {
	start, end int
}

func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	n := len(startTime)
	freeTimes := make([]free, n+1)
	maxTimes := make([]int, n)
	preTime := 0

	for i := 0; i < n-1; i++ {
		c := startTime[i+1] - preTime - (endTime[i] - startTime[i])
		maxTimes[i] = c
		freeTimes[i].end = startTime[i]
		freeTimes[i].start = preTime
		preTime = endTime[i]
	}
	//
	c := eventTime - preTime - (endTime[n-1] - startTime[n-1])
	maxTimes[n-1] = c

	freeTimes[n-1].end = startTime[n-1]
	freeTimes[n-1].start = preTime
	preTime = endTime[n-1]

	freeTimes[n].end = eventTime
	freeTimes[n].start = preTime

	sort.Slice(freeTimes, func(i, j int) bool {
		return freeTimes[i].end-freeTimes[i].start > freeTimes[j].end-freeTimes[j].start
	})

	res := 0
	for i := 0; i < n; i++ {
		currTime := endTime[i] - startTime[i]
		index := 0
		freeTime := freeTimes[index].end - freeTimes[index].start
		for {
			if freeTime < currTime {
				res = max(res, maxTimes[i])
				break
			}
			if freeTimes[index].end != startTime[i] && freeTimes[index].start != endTime[i] {
				res = max(res, maxTimes[i]+currTime)
				break
			}
			index++
			freeTime = freeTimes[index].end - freeTimes[index].start
		}
	}
	return res
}
