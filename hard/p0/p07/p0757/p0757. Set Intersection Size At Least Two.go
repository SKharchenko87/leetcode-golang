package p0757

import (
	"slices"
	"sort"
)

/*
123456789
---
  -----
       --


{1, 3}, {1, 4}, {2, 5}, {3, 5}
123456789
---
----
 ----
  ---


*/

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	lineVertical := []int{intervals[0][1] - 1, intervals[0][1]}
	var interval = make([]int, 2)
	for i := 1; i < len(intervals); i++ {
		interval = intervals[i]
		index, _ := slices.BinarySearch(lineVertical, interval[0])
		cnt := 0
		for ; index < len(lineVertical) && interval[0] <= lineVertical[index] && lineVertical[index] <= interval[1] && cnt < 2; index++ {
			cnt++
		}
		if cnt == 2 {
			continue
		}
		if cnt == 0 {
			lineVertical = append(lineVertical, interval[1]-1)
		}
		lineVertical = append(lineVertical, interval[1])
	}
	return len(lineVertical)
}
