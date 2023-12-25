package main

import "fmt"

func checkSumIncrement(x int, y int, img [][]int, m *int, n *int, sum *int, cnt *int) {
	if (*m > y && y >= 0) && (*n > x && x >= 0) {
		*sum = *sum + img[y][x]
		*cnt++
	}
}

func imageSmoother(img [][]int) [][]int {
	m := len(img)
	n := len(img[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			sum := img[i][j]
			cnt := 1
			checkSumIncrement(j-1, i-1, img, &m, &n, &sum, &cnt)
			checkSumIncrement(j, i-1, img, &m, &n, &sum, &cnt)
			checkSumIncrement(j+1, i-1, img, &m, &n, &sum, &cnt)
			checkSumIncrement(j-1, i, img, &m, &n, &sum, &cnt)
			checkSumIncrement(j+1, i, img, &m, &n, &sum, &cnt)
			checkSumIncrement(j-1, i+1, img, &m, &n, &sum, &cnt)
			checkSumIncrement(j, i+1, img, &m, &n, &sum, &cnt)
			checkSumIncrement(j+1, i+1, img, &m, &n, &sum, &cnt)
			res[i][j] = sum / cnt
		}
	}
	return res
}

func main() {
	//img := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	//img := [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}
	img := [][]int{{2, 3, 4}, {5, 6, 7}, {8, 9, 10}, {11, 12, 13}, {14, 15, 16}}
	fmt.Println(imageSmoother(img))
}
