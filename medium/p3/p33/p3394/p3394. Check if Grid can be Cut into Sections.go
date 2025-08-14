package p3394

import "sort"

func check(rectangles *[][]int, start, end int) bool {
	sort.Slice(*rectangles, func(i, j int) bool {
		if (*rectangles)[i][start] == (*rectangles)[j][start] {
			return (*rectangles)[i][end] < (*rectangles)[j][end]
		}
		return (*rectangles)[i][start] < (*rectangles)[j][start]
	})
	m := (*rectangles)[0][end]
	res := 0
	for i := 1; i < len((*rectangles)); i++ {
		if (*rectangles)[i][start] >= m {
			res++
			if res == 2 {
				return true
			}
		}
		if (*rectangles)[i][end] > m {
			m = (*rectangles)[i][end]
		}
	}
	return false
}

func checkValidCuts(n int, rectangles [][]int) bool {
	return check(&rectangles, 0, 2) || check(&rectangles, 1, 3)
}

func checkValidCuts0(n int, rectangles [][]int) bool {
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
