package p0476

import "math/bits"

func findComplement(num int) int {
	return (1<<bits.Len(uint(num)) - 1) ^ num
}

func findComplement1(num int) int {
	return (1<<bits.Len(uint(num)) - 1) &^ num
}

func findComplement0(num int) int {
	x := 0
	for i := 0; i < 32 && x < num; i++ {
		x = 1 << i
	}
	x--
	return x &^ num
}
