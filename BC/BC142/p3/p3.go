package p3

type scorecity struct {
	cityIndex int
	score     int
}

func maxScore(n int, k int, stayScore [][]int, travelScore [][]int) int {
	//m := []scorecity{}
	m := map[int]int{}
	for i, s := range stayScore[0] {
		m[i] = s
	}
	for i := 0; i < len(travelScore); i++ {
		for j := 0; j < len(travelScore[i]); j++ {
			if m[j] < travelScore[i][j] {
				m[j] = travelScore[i][j]
			}
		}
	}

	for i := 1; i < k; i++ {
		newM := map[int]int{}
		for city, v := range m {
			if newM[city] < v+stayScore[i][city] {
				newM[city] = v + stayScore[i][city]
			}
			for newCity, score := range travelScore[city] {
				if newM[newCity] < v+score {
					newM[newCity] = v + score
				}
			}
		}
		m = newM
	}
	res := -1
	for _, v := range m {
		if res < v {
			res = v
		}
	}
	return res
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
