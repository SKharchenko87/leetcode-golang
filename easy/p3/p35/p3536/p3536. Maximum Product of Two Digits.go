package p3536

func maxProduct(n int) int {
	var a, b, x int
	for n > 0 {
		x, n = n%10, n/10
		if x >= b {
			a, b = b, x
		} else if x > a {
			a = x
		}
	}
	return a * b
}

func maxProduct0(n int) int {
	a, n := n%10, n/10
	b, n := n%10, n/10
	a, b = min(a, b), max(a, b)
	var x int
	for n > 0 {
		x, n = n%10, n/10
		if x >= b {
			a, b = b, x
		} else if x > a {
			a = x
		}
	}
	return a * b
}
