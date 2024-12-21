package p3

import "sort"

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func checkValidCuts(n int, rectangles [][]int) bool {
	sort.Slice(rectangles, func(i, j int) bool {
		if rectangles[i][0] == rectangles[j][0] {
			return rectangles[i][2] < rectangles[j][2]
		}
		return rectangles[i][0] < rectangles[j][0]
	})
	m := rectangles[0][2]
	res := 0
	for i := 1; i < len(rectangles); i++ {
		if rectangles[i][0] >= m {
			res++
			if res == 2 {
				return true
			}
		}
		if rectangles[i][2] > m {
			m = rectangles[i][2]
		}

	}

	sort.Slice(rectangles, func(i, j int) bool {
		if rectangles[i][1] == rectangles[j][1] {
			return rectangles[i][3] < rectangles[j][3]
		}
		return rectangles[i][1] < rectangles[j][1]
	})
	m = rectangles[0][3]
	res = 0
	for i := 1; i < len(rectangles); i++ {
		if rectangles[i][1] >= m {
			res++
			if res == 2 {
				return true
			}
		}
		if rectangles[i][3] > m {
			m = rectangles[i][3]
		}
	}

	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
