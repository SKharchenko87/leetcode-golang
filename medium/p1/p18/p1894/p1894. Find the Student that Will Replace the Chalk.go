package p1894

import "slices"

func chalkReplacer(chalk []int, k int) int {
	l := len(chalk)
	for i := 1; i < l; i++ {
		chalk[i] += chalk[i-1]
	}
	k %= chalk[l-1]
	res, _ := slices.BinarySearch(chalk, k+1)
	return res
}

func chalkReplacer1(chalk []int, k int) int {
	l := len(chalk)
	prefSum := make([]int, l)
	prefSum[0] = chalk[0]
	for i := 1; i < l; i++ {
		prefSum[i] = prefSum[i-1] + chalk[i]
	}
	k %= prefSum[l-1]
	res, b := slices.BinarySearch(prefSum, k)
	if b {
		return res + 1
	}
	return res
}

func chalkReplacer0(chalk []int, k int) int {
	sum := 0
	for i := 0; i < len(chalk); i++ {
		sum += chalk[i]
	}
	k %= sum
	for i := 0; i < len(chalk); i++ {
		if k-chalk[i] < 0 {
			return i
		}
		k -= chalk[i]
	}
	return 0
}
