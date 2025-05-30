package p2359

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	if node1 == node2 {
		return node1
	}
	l := len(edges)
	dist1, dist2 := make([]int, l), make([]int, l)

	var traverse = func(node, end int, w []int) {
		index := node
		w[index] = 1
		var tmp int
		for index != end {
			tmp = w[index]
			index = edges[index]
			if index == -1 || w[index] > 0 {
				break
			}
			w[index] = tmp + 1
		}
	}

	traverse(node1, node2, dist1)
	traverse(node2, node1, dist2)

	maxDist := 2 * l
	resIndex := -1
	for i := 0; i < l; i++ {
		if dist1[i] > 0 && dist2[i] > 0 {
			if max(dist1[i], dist2[i]) == maxDist {
				resIndex = min(i, resIndex)
			} else if max(dist1[i], dist2[i]) < maxDist {
				maxDist = max(dist1[i], dist2[i])
				resIndex = i
			}
		}
	}
	return resIndex
}

func closestMeetingNode1(edges []int, node1, node2 int) int {
	if node1 == node2 {
		return node1
	}

	n := len(edges)
	dist1, dist2 := make([]int, n), make([]int, n)

	// Инициализация расстояний
	for i := range dist1 {
		dist1[i], dist2[i] = -1, -1
	}

	// Обход графа и заполнение расстояний от заданной вершины
	traverse := func(start int, dist []int) {
		for d, current := 0, start; current != -1 && dist[current] == -1; d++ {
			dist[current] = d
			current = edges[current]
		}
	}

	traverse(node1, dist1)
	traverse(node2, dist2)

	// Поиск общей вершины с минимальным максимальным расстоянием
	minDist := n * 2
	result := -1

	for i := 0; i < n; i++ {
		if dist1[i] != -1 && dist2[i] != -1 {
			maxDist := max(dist1[i], dist2[i])
			if maxDist < minDist || (maxDist == minDist && i < result) {
				minDist = maxDist
				result = i
			}
		}
	}

	return result
}

func closestMeetingNode0(edges []int, node1 int, node2 int) int {
	if node1 == node2 {
		return node1
	}
	l := len(edges)
	w1, w2 := make([]int, l), make([]int, l)
	index1 := node1
	w1[index1] = 1
	for {
		tmp := w1[index1]
		index1 = edges[index1]
		if index1 == -1 || w1[index1] > 0 {
			break
		}
		w1[index1] = tmp + 1
		if index1 == node2 {
			break
		}
	}

	index2 := node2
	w2[index2] = 1
	for {
		tmp := w2[index2]
		index2 = edges[index2]
		if index2 == -1 || w2[index2] > 0 {
			break
		}
		w2[index2] = tmp + 1
		if index2 == node1 {
			break
		}
	}
	sumMin := 2 * l
	resIndex := -1
	for i := 0; i < l; i++ {
		if w1[i] > 0 && w2[i] > 0 && max(w1[i], w2[i]) <= sumMin {
			if max(w1[i], w2[i]) == sumMin {
				resIndex = min(i, resIndex)
			} else {
				sumMin = max(w1[i], w2[i])
				resIndex = i
			}
		}
	}
	return resIndex
}
