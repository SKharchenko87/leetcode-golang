package p1356

import (
	"math/bits"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if bits.OnesCount(uint(arr[i])) == bits.OnesCount(uint(arr[j])) {
			return arr[i] < arr[j]
		}
		return bits.OnesCount(uint(arr[i])) < bits.OnesCount(uint(arr[j]))
	})
	return arr
}
