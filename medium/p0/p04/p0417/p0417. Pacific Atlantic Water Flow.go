package p0417

import "sort"

var DIRECTIONS = [4][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pacific := ocean(heights, 0, 0)
	atlantic := ocean(heights, n-1, m-1)
	sort.Slice(pacific, func(i, j int) bool {
		if pacific[i][0] == pacific[j][0] {
			return pacific[i][1] < pacific[j][1]
		}
		return pacific[i][0] < pacific[j][0]
	})
	sort.Slice(atlantic, func(i, j int) bool {
		if atlantic[i][0] == atlantic[j][0] {
			return atlantic[i][1] < atlantic[j][1]
		}
		return atlantic[i][0] < atlantic[j][0]
	})

	j := 0
	res := [][]int{}
	for i := 0; i < len(pacific); i++ {
		p := pacific[i]
		if j < len(atlantic) {
			for ; j < len(atlantic) && ((p[0] == atlantic[j][0] && p[1] > atlantic[j][1]) || p[0] > atlantic[j][0]); j++ {
			}
			if p[0] == atlantic[j][0] && p[1] == atlantic[j][1] {
				res = append(res, []int{p[0], p[1]})
			}
		}
	}
	return res

}

func ocean(heights [][]int, limitI, limitJ int) [][2]int {
	m, n := len(heights), len(heights[0])
	pacific := make([][2]int, 0, m*n)
	pacificVisited := make([][]bool, m)
	for j := 0; j < m; j++ {
		pacificVisited[j] = make([]bool, n)
	}
	pacificCandidate := make([][2]int, 0, m*n)
	for j := 0; j < m; j++ {
		if !pacificVisited[j][limitI] {
			pacificCandidate = append(pacificCandidate, [2]int{j, limitI})
			pacific = append(pacific, [2]int{j, limitI})
			pacificVisited[j][limitI] = true
		}
	}
	for i := 0; i < n; i++ {
		if !pacificVisited[limitJ][i] {
			pacificCandidate = append(pacificCandidate, [2]int{limitJ, i})
			pacific = append(pacific, [2]int{limitJ, i})
			pacificVisited[limitJ][i] = true
		}
	}

	for k := 0; k < len(pacificCandidate); k++ {
		for _, direction := range DIRECTIONS {
			j, i := pacificCandidate[k][0], pacificCandidate[k][1]
			jc, ic := pacificCandidate[k][0]+direction[0], pacificCandidate[k][1]+direction[1]
			if 0 <= jc && jc < m && 0 <= ic && ic < n && !pacificVisited[jc][ic] && heights[jc][ic] >= heights[j][i] {
				pacificCandidate = append(pacificCandidate, [2]int{jc, ic})
				pacific = append(pacific, [2]int{jc, ic})
				pacificVisited[jc][ic] = true
			}
		}
	}
	return pacific
}
