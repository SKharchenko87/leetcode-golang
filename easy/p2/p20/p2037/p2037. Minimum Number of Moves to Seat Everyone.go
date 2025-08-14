package p2037

import "sort"

func minMovesToSeat(seats []int, students []int) (res int) {
	sort.Ints(seats)
	sort.Ints(students)
	for i, seat := range seats {
		res += abs(seat - students[i])
	}
	return res
}

func minMovesToSeat1(seats []int, students []int) (res int) {
	cntArr := [101]int{}
	n := len(students)
	for i := 0; i < n; i++ {
		cntArr[seats[i]]++
		cntArr[students[i]]--
	}
	tmp := 0
	for _, cnt := range cntArr {
		res += abs(tmp)
		tmp += cnt
	}
	return res
}

func minMovesToSeat0(seats []int, students []int) (res int) {
	sort.Ints(seats)
	sort.Ints(students)
	for i, seat := range seats {
		res += abs(seat - students[i])
	}
	return res
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}
