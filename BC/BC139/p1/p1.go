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

func stableMountains(height []int, threshold int) []int {
	l := len(height)
	index := 0
	result := make([]int, l)
	for i := 1; i < l; i++ {
		if height[i-1] > threshold {
			result[index] = i
			index++
		}
	}
	return result[:index]
}
