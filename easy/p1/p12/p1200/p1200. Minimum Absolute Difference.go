package p1200

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	minDiff := math.MaxInt
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		diff := abs(arr[i+1] - arr[i])
		if minDiff > diff {
			minDiff = diff
			count = 1
		} else if minDiff == diff {
			count++
		}
	}
	data := make([]int, count*2)
	result := make([][]int, count)
	for i := 0; i < count; i++ {
		result[i] = data[i*2 : i*2+2]
	}
	index := 0
	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i+1]-arr[i]) == minDiff {
			result[index][0] = arr[i]
			result[index][1] = arr[i+1]
			index++
		}
	}
	return result
}

func minimumAbsDifference1(arr []int) [][]int {
	l := len(arr)
	data := make([]int, l*2)
	result := make([][]int, l)
	for i := 0; i < l; i++ {
		result[i] = data[i*2 : i*2+2]
	}

	sort.Ints(arr)

	minDiff := math.MaxInt
	count := 0
	for i := 0; i < l-1; i++ {
		diff := abs(arr[i+1] - arr[i])
		if minDiff > diff {
			minDiff = diff
			count = 1
			result[0][0] = arr[i]
			result[0][1] = arr[i+1]
		} else if minDiff == diff {
			result[count][0] = arr[i]
			result[count][1] = arr[i+1]
			count++
		}
	}

	return result[:count]
}

func minimumAbsDifference0(arr []int) [][]int {
	sort.Ints(arr)
	minDiff := math.MaxInt
	for i := 0; i < len(arr)-1; i++ {
		minDiff = min(minDiff, abs(arr[i+1]-arr[i]))
	}
	var result [][]int
	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i+1]-arr[i]) == minDiff {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}
	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
