package p3740

import "math"

func minimumDistance(nums []int) int {
	matrix := [101][3]int{}
	res := math.MaxInt
	for i, n := range nums {
		move(&matrix[n], i)
		zero := 0
		if n == nums[0] {
			zero = 1
		}
		res = min(getDistance(matrix[n], zero), res)
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}

func move(indexes *[3]int, i int) {
	indexes[0] = indexes[1]
	indexes[1] = indexes[2]
	indexes[2] = i
}

func getDistance(indexes [3]int, zero int) int {
	if indexes[zero] == 0 {
		return math.MaxInt
	}
	return abs(indexes[0]-indexes[1]) + abs(indexes[1]-indexes[2]) + abs(indexes[0]-indexes[2])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
