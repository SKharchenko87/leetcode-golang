package p2048

import (
	"slices"
)

const LIMIT = 1224444 + 1

var allBeautifulNumbers []int = make([]int, 0, 110)

func init() {
	for i := 1; i < LIMIT; i++ {
		if isBalanced(i) {
			allBeautifulNumbers = append(allBeautifulNumbers, i)
		}
	}
}

func nextBeautifulNumber0(n int) int {
	index, _ := slices.BinarySearch(allBeautifulNumbers, n+1)
	return allBeautifulNumbers[index]
}

func isBalanced(n int) bool {
	digits := [10]int{}
	for n > 0 {
		digits[n%10]++
		n /= 10
	}
	for digit, cnt := range digits {
		if cnt > 0 && digit != cnt {
			return false
		}
	}
	return true
}
