package p1976

import "container/heap"

const mod = 1e9 + 7

type Edge struct {
	node int
	time int64
}

type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].time < pq[j].time
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Edge)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func countPaths(n int, roads [][]int) int {
	// Создаем граф
	graph := make([][]Edge, n)
	for _, road := range roads {
		u, v, time := road[0], road[1], int64(road[2])
		graph[u] = append(graph[u], Edge{v, time})
		graph[v] = append(graph[v], Edge{u, time})
	}

	// Инициализируем массивы для хранения минимального времени и количества способов
	dist := make([]int64, n)
	for i := range dist {
		dist[i] = 1 << 62 // Устанавливаем начальное значение как "бесконечность"
	}
	ways := make([]int, n)
	dist[0] = 0
	ways[0] = 1

	// Инициализируем приоритетную очередь
	pq := &PriorityQueue{}
	heap.Push(pq, &Edge{0, 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Edge)
		u := current.node
		currentTime := current.time

		if currentTime > dist[u] {
			continue
		}

		for _, edge := range graph[u] {
			v := edge.node
			time := edge.time

			if dist[v] > dist[u]+time {
				dist[v] = dist[u] + time
				ways[v] = ways[u]
				heap.Push(pq, &Edge{v, dist[v]})
			} else if dist[v] == dist[u]+time {
				ways[v] = (ways[v] + ways[u]) % mod
			}
		}
	}

	return ways[n-1]
}
