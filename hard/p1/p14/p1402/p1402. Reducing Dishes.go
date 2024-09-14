package p1402

import "sort"

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	var sum, maxVal int
	for i := len(satisfaction) - 1; i >= 0; i-- {
		maxVal = max(maxVal, satisfaction[i]+maxVal+sum)
		sum += satisfaction[i]
	}
	return maxVal
}

func maxSatisfaction0(satisfaction []int) int {
	sort.Ints(satisfaction)
	var sum, maxVal int
	for i := 0; i < len(satisfaction); i++ {
		sum += satisfaction[i]
		maxVal += satisfaction[i] * (i + 1)
	}
	tmpVal := maxVal
	for i := 1; i < len(satisfaction); i++ {
		tmpVal = tmpVal - sum
		sum = sum - satisfaction[i-1]
		if tmpVal > maxVal {
			maxVal = tmpVal
		}
	}
	if maxVal < 0 {
		return 0
	}
	return maxVal
}
