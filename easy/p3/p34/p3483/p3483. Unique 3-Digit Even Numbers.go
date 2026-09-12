package p3483

func totalNumbers(digits []int) int {
	m := [10]int{}
	for _, digit := range digits {
		m[digit]++
	}
	res := 0
	for i := 100; i < 1000; i += 2 {
		a, b, c := getABC(i)
		m[a]--
		m[b]--
		m[c]--
		if m[a] >= 0 && m[b] >= 0 && m[c] >= 0 {
			res++
		}
		m[a]++
		m[b]++
		m[c]++
	}
	return res
}

func getABC(n int) (a, b, c int) {
	c = n % 10
	n /= 10
	b = n % 10
	n /= 10
	a = n % 10
	return a, b, c
}
