package p1317

func checkZero(n int) bool {
	for n > 0 {
		if n%10 == 0 {
			return false
		}
		n /= 10
	}
	return true
}

func getNoZeroIntegers(n int) []int {
	for a := 1; a < n; a++ {
		b := n - a
		if checkZero(a) && checkZero(b) {
			return []int{a, b}
		}
	}
	return []int{}
}
