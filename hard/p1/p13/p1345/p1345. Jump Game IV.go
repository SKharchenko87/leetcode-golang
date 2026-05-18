package p1345

/*ToDo*/

func _minJumps(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return n - 1
	}
	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		if _, ok := m[arr[i]]; !ok {
			m[arr[i]] = i
		}
	}
	tmp := make([]int, 2, n)
	tmp[0] = 0
	tmp[1] = 1

	for i := 0; i < len(tmp); i++ {
		index := tmp[i]
		x := m[arr[index]] + 1

		if l := index - 1; l >= 0 && m[arr[l]] > x {
			m[arr[l]] = x
			tmp = append(tmp, l)
		}

		if r := index + 1; r < n && m[arr[r]] >= x {
			m[arr[r]] = x
			tmp = append(tmp, r)
		}
	}

	return m[arr[n-1]]
}

func minJumps(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	// Мапа для быстрого поиска всех индексов с одинаковыми значениями
	m := make(map[int][]int)
	for i, v := range arr {
		m[v] = append(m[v], i)
	}

	// Массив для хранения минимального количества шагов до каждого индекса
	// Инициализируем -1 (не посещено)
	counts := make([]int, n)
	for i := range counts {
		counts[i] = -1
	}
	counts[0] = 0

	// Очередь для BFS
	queue := make([]int, 0, n)
	queue = append(queue, 0)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Если дошли до конца, возвращаем результат
		if curr == n-1 {
			return counts[curr]
		}

		nextSteps := counts[curr] + 1

		// Вариант 1: Прыжок вправо (curr + 1)
		if curr+1 < n && counts[curr+1] == -1 {
			counts[curr+1] = nextSteps
			queue = append(queue, curr+1)
		}

		// Вариант 2: Прыжок влево (curr - 1)
		if curr-1 >= 0 && counts[curr-1] == -1 {
			counts[curr-1] = nextSteps
			queue = append(queue, curr-1)
		}

		// Вариант 3: Прыжки по одинаковым значениям
		if indices, ok := m[arr[curr]]; ok {
			for _, idx := range indices {
				if counts[idx] == -1 {
					counts[idx] = nextSteps
					queue = append(queue, idx)
				}
			}
			// КРИТИЧНО ДЛЯ ХИТРЫХ ТЕСТОВ И ТLЕ:
			// Удалили значение из мапы, чтобы больше никогда не перебирать эти индексы
			delete(m, arr[curr])
		}
	}

	return counts[n-1]
}
