package p0873

import "slices"

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	res := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			x0, x1 := arr[i], arr[j]

			tmpLen := 0
			for _, found := slices.BinarySearch(arr, x0+x1); found; {
				tmpLen++
				x0, x1 = x1, x0+x1
				_, found = slices.BinarySearch(arr, x0+x1)
			}
			if res < tmpLen {
				res = tmpLen
			}
		}
	}
	if res == 0 {
		return 0
	}
	return res + 2
}
