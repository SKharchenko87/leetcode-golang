package p0539

import (
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	minutes := [1440]bool{}
	for _, timePoint := range timePoints {
		minute := getMinute(timePoint)
		if minutes[minute] {
			return 0
		}
		minutes[minute] = true
	}

	minMinutes := 1440
	startPosition := 0
	for !minutes[startPosition] {
		startPosition++
	}
	first := startPosition

	for i := startPosition + 1; i < 1440; i++ {
		if minutes[i] {
			minMinutes = min(i-startPosition, minMinutes)
			startPosition = i
		}
	}
	minMinutes = min(first+1440-startPosition, minMinutes)
	return minMinutes
}

func findMinDifference0(timePoints []string) int {
	l := len(timePoints)
	sort.Strings(timePoints)
	minTimeDiff := getMinute(timePoints[0]) + 1440 - getMinute(timePoints[l-1])
	for i := 0; i < l-1; i++ {
		if c := getMinute(timePoints[i+1]) - getMinute(timePoints[i]); minTimeDiff > c {
			minTimeDiff = c
		}
	}
	return minTimeDiff
}

func getMinute(time string) int {
	return int(time[0]-'0')*600 + int(time[1]-'0')*60 + int(time[3]-'0')*10 + int(time[4]-'0')
}

func getMinute0(time string) int {
	hour, _ := strconv.Atoi(time[:2])
	minus, _ := strconv.Atoi(time[3:])
	return hour*60 + minus
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
