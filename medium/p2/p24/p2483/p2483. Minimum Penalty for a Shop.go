package p2483

import "strings"

func bestClosingTime(customers string) int {
	y := strings.Count(customers, "Y")
	hour := 0
	minimumPenalty := y
	n := 0
	for i := 0; i < len(customers); i++ {
		if customers[i] == 'Y' {
			y--
		} else {
			n++
		}
		if n+y < minimumPenalty {
			minimumPenalty = n + y
			hour = i + 1
		}
	}
	return hour
}
