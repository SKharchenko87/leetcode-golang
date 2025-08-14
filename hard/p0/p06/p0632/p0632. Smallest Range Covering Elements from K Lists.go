package p0632

import (
	"math"
	"sort"
)

func smallestRange(nums [][]int) []int {
	merged := make([][2]int, 0)

	// Merge all lists with their list index
	for i := range nums {
		for _, num := range nums[i] {
			merged = append(merged, [2]int{num, i})
		}
	}

	// Sort the merged list
	sort.Slice(merged, func(i, j int) bool {
		return merged[i][0] < merged[j][0]
	})

	// Two pointers to track the smallest range
	freq := make(map[int]int)
	left, count := 0, 0
	rangeStart, rangeEnd := 0, math.MaxInt32

	for right := 0; right < len(merged); right++ {
		freq[merged[right][1]]++
		if freq[merged[right][1]] == 1 {
			count++
		}

		// When all lists are represented, try to shrink the window
		for count == len(nums) {
			curRange := merged[right][0] - merged[left][0]
			if curRange < rangeEnd-rangeStart {
				rangeStart = merged[left][0]
				rangeEnd = merged[right][0]
			}

			freq[merged[left][1]]--
			if freq[merged[left][1]] == 0 {
				count--
			}
			left++
		}
	}

	return []int{rangeStart, rangeEnd}
}

func smallestRange0(nums [][]int) []int {
	k := len(nums)
	allLen := 0
	for i := 0; i < k; i++ {
		allLen += len(nums[i])
	}
	all := make([][2]int, 0, allLen)
	for i := 0; i < k; i++ {
		for j := 0; j < len(nums[i]); j++ {
			all = append(all, [2]int{nums[i][j], i})
		}
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i][0] < all[j][0]
	})

	freq := make(map[int]int, k)
	count := 0
	left := 0
	rangeStart, rangeEnd := 0, allLen
	for right := 0; right < allLen; right++ {
		freq[all[right][1]]++
		if freq[all[right][1]] == 1 {
			count++
		}

		for count == k {
			curRange := all[right][0] - all[left][0]
			if curRange < rangeEnd-rangeStart {
				rangeStart = all[left][0]
				rangeEnd = all[right][0]
			}

			freq[all[left][1]]--
			if freq[all[left][1]] == 0 {
				count--
			}
			left++
		}
	}
	return []int{rangeStart, rangeEnd}
}
