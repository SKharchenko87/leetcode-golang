package p1857

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)

	// Строим граф
	graph := make([][]int, n)
	indegree := make([]int, n)

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
		indegree[to]++
	}

	// dp[i][c] = максимальное количество цвета c на ЛЮБОМ пути, заканчивающемся в узле i
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 26)
	}

	// Топологическая сортировка + DP одновременно
	queue := []int{}
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
			// ТОЛЬКО сейчас инициализируем цвет стартового узла!
			colorIndex := int(colors[i] - 'a')
			dp[i][colorIndex] = 1
		}
	}

	processed := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		processed++

		// Обрабатываем всех соседей
		for _, neighbor := range graph[curr] {
			neighborColor := int(colors[neighbor] - 'a')

			// Обновляем DP для соседа
			for color := 0; color < 26; color++ {
				if color == neighborColor {
					// Если цвет соседа совпадает с этим цветом, добавляем +1
					dp[neighbor][color] = max(dp[neighbor][color], dp[curr][color]+1)
				} else {
					// Иначе просто переносим количество
					dp[neighbor][color] = max(dp[neighbor][color], dp[curr][color])
				}
			}

			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Проверяем цикл
	if processed != n {
		return -1
	}

	// Находим максимум
	result := 0
	for i := 0; i < n; i++ {
		for color := 0; color < 26; color++ {
			result = max(result, dp[i][color])
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TLE
func largestPathValue0(colors string, edges [][]int) int {
	n := len(colors)
	graph := make(map[int]map[int]struct{}, len(colors))
	calc := make([][26]int, n)
	startPosition := make([]int, n)
	for _, edge := range edges {
		if _, ok := graph[edge[0]]; !ok {
			graph[edge[0]] = make(map[int]struct{})
		}
		graph[edge[0]][edge[1]] = struct{}{}
		startPosition[edge[1]] = -1
	}

	var dfs func(u int, prev [26]int, prevIndex int) int
	dfs = func(u int, prev [26]int, prevIndex int) int {
		//if u <= prevIndex {
		//	return -1
		//}
		calc[u] = prev
		calc[u][colors[u]-'a']++
		res := 0
		for v := range graph[u] {
			x := dfs(v, calc[u], u)
			if x == -1 {
				return -1
			}
			res = max(x, res)
		}
		if len(graph[u]) == 0 {
			for i := 0; i < 26; i++ {
				res = max(calc[u][i], res)
			}
		}
		return res
	}

	result := -1
	for i := 0; i < n; i++ {
		if startPosition[i] == 0 {
			x := dfs(i, [26]int{}, -1)
			if x == -1 {
				return -1
			}
			result = max(result, x)
		}
	}
	return result
}
