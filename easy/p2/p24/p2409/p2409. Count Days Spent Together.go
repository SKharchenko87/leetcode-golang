package p2409

import "strconv"

var dayOfMonth = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var days = [12]int{}

func init() {
	days[0] = 0
	for i := 1; i < 12; i++ {
		days[i] = dayOfMonth[i-1] + days[i-1]
	}
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	arriveAliceMM, _ := strconv.Atoi(arriveAlice[:2])
	arriveAliceDD, _ := strconv.Atoi(arriveAlice[3:])
	leaveAliceMM, _ := strconv.Atoi(leaveAlice[:2])
	leaveAliceDD, _ := strconv.Atoi(leaveAlice[3:])

	arriveBobMM, _ := strconv.Atoi(arriveBob[:2])
	arriveBobDD, _ := strconv.Atoi(arriveBob[3:])
	leaveBobMM, _ := strconv.Atoi(leaveBob[:2])
	leaveBobDD, _ := strconv.Atoi(leaveBob[3:])

	a0, a1 := days[arriveAliceMM-1]+arriveAliceDD, days[leaveAliceMM-1]+leaveAliceDD
	b0, b1 := days[arriveBobMM-1]+arriveBobDD, days[leaveBobMM-1]+leaveBobDD
	if a0 > b0 {
		a0, a1, b0, b1 = b0, b1, a0, a1
	}
	if b0 > a1 {
		return 0
	}
	return min(a1, b1) - b0 + 1
}
