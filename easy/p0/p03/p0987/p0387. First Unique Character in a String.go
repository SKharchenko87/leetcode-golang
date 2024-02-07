package p0987

import (
	"math"
	"sort"
)

// ToDo benchMark
func firstUniqChar(s string) int {
	arr := make([]int, 26)
	for i, c := range s {
		if arr[c-'a'] == 0 {
			arr[c-'a'] = i + 1
		} else {
			arr[c-'a'] = math.MaxInt
		}
	}
	sort.Ints(arr)
	for i := 0; i < 26; i++ {
		if arr[i] != 0 && arr[i] != math.MaxInt {
			return arr[i] - 1
		}
	}
	return -1
}
