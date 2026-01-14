package p3454

import "sort"

// ToDo Extra Hard

// Event представляет собой горизонтальную границу квадрата
type Event struct {
	y      float64
	x1, x2 float64
	typ    int // 1 для нижней границы (добавление), -1 для верхней (удаление)
}

// Node узла дерева отрезков
type Node struct {
	count  int     // Сколько квадратов полностью покрывают этот диапазон
	length float64 // Суммарная длина покрытия в этом диапазоне (с учетом наслоений)
}

func separateSquares(squares [][]int) float64 {
	n := len(squares)
	events := make([]Event, 0, 2*n)
	xSet := make(map[float64]struct{})

	for _, s := range squares {
		x, y, l := float64(s[0]), float64(s[1]), float64(s[2])
		events = append(events, Event{y, x, x + l, 1})
		events = append(events, Event{y + l, x, x + l, -1})
		xSet[x] = struct{}{}
		xSet[x+l] = struct{}{}
	}

	// Сортируем события по Y (снизу вверх)
	sort.Slice(events, func(i, j int) bool {
		return events[i].y < events[j].y
	})

	// Собираем и сортируем уникальные X для сжатия координат
	uniqueX := make([]float64, 0, len(xSet))
	for x := range xSet {
		uniqueX = append(uniqueX, x)
	}
	sort.Float64s(uniqueX)

	// Маппинг X-координаты в индекс для дерева отрезков
	xIndex := make(map[float64]int)
	for i, x := range uniqueX {
		xIndex[x] = i
	}

	m := len(uniqueX)
	tree := make([]Node, 4*m)

	// Функция обновления дерева отрезков
	var update func(v, tl, tr, l, r, val int)
	update = func(v, tl, tr, l, r, val int) {
		if l > r {
			return
		}
		if l == tl && r == tr {
			tree[v].count += val
		} else {
			tm := (tl + tr) / 2
			update(2*v, tl, tm, l, min(r, tm), val)
			update(2*v+1, tm+1, tr, max(l, tm+1), r, val)
		}

		// Расчет длины покрытия для текущего узла
		if tree[v].count > 0 {
			// Если count > 0, значит весь интервал [uniqueX[tl], uniqueX[tr+1]] покрыт
			tree[v].length = uniqueX[tr+1] - uniqueX[tl]
		} else {
			// Если count == 0, длина равна сумме длин детей (если не лист)
			if tl != tr {
				tree[v].length = tree[2*v].length + tree[2*v+1].length
			} else {
				tree[v].length = 0
			}
		}
	}

	// ПЕРВЫЙ ПРОХОД: Считаем общую площадь
	var totalArea float64
	for i := 0; i < len(events)-1; i++ {
		e := events[i]
		update(1, 0, m-2, xIndex[e.x1], xIndex[e.x2]-1, e.typ)

		// tree[1].length — это текущая ширина объединения квадратов по X
		height := events[i+1].y - e.y
		totalArea += height * tree[1].length
	}

	// Очищаем дерево перед вторым проходом
	for i := range tree {
		tree[i] = Node{}
	}

	// ВТОРОЙ ПРОХОД: Ищем линию, делящую площадь пополам
	targetArea := totalArea / 2.0
	var currentArea float64

	for i := 0; i < len(events)-1; i++ {
		e := events[i]
		update(1, 0, m-2, xIndex[e.x1], xIndex[e.x2]-1, e.typ)

		width := tree[1].length
		height := events[i+1].y - e.y
		sliceArea := width * height

		if currentArea+sliceArea >= targetArea {
			// Нужная линия находится внутри текущего интервала Y
			// Используем линейную интерполяцию:
			// targetArea = currentArea + width * (y_ans - events[i].y)
			offset := (targetArea - currentArea) / width
			return events[i].y + offset
		}
		currentArea += sliceArea
	}

	return events[len(events)-1].y
}
