package p0857

import (
	"math"
	"sort"
)

type my struct {
	index       int
	quality     int
	coefficient float64
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {

	n := len(quality)
	tmp := make([]my, n)
	for i := 0; i < n; i++ {
		tmp[i] = my{i, quality[i], float64(wage[i]) / float64(quality[i])}
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].quality < tmp[j].quality
	})
	for j := 0; j < n; j++ {
		tmp[j].index = j
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].coefficient > tmp[j].coefficient
	})

	sort.Ints(quality)
	var res = math.MaxFloat64
	for i := 0; i < n-k+1; i++ {
		maxCoefficient := tmp[i].coefficient
		quality[tmp[i].index] = -1
		var tt float64
		j, c := 0, 0
		for c != k-1 {
			if quality[j] != -1 {
				tt += maxCoefficient * float64(quality[j])
				c++
			}
			j++
		}
		tt += maxCoefficient * float64(tmp[i].quality)
		if res > tt {
			res = tt
		}
	}
	//for i := 0; i < n; i++ {
	//	println(tmp2[i])
	//}
	return math.Round(res*100_000) / 100_000
}
