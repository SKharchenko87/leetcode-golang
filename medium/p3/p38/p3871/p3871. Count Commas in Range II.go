package p3871

import (
	"slices"
)

var pow10 = []int64{0, 1, 10, 100, 1000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000,
	10_000_000_000, 100_000_000_000, 1_000_000_000_000, 10_000_000_000_000, 100_000_000_000_000, 1_000_000_000_000_000, 10_000_000_000_000_000}
var multiplicator = []int64{0, 0, 0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5}
var cnt = make([]int64, 18)

func init() {
	for i := 5; i < 18; i++ {
		cnt[i] = cnt[i-1] + multiplicator[i]*(pow10[i]-pow10[i-1])
	}
}

func countCommas(n int64) int64 {
	if n < 1000 {
		return 0
	}
	index, ok := slices.BinarySearch(pow10, n)
	if ok {
		return cnt[index] + multiplicator[index+1]
	}
	return cnt[index-1] + (n-pow10[index-1]+1)*multiplicator[index]
}
