package p4

func maxCollectedFruits(fruits [][]int) int {
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

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
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
