package p1780

import (
	"slices"
)

func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}

var threes = []int{1, 3, 9, 27, 81, 243, 729, 2187, 6561, 19683, 59049, 177147, 531441, 1594323, 4782969}
var sums = []int{1, 4, 13, 40, 121, 364, 1093, 3280, 9841, 29524, 88573, 265720, 797161, 2391484, 7174453}

func checkPowersOfThree0(n int) bool {
	index, exists := slices.BinarySearch(threes, n)
	if exists {
		return true
	}
	if n > sums[index-1] {
		return false
	}
	return checkPowersOfThree(n - threes[index-1])
}
