package p1578

func minCost(colors string, neededTime []int) int {
	var prev byte
	result := 0
	for i := 0; i < len(neededTime); i++ {
		if colors[i] == prev {
			if neededTime[i] <= neededTime[i-1] {
				result += neededTime[i]
				neededTime[i] = neededTime[i-1]
			} else {
				result += neededTime[i-1]
			}
		}
		prev = colors[i]
	}
	return result
}

func minCost2(colors string, neededTime []int) int {
	var prev byte
	result := 0
	for i := len(neededTime) - 1; i >= 0; i-- {
		if colors[i] == prev {
			if neededTime[i] <= neededTime[i+1] {
				result += neededTime[i]
				neededTime[i] = neededTime[i+1]
			} else {
				result += neededTime[i+1]
			}
		}
		prev = colors[i]
	}
	return result
}

func minCost1(colors string, neededTime []int) int {
	var prev byte
	result := 0
	for i := len(neededTime) - 1; i >= 0; i-- {
		if colors[i] == prev {
			result += min(neededTime[i], neededTime[i+1])
			m := max(neededTime[i], neededTime[i+1])
			neededTime[i] = m
			neededTime[i+1] = m
		}
		prev = colors[i]
	}
	return result
}

func minCost0(colors string, neededTime []int) int {
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
