package p1671

import (
	"math"
	"sort"
)

func getLongestIncreasingSubsequenceLength(v []int) []int {
	lisLen := make([]int, len(v))
	for i := range lisLen {
		lisLen[i] = 1
	}
	lis := []int{v[0]}

	for i := 1; i < len(v); i++ {
		index := sort.Search(len(lis), func(j int) bool {
			return lis[j] >= v[i]
		})

		if index == len(lis) {
			lis = append(lis, v[i])
		} else {
			lis[index] = v[i]
		}

		lisLen[i] = len(lis)
	}

	return lisLen
}

func minimumMountainRemovals(nums []int) int {
	N := len(nums)
	lisLength := getLongestIncreasingSubsequenceLength(nums)

	// Reverse the nums slice
	for i, j := 0, N-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	ldsLength := getLongestIncreasingSubsequenceLength(nums)

	// Reverse the length array to correctly depict the length of longest
	// decreasing subsequence for each index.
	for i, j := 0, N-1; i < j; i, j = i+1, j-1 {
		ldsLength[i], ldsLength[j] = ldsLength[j], ldsLength[i]
	}

	minRemovals := math.MaxInt32
	for i := 0; i < N; i++ {
		if lisLength[i] > 1 && ldsLength[i] > 1 {
			minRemovals = min(minRemovals, N-lisLength[i]-ldsLength[i]+1)
		}
	}

	return minRemovals
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
