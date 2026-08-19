package p1386

import "sort"

const (
	maskLeft  = 0b1_0000_1111_1
	maskRight = 0b1_1111_0000_1
	maskMid   = 0b111_0000_111
	mask2     = 0b1_0000_0000_1
)

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	res := 2 * n
	m := make(map[int]int, 100)
	for i := 0; i < len(reservedSeats); i++ {
		m[reservedSeats[i][0]] = m[reservedSeats[i][0]] | (1 << (10 - reservedSeats[i][1]))
	}
	for _, v := range m {
		if v|mask2 == mask2 {
		} else if v|maskLeft == maskLeft || v|maskRight == maskRight || v|maskMid == maskMid {
			res--
		} else {
			res -= 2
		}
	}

	return res
}

func maxNumberOfFamilies0(n int, reservedSeats [][]int) int {
	res := 2 * n
	sort.Slice(reservedSeats, func(i, j int) bool {
		if reservedSeats[i][0] == reservedSeats[j][0] {
			return reservedSeats[i][1] < reservedSeats[j][1]
		}
		return reservedSeats[i][0] < reservedSeats[j][0]
	})

	for i := 0; i < len(reservedSeats); {
		cur := 1 << (10 - reservedSeats[i][1])
		i++
		for ; i < len(reservedSeats) && reservedSeats[i-1][0] == reservedSeats[i][0]; i++ {
			cur |= 1 << (10 - reservedSeats[i][1])
		}
		if cur|mask2 == mask2 {
		} else if cur|maskLeft == maskLeft || cur|maskRight == maskRight || cur|maskMid == maskMid {
			res--
		} else {
			res -= 2
		}
		if i < len(reservedSeats) {
			cur = 1 << (10 - reservedSeats[i][1])
		}
	}
	return res
}
