package p0435

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Intervals struct {
	overlap map[int]interface{}
	length  int
	cnt     int
}

func (this *Intervals) setLength(length int) {
	this.length = length
}

func (this *Intervals) String() string {
	sb := strings.Builder{}
	sb.WriteString("{")
	for k, _ := range this.overlap {
		sb.WriteString(strconv.Itoa(k))
		sb.WriteString(" ")
	}
	sb.WriteString("}")
	return fmt.Sprintf("{overlap:%s length:%s cnt:%s}", sb.String(), strconv.Itoa(this.length), strconv.Itoa(this.cnt))
}

func eraseOverlapIntervals(intervals [][]int) int {
	l := len(intervals)
	count := -1
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	last := intervals[0][1]
	for i := 0; i < l; i++ {
		currentFirst := intervals[i][0]
		currentLast := intervals[i][1]
		if currentFirst < last {
			count++
			last = min(last, currentLast)
		} else {
			last = currentLast
		}
	}
	return count
}
