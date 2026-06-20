package p1840

import "sort"

func maxBuilding(n int, restrictions [][]int) int {
	// Add restriction (1, 0)
	r := append(restrictions, []int{1, 0})

	// Sort by position
	sort.Slice(r, func(i, j int) bool {
		return r[i][0] < r[j][0]
	})
	// Add restriction (n, n-1)
	if r[len(r)-1][0] != n {
		r = append(r, []int{n, n - 1})
	}
	m := len(r)

	// Pass restrictions from left to right
	for i := 1; i < m; i++ {
		dist := r[i][0] - r[i-1][0]
		r[i][1] = min(r[i][1], r[i-1][1]+dist)
	}

	// Pass restrictions from right to left
	for i := m - 2; i >= 0; i-- {
		dist := r[i+1][0] - r[i][0]
		r[i][1] = min(r[i][1], r[i+1][1]+dist)
	}

	ans := 0
	for i := 0; i < m-1; i++ {
		ans = max(ans, maxHeightOfSection(r, i))
	}

	return ans
}

// Calculate the maximum height of the buildings between r[i][0] and r[i][1]
func maxHeightOfSection(restrictions [][]int, i int) int {
	dist := restrictions[i+1][0] - restrictions[i][0]
	return (dist + restrictions[i][1] + restrictions[i+1][1]) / 2
}
