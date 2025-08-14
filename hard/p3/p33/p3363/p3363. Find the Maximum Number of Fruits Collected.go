package p3363

func maxCollectedFruits(fruits [][]int) int {
	n := len(fruits)
	child0Basket := 0
	for i := 0; i < n; i++ {
		child0Basket += fruits[i][i]
	}

	for j := 0; j < n/2; j++ {
		for i := n - 1; i >= n-1-j; i-- {
			fruits[i][j] = maxValPrev(fruits, i, j, 1, DOWN)
		}
	}
	for j := n / 2; j < n-1; j++ {
		for i := j + 1; i < n; i++ {
			fruits[i][j] = maxValPrev(fruits, i, j, 1, DOWN)
		}
	}

	for i := 0; i < n/2; i++ {
		for j := n - 1 - i; j < n; j++ {
			fruits[i][j] = maxValPrev(fruits, i, j, 0, RIGHT)
		}
	}
	for i := n / 2; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			fruits[i][j] = maxValPrev(fruits, i, j, 0, RIGHT)
		}
	}

	child1Basket := fruits[n-2][n-1]
	child2Basket := fruits[n-1][n-2]

	return child0Basket + child1Basket + child2Basket
}

func maxCollectedFruits0(fruits [][]int) int {
	n := len(fruits)
	child0Basket := 0
	for i := 0; i < n; i++ {
		child0Basket += fruits[i][i]
		fruits[i][i] = 0
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fruits[i][j] = 0
		}
	}

	for j := 0; j < n/2; j++ {
		for i := n - 1; i >= n-1-j; i-- {
			fruits[i][j] = maxValPrev(fruits, i, j, 1, DOWN)
		}
	}
	for j := n / 2; j < n-1; j++ {
		for i := j + 1; i < n; i++ {
			fruits[i][j] = maxValPrev(fruits, i, j, 1, DOWN)
		}
	}

	for i := 0; i < n/2; i++ {
		for j := n - 1 - i; j < n; j++ {
			fruits[i][j] = maxValPrev(fruits, i, j, 0, RIGHT)
		}
	}
	for i := n / 2; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			fruits[i][j] = maxValPrev(fruits, i, j, 0, RIGHT)
		}
	}

	child1Basket := fruits[n-2][n-1]
	child2Basket := fruits[n-1][n-2]

	return child0Basket + child1Basket + child2Basket
}

var directions = [][]int{{-1, -1}, {-1, 0}, {-1, 1}}

func maxValPrev(fruits [][]int, i, j int, mod int, positionTriangle int) int {
	n := len(fruits)
	x := 0
	for _, dir := range directions {
		currI, currJ := i+dir[mod%2], j+dir[(mod+1)%2] //10  DOWN
		if getTriangle(currI, currJ, n) == positionTriangle {
			x = max(x, fruits[currI][currJ])
		}
	}
	return fruits[i][j] + x
}

func isMainDiagonal(i, j, n int) int {
	return j - i
}

func isSubDiagonal(i, j, n int) int {
	return n - 1 - i - j
}

const (
	UNCNOWN = -1
	UP      = iota
	RIGHT
	DOWN
	LEFT
)

func getTriangle(i, j, n int) int {
	if !(0 <= i && i < n && 0 <= j && j < n) {
		return UNCNOWN
	}
	if isMainDiagonal(i, j, n) >= 0 && isSubDiagonal(i, j, n) > 0 {
		return UP
	} else if isMainDiagonal(i, j, n) > 0 && isSubDiagonal(i, j, n) <= 0 {
		return RIGHT
	} else if isMainDiagonal(i, j, n) <= 0 && isSubDiagonal(i, j, n) > 0 {
		return LEFT
	} else if isMainDiagonal(i, j, n) < 0 && isSubDiagonal(i, j, n) <= 0 {
		return DOWN
	}
	return UNCNOWN
}

func maxCollectedFruits2(fruits [][]int) int {
	n := len(fruits)
	child0Basket := 0
	for i := 0; i < n; i++ {
		child0Basket += fruits[i][i]
		// fruits[i][i] = 0
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fruits[i][j] = 0
		}
	}

	for j := 0; j < n/2; j++ {
		for i := n - 1; i >= n-1-j; i-- {
			fruits[i][j] = maxValPrev1(fruits, i, j)
		}
	}
	for j := n / 2; j < n-1; j++ {
		for i := j + 1; i < n; i++ {
			fruits[i][j] = maxValPrev1(fruits, i, j)
		}
	}

	for i := 0; i < n/2; i++ {
		for j := n - 1 - i; j < n; j++ {
			fruits[i][j] = maxValPrev2(fruits, i, j)
		}
	}
	for i := n / 2; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			fruits[i][j] = maxValPrev2(fruits, i, j)
		}
	}

	//fmt.Println(fruits)

	child1Basket := fruits[n-2][n-1]
	child2Basket := fruits[n-1][n-2]

	return child0Basket + child1Basket + child2Basket
}

func maxValPrev1(fruits [][]int, i, j int) int {
	n := len(fruits)
	var x0, x1, x2 int
	if j > 0 {
		currI, currJ := i-1, j-1
		if i-1 >= 0 {
			x0 = fruits[currI][currJ]
		}
		currI, currJ = i, j-1
		x1 = fruits[currI][currJ]

		currI, currJ = i+1, j-1
		if i+1 < n {
			x2 = fruits[currI][currJ]
		}
	}
	return fruits[i][j] + max(x0, x1, x2)
}

func maxValPrev2(fruits [][]int, i, j int) int {
	n := len(fruits)
	var x0, x1, x2 int
	if i > 0 {
		currI, currJ := i-1, j-1
		if j-1 >= 0 && currJ-currI > 0 && n-1-currI-currJ <= 0 {
			x0 = fruits[currI][currJ]
		}

		currI, currJ = i-1, j
		if currJ-currI > 0 && n-1-currI-currJ <= 0 {
			x1 = fruits[currI][currJ]
		}

		currI, currJ = i-1, j+1
		if j+1 < n && n-1-currI-currJ <= 0 {
			x2 = fruits[currI][currJ]
		}
	}
	return fruits[i][j] + max(x0, x1, x2)
}

func maxValPrev1_1(fruits [][]int, i, j int) int {
	n := len(fruits)
	var x int
	x = fruits[i-1][j-1]
	x = max(fruits[i][j-1], x)
	if i+1 < n {
		x = max(fruits[i+1][j-1], x)
	}
	return fruits[i][j] + x
}

func maxValPrev1_2(fruits [][]int, i, j int) int {
	n := len(fruits)
	var x int
	x = fruits[i-1][j-1]
	x = max(fruits[i-1][j], x)
	if j+1 < n {
		x = max(fruits[i-1][j+1], x)
	}
	return fruits[i][j] + x
}

func maxCollectedFruits1(fruits [][]int) int {
	n := len(fruits)
	child0Basket := 0
	for i := 0; i < n; i++ {
		child0Basket += fruits[i][i]
	}

	for i := 0; i < n-1; i++ {
		fruits[i][n-1-i-1] = 0
	}
	for i := 0; i < n-2; i++ {
		fruits[i][n-2-i-1] = 0
	}

	for j := 1; j < n/2; j++ {
		for i := n - 1; i >= n-1-j; i-- {
			fruits[i][j] = maxValPrev1_1(fruits, i, j)
		}
	}
	for j := n / 2; j < n-1; j++ {
		for i := j + 1; i < n; i++ {
			fruits[i][j] = maxValPrev1_1(fruits, i, j)
		}
	}

	for i := 1; i < n/2; i++ {
		for j := n - 1 - i; j < n; j++ {
			fruits[i][j] = maxValPrev1_2(fruits, i, j)
		}
	}
	for i := n / 2; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			fruits[i][j] = maxValPrev1_2(fruits, i, j)
		}
	}

	//fmt.Println(fruits)
	/*
		16,  3, 11, 14, 14
		 3,  0, 10, 13, 14
		 7, 18,  8,  7, 18
		 7,  8,  5,  7,  5
		 0, 14,  8,  1,  0
	*/
	/*
		16,  3, 11,  0, 14
		 3,  0,  0, 13, 14
		 7,  0,  8,  7, 18
		 0,  8,  5,  7,  5
		 0, 14,  8,  1,  0
	*/
	//[[16 3 11 0 14] [3 0 0 27 28] [7 0 8 35 46] [0 15 20 7 51] [0 14 23 24 0]]
	//[[ 0 3 11 0 14] [3 0 0 27 28] [7 0 0 35 46] [0 15 20 0 51] [0 14 23 24 0]]
	//[[ 0 0  0 0 14] [0 0 0 27 28] [0 0 8 35 46] [0  8 19 7 51] [0 14 22 23 0]]
	child1Basket := fruits[n-2][n-1]
	child2Basket := fruits[n-1][n-2]

	return child0Basket + child1Basket + child2Basket
}

func maxCollectedFruits3(fruits [][]int) int {
	n := len(fruits)
	green, red, blue := 0, 0, 0
	for i := 0; i < n; i++ {
		green += fruits[i][i]
		fruits[i][i] = 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < n-i-1 {
				fruits[i][j] = 0
			}
		}
	}
	//red
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			var l, m, r int
			if j-1 >= 0 {
				l = fruits[i-1][j-1]
			}
			m = fruits[i-1][j]
			if j+1 < n {
				r = fruits[i-1][j+1]
			}
			fruits[i][j] += max(max(l, m), r)
		}
	}
	//blue
	for j := 1; j < n-1; j++ {
		for i := n - 1; i >= n-1-i; i-- {
			var u, m, d int
			if i-1 >= 0 {
				u = fruits[i-1][j-1]
			}
			m = fruits[i][j-1]
			if i+1 < n {
				d = fruits[i+1][j-1]
			}
			fruits[i][j] += max(max(u, m), d)
		}
	}

	//fmt.Println(fruits)
	red = fruits[n-2][n-1]
	blue = fruits[n-1][n-2]
	return green + red + blue
}

func maxCollectedFruits4(fruits [][]int) int {
	n := len(fruits)
	green, red, blue := 0, 0, 0
	for i := 0; i < n; i++ {
		green += fruits[i][i]
	}

	//for i := 0; i < n-1; i++ {
	//	fruits[i][n-1-i-1] = 0
	//}
	//for i := 0; i < n-2; i++ {
	//	fruits[i][n-2-i-1] = 0
	//}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < n-i-1 {
				fruits[i][j] = 0
			}
		}
	}

	//red
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			var l, m, r int
			if j-1 >= 0 {
				l = fruits[i-1][j-1]
			}
			m = fruits[i-1][j]
			if j+1 < n {
				r = fruits[i-1][j+1]
			}
			fruits[i][j] += max(max(l, m), r)
		}
	}
	//blue
	for j := 1; j < n-1; j++ {
		for i := n - 1; i >= n-1-i; i-- {
			var u, m, d int
			if i-1 >= 0 {
				u = fruits[i-1][j-1]
			}
			m = fruits[i][j-1]
			if i+1 < n {
				d = fruits[i+1][j-1]
			}
			fruits[i][j] += max(max(u, m), d)
		}
	}

	//fmt.Println(fruits)
	red = fruits[n-2][n-1]
	blue = fruits[n-1][n-2]
	return green + red + blue
}
