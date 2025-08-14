package p2192

import (
	"sort"
)

func getAncestors(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	for _, edge := range edges {
		if graph[edge[0]] == nil {
			graph[edge[0]] = []int{}
		}
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}
	answer := make([][]int, n)
	var dfs func(u, v int)
	dfs = func(u, v int) {
		for _, i := range graph[v] {
			if answer[v] == nil {
				answer[v] = []int{}
			}
			if len(answer[i]) > 0 && answer[i][len(answer[i])-1] == u {
				continue
			}
			answer[i] = append(answer[i], u)
			dfs(u, i)
		}
	}
	for i := 0; i < n; i++ {
		dfs(i, i)
	}
	return answer
}

type node struct {
	parents map[int]int
}

func getAncestors3(n int, edges [][]int) [][]int {
	answer := make([][]int, n)
	for _, edge := range edges {
		if answer[edge[1]] == nil {
			answer[edge[1]] = []int{}
		}
		i := 0
		for ; i < len(answer[edge[1]]) && answer[edge[1]][i] < edge[0]; i++ {
		}
		answer[edge[1]] = *insert(&answer[edge[1]], i, edge[0])
	}
	visited := make([]bool, n)
	var rec func(x int) []int
	rec = func(x int) []int {
		if visited[x] {
			return answer[x]
		}
		if answer[x] == nil {
			answer[x] = []int{}
			visited[x] = true
			return answer[x]
		}
		for k := 0; k < len(answer[x]); k++ {
			p := answer[x][k]
			pm := rec(p)
			i := 0
			for j := 0; j < len(pm); j++ {
				for ; i < len(answer[x]) && answer[x][i] < pm[j]; i++ {
				}
				if i >= len(answer[x]) || answer[x][i] != pm[j] {
					answer[x] = *insert(&answer[x], i, pm[j])
				}
			}
		}
		visited[x] = true
		return answer[x]
	}
	for i := 0; i < n; i++ {
		answer[i] = rec(i)
	}
	return answer
}

func insert(arr *[]int, position int, value int) *[]int {
	var mass []int = *arr
	mass = append(mass[:position], append([]int{value}, mass[position:]...)...)
	return &mass
}

func getAncestors2(n int, edges [][]int) [][]int {
	nods := make([]node, n)
	for _, edge := range edges {
		if nods[edge[1]].parents == nil {
			nods[edge[1]].parents = make(map[int]int)
		}
		nods[edge[1]].parents[edge[0]]++
	}
	visited := make([]bool, n)
	var rec func(x int) map[int]int
	rec = func(x int) map[int]int {
		if nods[x].parents == nil {
			return map[int]int{}
		}
		if visited[x] {
			return nods[x].parents
		}
		tmp := map[int]int{}
		for p, _ := range nods[x].parents {
			pm := rec(p)
			for pk, _ := range pm {
				tmp[pk]++
			}
		}
		for pk, _ := range tmp {
			nods[x].parents[pk]++
		}
		visited[x] = true
		return nods[x].parents
	}
	answer := make([][]int, n)
	for i := 0; i < n; i++ {
		answer[i] = []int{}
		for k, _ := range rec(i) {
			answer[i] = append(answer[i], k)
		}
		sort.Ints(answer[i])
	}
	return answer
}

func getAncestors1(n int, edges [][]int) [][]int {
	answer := make([][]int, n)
	m := make(map[int]map[int]int, n)
	for _, edge := range edges {
		if _, ok := m[edge[1]]; !ok {
			m[edge[1]] = map[int]int{}
		}
		m[edge[1]][edge[0]]++

	}
	visited := map[int]bool{}
	var recursive func(x, y int)
	recursive = func(x, y int) {
		m[x][y]++
		if w, ok := m[y]; !ok {
			return
		} else {
			for y0, _ := range w {
				m[x][y0]++
				if !visited[y] {
					recursive(x, y0)
				}
			}
		}
		visited[x] = true
	}
	for k := 0; k < n; k++ {
		v := m[k]
		for y, _ := range v {
			recursive(k, y)
		}
		tmp := []int{}
		for w, _ := range m[k] {
			tmp = append(tmp, w)
		}
		sort.Ints(tmp)
		answer[k] = tmp
	}
	return answer
}

func getAncestors0(n int, edges [][]int) [][]int {
	answer := make([][]int, n)
	m := make(map[int][]int, n)
	for _, edge := range edges {
		if _, ok := m[edge[1]]; ok {
			m[edge[1]] = append(m[edge[1]], edge[0])
		} else {
			m[edge[1]] = []int{edge[0]}
		}
	}
	for i := 0; i < n; i++ {
		answer[i] = []int{}
		if _, ok := m[i]; ok {
			recur(m, &answer[i], i)
			sort.Ints(answer[i])
			for j := 1; j < len(answer[i]); j++ {
				if answer[i][j] == answer[i][j-1] {
					answer[i] = append(answer[i][:j], answer[i][j+1:]...)
					j--
				}
			}
		}
	}
	return answer
}

func recur(m map[int][]int, res *[]int, x int) {
	if v, ok := m[x]; ok {
		for i := 0; i < len(v); i++ {
			*res = append(*res, v[i])
			recur(m, res, v[i])
		}
	}
}
