package p2054

import "sort"

func maxTwoEvents(events [][]int) int {
	e := make([][]int, len(events))
	copy(e, events)

	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	sort.Slice(e, func(i, j int) bool {
		return e[i][2] > e[j][2]
	})

	index := 0
	res := e[0][2]
	for i := 0; i < len(events); i++ {
		for ; index < len(e) && e[index][0] <= events[i][1]; index++ {
		}
		if index == len(e) {
			return res
		}
		res = max(res, events[i][2]+e[index][2])
	}

	return res
}

// OutOfMemory
func maxTwoEvents0(events [][]int) int {

	maxEndTime := 0
	for _, event := range events {
		maxEndTime = max(maxEndTime, event[1])
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	dp := make([][2]int, maxEndTime+2)
	index := 0
	for i := 1; i < maxEndTime+2; i++ {
		for ; index < len(events) && events[index][1] == i; index++ {
			startTime := events[index][0]
			endTime := events[index][1]
			value := events[index][2]
			dp[endTime][0] = max(dp[endTime][0], value)
			dp[endTime][1] = max(dp[endTime][1], dp[startTime-1][0]+value)
		}
		dp[i][0] = max(dp[i][0], dp[i-1][0])
		dp[i][1] = max(dp[i][1], dp[i-1][1])
	}
	return dp[len(dp)-1][1]
}
