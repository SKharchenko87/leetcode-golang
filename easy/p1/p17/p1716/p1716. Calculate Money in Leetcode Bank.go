package p1716

/*
28 = 28+0*7		1 -  7	 	weekNumber:0	Monday Money: 1
35 = 28+1*7		2 -  8	 	weekNumber:1	Monday Money: 2
42 = 28+2*7		3 -  9	 	weekNumber:2	Monday Money: 3
49 = 28+3*7		4 - 10	 	weekNumber:3	Monday Money: 4
56 = 28+4*7		5 - 11	 	weekNumber:4	Monday Money: 5

res := 28*(weeks) + weeks*(weeks-1)/2*7
res += (weeks+1)*daysOfLastWeek + daysOfLastWeek*(daysOfLastWeek-1)/2
*/
func totalMoney(n int) int {
	n--
	weeks := n / 7
	daysOfLastWeek := n%7 + 1

	res := (7*weeks*(weeks+7) + daysOfLastWeek*(2*weeks+daysOfLastWeek+1)) / 2

	return res
}
