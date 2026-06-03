package p3635

import "math"

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	finishLand := finish(landStartTime, landDuration)
	finishWater := finish(waterStartTime, waterDuration)

	res := math.MaxInt
	earliestPossibleTime(&res, finishWater, landStartTime, landDuration)
	earliestPossibleTime(&res, finishLand, waterStartTime, waterDuration)

	return res
}

func finish(landStartTime []int, landDuration []int) int {
	res := math.MaxInt
	for i := 0; i < len(landStartTime); i++ {
		res = min(res, landStartTime[i]+landDuration[i])
	}
	return res
}

func earliestPossibleTime(res *int, finish int, start, duration []int) {
	for i := 0; i < len(duration); i++ {
		*res = min(*res, max(finish, start[i])+duration[i])
	}
}

func earliestFinishTime0(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	n, m := len(landStartTime), len(waterStartTime)
	finishLand := math.MaxInt
	for i := 0; i < n; i++ {
		finishLand = min(finishLand, landStartTime[i]+landDuration[i])
	}
	finishWater := math.MaxInt
	for i := 0; i < m; i++ {
		finishWater = min(finishWater, waterStartTime[i]+waterDuration[i])
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		res = min(res, max(finishWater, landStartTime[i])+landDuration[i])
	}
	for i := 0; i < m; i++ {
		res = min(res, max(finishLand, waterStartTime[i])+waterDuration[i])
	}
	return res
}
