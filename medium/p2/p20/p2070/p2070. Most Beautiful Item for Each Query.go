package p2070

import (
	"slices"
	"sort"
)

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		if items[i][0] == items[j][0] {
			return items[i][1] < items[j][1]
		}
		return items[i][0] < items[j][0]
	})

	maxVal := items[0][1]
	for i := 1; i < len(items); i++ {
		maxVal = max(maxVal, items[i][1])
		items[i][1] = maxVal
	}

	for i := 0; i < len(queries); i++ {
		index, _ := slices.BinarySearchFunc(items, []int{queries[i] + 1, 0}, func(a, b []int) int {
			return a[0] - b[0]
		})
		index--
		if index == -1 {
			queries[i] = 0
		} else {
			queries[i] = items[index][1]
		}
	}
	return queries
}

func maximumBeauty0(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		if items[i][0] == items[j][0] {
			return items[i][1] < items[j][1]
		}
		return items[i][0] < items[j][0]
	})

	newItem := make([][]int, 0, len(items))
	maxVal := items[0][1]
	for i := 1; i < len(items); i++ {
		if items[i][0] != items[i-1][0] {
			if maxVal < items[i-1][1] {
				maxVal = items[i-1][1]
			}
			newItem = append(newItem, []int{items[i-1][0], maxVal})
		}
	}
	newItem = append(newItem, []int{items[len(items)-1][0], max(maxVal, items[len(items)-1][1])})
	items = newItem

	for i := 0; i < len(queries); i++ {
		index, _ := slices.BinarySearchFunc(items, []int{queries[i] + 1, 0}, func(a, b []int) int {
			return a[0] - b[0]
		})
		index--
		if index == -1 {
			queries[i] = 0
		} else {
			queries[i] = items[index][1]
		}
	}
	return queries
}
