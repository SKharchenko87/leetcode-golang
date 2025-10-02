package p2523

import "math"

func closestPrimes(left int, right int) []int {
	result := make([]int, 2)
	result[0] = -1
	result[1] = -1
	if left == right {
		return result
	}
	if left > 3 {
		lastPrime := -1_000_000
		maxLen := 1_000_000
		if left%2 == 0 {
			left++
		}
		for i := left; i <= right; i += 2 {
			if isPrime(i) {
				if maxLen > i-lastPrime {
					maxLen = i - lastPrime
					result[0] = lastPrime
					result[1] = i
					if maxLen == 2 {
						return result
					}
				}
				lastPrime = i
			}
		}
	} else if left == 3 {
		if right >= 5 {
			result[0] = 3
			result[1] = 5
		}
	} else {
		if left == 1 {
			left++
		}
		result[0] = left
		result[1] = left + 1
	}
	return result
}

func isPrime(n int) bool {
	if n <= 3 {
		return true
	} else if n%2 == 0 {
		return false
	} else if n%3 == 0 {
		return false
	}
	maxN := int(math.Sqrt(float64(n)))
	for i := 5; i <= maxN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPrime0(n int) bool {
	if n <= 3 {
		return true
	} else if n%2 == 0 {
		return false
	} else if n%3 == 0 {
		return false
	}
	for i := 5; i < n/4; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
