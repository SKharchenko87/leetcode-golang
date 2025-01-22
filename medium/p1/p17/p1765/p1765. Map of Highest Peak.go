package p1765

import "container/heap"

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				res[i][j] = 0
			} else {
				res[i][j] = 1_000_001
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a, b, c := res[i][j], 1_000_001, 1_000_001
			if i-1 >= 0 {
				b = res[i-1][j] + 1
			}
			if j-1 >= 0 {
				c = res[i][j-1] + 1
			}
			res[i][j] = min(a, b, c)
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			a, b, c := res[i][j], 1_000_001, 1_000_001
			if i+1 < m {
				b = res[i+1][j] + 1
			}
			if j+1 < n {
				c = res[i][j+1] + 1
			}
			res[i][j] = min(a, b, c)
		}
	}
	return res
}

type Point struct {
	i, j   int
	height int
}

type PriorityQueue []*Point

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].height < pq[j].height
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Point))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func highestPeak1(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	x := m * n
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	pq := PriorityQueue{}
	heap.Init(&pq)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				heap.Push(&pq, &Point{i, j, 0})
				x--
			}
		}
	}
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for pq.Len() > 0 {
		cell := heap.Pop(&pq).(*Point)
		for _, direction := range directions {
			i, j := cell.i+direction[0], cell.j+direction[1]
			if 0 <= i && i < m && 0 <= j && j < n && res[i][j] == 0 && isWater[i][j] != 1 {
				res[i][j] = cell.height + 1
				heap.Push(&pq, &Point{i, j, res[i][j]})
				x--
				if x == 0 {
					return res
				}
			}
		}
	}
	return res
}

func highestPeak0(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	pq := PriorityQueue{}
	heap.Init(&pq)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				heap.Push(&pq, &Point{i, j, 0})
			}
		}
	}
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for pq.Len() > 0 {
		cell := heap.Pop(&pq).(*Point)
		for _, direction := range directions {
			i, j := cell.i+direction[0], cell.j+direction[1]
			if 0 <= i && i < m && 0 <= j && j < n && res[i][j] == 0 && isWater[i][j] != 1 {
				res[i][j] = cell.height + 1
				heap.Push(&pq, &Point{i, j, res[i][j]})
			}
		}
	}
	return res
}
