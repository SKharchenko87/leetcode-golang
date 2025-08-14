package p0802

import (
	"sort"
)

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	safe := make([]int, n) // 0: не посещён; 1: посещён, не безопасный; 2: безопасный
	res := []int{}

	// DFS для проверки, является ли узел безопасным
	var isSafe func(int) bool
	isSafe = func(node int) bool {
		// Если узел уже посещён, проверяем его статус
		if safe[node] != 0 {
			return safe[node] == 2
		}
		// Помечаем узел как посещённый, но ещё не безопасный
		safe[node] = 1
		// Проверяем всех соседей
		for _, neighbor := range graph[node] {
			if !isSafe(neighbor) { // Если хотя бы один сосед не безопасен
				return false
			}
		}
		// Если все соседи безопасны, помечаем узел как безопасный
		safe[node] = 2
		return true
	}

	// Проверяем каждый узел
	for i := 0; i < n; i++ {
		if isSafe(i) {
			res = append(res, i)
		}
	}

	return res
}

func eventualSafeNodes2(graph [][]int) []int {
	l := len(graph)
	lengthsGraph := make([]int, l)
	reversGraph := make([][]int, l)
	candidates := make([]int, 0, l)

	for i := 0; i < l; i++ {
		lengthsGraph[i] = len(graph[i])
		if lengthsGraph[i] == 0 {
			candidates = append(candidates, i)
		}
	}
	for i := 0; i < l; i++ {
		for _, j := range graph[i] {
			if i != j {
				reversGraph[j] = append(reversGraph[j], i)
			}
		}
	}

	res := make([]int, 0, l)
	for index := 0; index < len(candidates); index++ {
		candidate := candidates[index]
		res = append(res, candidate)
		for _, j := range reversGraph[candidate] {
			lengthsGraph[j]--
			if lengthsGraph[j] == 0 {
				candidates = append(candidates, j)
			}
		}
	}

	sort.Ints(res)
	return res
}

func eventualSafeNodes1(graph [][]int) []int {
	lengthsGraph := make([]int, len(graph))
	reversGraph := make([][]int, len(graph))
	candidates := []int{}
	res := []int{}
	for i := 0; i < len(graph); i++ {
		reversGraph[i] = []int{}
		lengthsGraph[i] = len(graph[i])
		if lengthsGraph[i] == 0 {
			candidates = append(candidates, i)
		}
	}
	for i := 0; i < len(graph); i++ {
		for _, j := range graph[i] {
			if i != j {
				reversGraph[j] = append(reversGraph[j], i)
			}
		}
	}
	newCandidate := []int{}
	for len(candidates) > 0 {
		for _, candidate := range candidates {
			res = append(res, candidate)
			for _, j := range reversGraph[candidate] {
				lengthsGraph[j]--
				//fmt.Println(j, lengthsGraph, candidates)
				if lengthsGraph[j] == 0 {
					newCandidate = append(newCandidate, j)
				}
			}
		}
		candidates, newCandidate = newCandidate, candidates[:0]
	}
	sort.Ints(res)
	return res
}

func eventualSafeNodes0(graph [][]int) []int {
	lengthsGraph := make([]int, len(graph))
	reversGraph := make([][]int, len(graph))
	candidates := []int{}
	resMap := make(map[int]bool)
	res := []int{}
	for i := 0; i < len(graph); i++ {
		reversGraph[i] = []int{}
		lengthsGraph[i] = len(graph[i])
		if lengthsGraph[i] == 0 {
			candidates = append(candidates, i)
		}
	}
	for i := 0; i < len(graph); i++ {
		for _, j := range graph[i] {
			if i != j {
				reversGraph[j] = append(reversGraph[j], i)
			}
		}
	}
	newCandidate := []int{}
	for len(candidates) > 0 {
		for _, candidate := range candidates {
			resMap[candidate] = true
			for _, j := range reversGraph[candidate] {
				lengthsGraph[j]--
				//fmt.Println(j, lengthsGraph, candidates)
				if lengthsGraph[j] == 0 {
					newCandidate = append(newCandidate, j)
				}
			}
		}
		candidates, newCandidate = newCandidate, candidates[:0]
	}
	for i := range resMap {
		res = append(res, i)
	}
	sort.Ints(res)
	return res
}
