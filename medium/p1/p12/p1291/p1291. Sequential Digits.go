package p1291

func sequentialDigits(low int, high int) []int {
	m := [9][3]int{
		[3]int{1, 1, 9},
		[3]int{12, 11, 89},
		[3]int{123, 111, 789},
		[3]int{1234, 1111, 6789},
		[3]int{12345, 11111, 56789},
		[3]int{123456, 111111, 456789},
		[3]int{1234567, 1111111, 3456789},
		[3]int{12345678, 11111111, 23456789},
		[3]int{123456789, 111111111, 123456789},
	}
	res := []int{}
	for _, c := range m {
		if low <= c[2] {
			for i := c[0]; i <= c[2]; i = i + c[1] {
				if i > high {
					break
				}
				if low <= i && i <= high {
					res = append(res, i)
				}
			}
		}
		if c[0] > high {
			return res
		}
	}
	return res
}
