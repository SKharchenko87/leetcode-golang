package p1514

import "math"

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	graph := make(map[int]map[int]float64, n)
	for i := 0; i < len(edges); i++ {
		if _, ok := graph[edges[i][0]]; !ok {
			graph[edges[i][0]] = make(map[int]float64)
		}
		graph[edges[i][0]][edges[i][1]] = succProb[i]

		if _, ok := graph[edges[i][1]]; !ok {
			graph[edges[i][1]] = make(map[int]float64)
		}
		graph[edges[i][1]][edges[i][0]] = succProb[i]
	}
	candidate := make(map[int]float64, n)
	candidate = graph[start_node]
	visited := make(map[int]float64, n)
	visited = graph[start_node]
	for len(candidate) != 0 {
		newCandidate := make(map[int]float64, n)
		for k, v := range candidate {
			for i, f := range graph[k] {
				if f*v > visited[i] {
					visited[i] = f * v
					if i != end_node {
						newCandidate[i] = f * v
					}
				}
			}
		}
		candidate = newCandidate
	}
	return math.Round(visited[end_node]*100_000) / 100_000
}

func maxProbability0(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	visited := make(map[int]float64, n)
	candidate := make(map[int]float64, n)
	for i := 0; i < len(edges); i++ {
		if edges[i][0] == start_node {
			visited[edges[i][1]] = succProb[i]
			if edges[i][1] != end_node {
				candidate[edges[i][1]] = succProb[i]
			}
		}
		if edges[i][1] == start_node {
			visited[edges[i][0]] = succProb[i]
			if edges[i][0] != end_node {
				candidate[edges[i][0]] = succProb[i]
			}
		}
	}
	for len(candidate) != 0 {
		newCandidate := make(map[int]float64, n)
		for k, v := range candidate {
			for i := 0; i < len(edges); i++ {
				if edges[i][0] == k && succProb[i]*v > visited[edges[i][1]] {
					visited[edges[i][1]] = succProb[i] * v
					if edges[i][1] != end_node {
						newCandidate[edges[i][1]] = succProb[i] * v
					}
				}
				if edges[i][1] == k && succProb[i]*v > visited[edges[i][0]] {
					visited[edges[i][0]] = succProb[i] * v
					if edges[i][0] != end_node {
						newCandidate[edges[i][0]] = succProb[i] * v
					}
				}
			}
		}
		candidate = newCandidate
	}
	return math.Round(visited[end_node]*100_000) / 100_000
}
