package p2894

func differenceOfSums(n int, m int) int {
	sumN := n * (n + 1) / 2
	x := n / m
	sumX := x * (x + 1) / 2
	sumM := m * sumX
	return sumN - 2*sumM
}
