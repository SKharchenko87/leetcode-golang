package p1137

//ToDo banchmark
func tribonacciDP(n int, m *map[int]int) int {
	if v, ok := (*m)[n]; ok {
		return v
	} else {
		(*m)[n] = tribonacciDP(n-3, m) + tribonacciDP(n-2, m) + tribonacciDP(n-1, m)
	}
	return (*m)[n]
}

//func tribonacci(n int) int {
//	m := map[int]int{0: 0, 1: 1, 2: 1, 3: 2, 4: 4, 5: 7, 6: 13, 7: 24, 8: 44, 9: 81, 10: 149, 11: 274, 12: 504, 13: 927, 14: 1705, 15: 3136, 16: 5768, 17: 10609, 18: 19513, 19: 35890, 20: 66012, 21: 121415, 22: 223317, 23: 410744, 24: 755476, 25: 1389537, 26: 2555757, 27: 4700770, 28: 8646064, 29: 15902591, 30: 29249425, 31: 53798080, 32: 98950096, 33: 181997601, 34: 334745777, 35: 615693474, 36: 1132436852, 37: 2082876103}
//	if v, ok := m[n]; ok {
//		return v
//	}
//	return tribonacciDP(n, &m)
//}

func tribonacci(n int) int {
	m := map[byte]int32{0: 0, 1: 1, 2: 1, 3: 2, 4: 4, 5: 7, 6: 13, 7: 24, 8: 44, 9: 81, 10: 149, 11: 274, 12: 504, 13: 927, 14: 1705, 15: 3136, 16: 5768, 17: 10609, 18: 19513, 19: 35890, 20: 66012, 21: 121415, 22: 223317, 23: 410744, 24: 755476, 25: 1389537, 26: 2555757, 27: 4700770, 28: 8646064, 29: 15902591, 30: 29249425, 31: 53798080, 32: 98950096, 33: 181997601, 34: 334745777, 35: 615693474, 36: 1132436852, 37: 2082876103}
	return int(m[byte(n)])
}
