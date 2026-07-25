package p3658

/*
sumOdd, sumEven = n*n
sumEven, sumOdd = n*(n+1)
=> gcd(n*n, n*(n+1)) => n
*/
func gcdOfOddEvenSums(n int) int {
	return n
}

func gcdOfOddEvenSums0(n int) int {
	sum := (n*2 + 1) / 2 * (n*2 + 1)
	sumEven := n * (n + 1)
	sumOdd := sum - sumEven
	return gcd(sumEven, sumOdd)
}

func gcd(a, b int) int {
	a, b = min(a, b), max(a, b)
	for a != 0 {
		b, a = a, b%a
	}
	return b
}
