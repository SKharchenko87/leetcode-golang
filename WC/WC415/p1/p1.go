package p1

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

func getSneakyNumbers(nums []int) []int {
	a := [101]int{}
	res := make([]int, 0, 2)
	for _, num := range nums {
		a[num]++
		if a[num] > 1 {
			res = append(res, num)
		}
	}
	return res
}
