package p2285

import "sort"

func maximumImportance(n int, roads [][]int) int64 {
	cities := make([]int, 50001)
	for i := range roads {
		cities[roads[i][0]]++
		cities[roads[i][1]]++
	}
	var res int64
	sort.Ints(cities)
	for i := 50000; i > 0 && cities[i] > 0; i-- {
		res += int64(n * cities[i])
		n--
	}
	return res
}
