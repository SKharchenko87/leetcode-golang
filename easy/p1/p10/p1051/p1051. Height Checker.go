package p1051

import "sort"

//ToDo bench mark

// Сортировка подсчётом (англ. Counting sort) с оптимизацией
func heightChecker(heights []int) int {
	cntArr := [100]int{}
	maxHeight := 0
	minHeight := 100
	for _, height := range heights {
		cntArr[height-1]++
		if height > maxHeight {
			maxHeight = height
		}
		if height-1 < minHeight {
			minHeight = height - 1
		}
	}
	res := len(heights)
	index := 0
	for i := minHeight; i < maxHeight; i++ {
		for j := 0; j < cntArr[i]; j++ {
			if heights[index] == i+1 {
				res--
			}
			index++
		}
	}
	return res
}

// Сортировка подсчётом (англ. Counting sort)
func heightChecker1(heights []int) int {
	cntArr := make([]int, 100)
	for _, height := range heights {
		cntArr[height-1]++
	}
	res := len(heights)
	index := 0
	for i, c := range cntArr {
		for j := 0; j < c; j++ {
			if heights[index] == i+1 {
				res--
			}
			index++
		}
	}
	return res
}

// В лоб (pdqsort)
func heightChecker0(heights []int) int {
	newHeights := make([]int, len(heights))
	copy(newHeights, heights)
	sort.Ints(newHeights)
	cnt := 0
	for i, height := range newHeights {
		if heights[i] != height {
			cnt++
		}
	}
	return cnt
}
