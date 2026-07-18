package p1979

import "slices"

func findGCD(nums []int) int {
	mn, mx := slices.Min(nums), slices.Max(nums)
	return gcd(mx, mn)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
