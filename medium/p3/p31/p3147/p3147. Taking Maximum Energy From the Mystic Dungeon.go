package p3147

import "math"

func maximumEnergy(energy []int, k int) int {
	l := len(energy)
	res := -1001
	for j := 0; j < k; j++ {
		tmp := 0
		for i := l - j - 1; i >= 0; i -= k {
			tmp += energy[i]
			res = max(res, tmp)
		}
	}
	return res
}

func maximumEnergy1(energy []int, k int) int {
	tmp := make([]int, k)
	res := -1001
	for i := len(energy) - 1; i >= 0; i-- {
		j := i % k
		tmp[j] += energy[i]
		res = max(res, tmp[j])
	}
	return res
}

func maximumEnergy0(energy []int, k int) int {
	tmp := make([]int, k)
	j := k - 1
	res := math.MinInt
	for i := len(energy) - 1; i >= 0; i-- {
		j += k
		j %= k
		tmp[j] += energy[i]
		res = max(res, tmp[j])
		j--
	}
	return res
}
