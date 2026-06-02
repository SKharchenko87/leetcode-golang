package p3633

import "math"

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	n := len(landStartTime)
	m := len(waterStartTime)
	res := math.MaxInt32
	for i := range n {
		land := landStartTime[i] + landDuration[i]
		for j := range m {
			res = min(res, max(land, waterStartTime[j])+waterDuration[j])
		}
	}
	for j := range m {
		water := waterStartTime[j] + waterDuration[j]
		for i := range n {
			res = min(res, max(water, landStartTime[i])+landDuration[i])
		}
	}
	return res
}
