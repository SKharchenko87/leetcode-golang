package p2290

import "container/heap"

var directions = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

type candidateHeap struct {
	from [2]int
	to   [2]int
	cost int
}

type MinHeap []candidateHeap

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(candidateHeap))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minimumObstacles1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := map[[2]int]int{{0, 0}: grid[0][0]}

	candidates := &MinHeap{}
	heap.Init(candidates)
	cm := map[[2]int]int{}
	if m > 1 {
		heap.Push(candidates, candidateHeap{from: [2]int{0, 0}, to: [2]int{1, 0}, cost: visited[[2]int{0, 0}] + grid[1][0]})
		cm[[2]int{1, 0}] = visited[[2]int{0, 0}] + grid[1][0]
		//candidates = append(candidates, candidate{from: [2]int{0, 0}, to: [2]int{1, 0}})
	}
	if n > 1 {
		heap.Push(candidates, candidateHeap{from: [2]int{0, 0}, to: [2]int{0, 1}, cost: visited[[2]int{0, 0}] + grid[0][1]})
		cm[[2]int{0, 1}] = visited[[2]int{0, 0}] + grid[0][1]
		//candidates = append(candidates, candidate{from: [2]int{0, 0}, to: [2]int{0, 1}})
	}

	for candidates.Len() > 0 {

		c := heap.Pop(candidates).(candidateHeap)
		delete(cm, c.to)
		if v, ok := visited[c.to]; !ok || v > visited[c.from]+grid[c.to[0]][c.to[1]] {
			visited[c.to] = visited[c.from] + grid[c.to[0]][c.to[1]]
			if c.to[0] == m-1 && c.to[1] == n-1 {
				break
			}
			for _, direction := range directions {
				i := c.to[0] + direction[0]
				j := c.to[1] + direction[1]
				if i < 0 || i >= m || j < 0 || j >= n {
					continue
				}
				if v, ok := visited[[2]int{i, j}]; !ok || v > visited[[2]int{c.to[0], c.to[1]}]+grid[i][j] {
					if v, ok := cm[[2]int{i, j}]; !ok || v > visited[[2]int{c.to[0], c.to[1]}]+grid[i][j] {
						heap.Push(candidates, candidateHeap{from: [2]int{c.to[0], c.to[1]}, to: [2]int{i, j}, cost: visited[[2]int{c.to[0], c.to[1]}] + grid[i][j]})
						cm[[2]int{i, j}] = visited[[2]int{c.to[0], c.to[1]}] + grid[i][j]
					}
					//candidates = append(candidates, candidate{from: [2]int{c.to[0], c.to[1]}, to: [2]int{i, j}})
				}

			}
		}

	}
	return visited[[2]int{m - 1, n - 1}]
}

type candidate struct {
	from [2]int
	to   [2]int
}

// TLE
func minimumObstacles0(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := map[[2]int]int{{0, 0}: grid[0][0]}
	candidates := make([]candidate, 0, m*n)
	if m > 1 {
		candidates = append(candidates, candidate{from: [2]int{0, 0}, to: [2]int{1, 0}})
	}
	if n > 1 {
		candidates = append(candidates, candidate{from: [2]int{0, 0}, to: [2]int{0, 1}})
	}

	index := 0
	for index < len(candidates) {
		c := candidates[index]

		if v, ok := visited[c.to]; !ok || v > visited[c.from]+grid[c.to[0]][c.to[1]] {
			visited[c.to] = visited[c.from] + grid[c.to[0]][c.to[1]]
			for _, direction := range directions {
				i := c.to[0] + direction[0]
				j := c.to[1] + direction[1]
				if i < 0 || i >= m || j < 0 || j >= n {
					continue
				}
				if v, ok := visited[[2]int{i, j}]; !ok || v > visited[[2]int{c.to[0], c.to[1]}]+grid[i][j] {
					candidates = append(candidates, candidate{from: [2]int{c.to[0], c.to[1]}, to: [2]int{i, j}})
				}

			}
		}
		index++

	}
	return visited[[2]int{m - 1, n - 1}]
}

func minimumObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true
	currents := [][2]int{[2]int{0, 0}}
	prior1 := [][2]int{}
	step := 0
	for len(currents) > 0 {
		prior1 = prior1[:0]
		//prior1 = [][2]int{}
		for c := 0; c < len(currents); c++ {
			for _, direction := range directions {
				i := currents[c][0] + direction[0]
				j := currents[c][1] + direction[1]

				if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] {
					continue
				}

				if i == m-1 && j == n-1 {
					return step
				}

				visited[i][j] = true

				if grid[i][j] == 0 {
					currents = append(currents, [2]int{i, j})
				} else {
					prior1 = append(prior1, [2]int{i, j})
				}

			}
		}
		currents, prior1 = prior1, currents
		//currents = prior1
		step++
	}
	return -1
}
