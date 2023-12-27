package p1578

func minCost(colors string, neededTime []int) int {
	var sum, maxCurrent int
	for i := 1; i < len(colors); i++ {
		if colors[i-1] == colors[i] {
			if maxCurrent == 0 {
				sum, maxCurrent = sum+neededTime[i-1], neededTime[i-1]
			}
			sum += neededTime[i]
			maxCurrent = max(maxCurrent, neededTime[i])
		} else {
			sum, maxCurrent = sum-maxCurrent, 0
		}
	}
	return sum - maxCurrent
}
