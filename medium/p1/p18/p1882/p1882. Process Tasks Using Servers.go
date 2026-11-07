package p1882

import "container/heap"

type server struct {
	weight int
	index  int
	time   int
}

type MinHeap []server

// Все методы используют указатель на MinHeap
func (h *MinHeap) Len() int { return len(*h) }
func (h *MinHeap) Less(i, j int) bool {
	if (*h)[i].time != (*h)[j].time {
		return (*h)[i].time < (*h)[j].time
	}
	if (*h)[i].weight == (*h)[j].weight {
		// Исправлено: сравниваем index серверов, а не индексы слайса
		return (*h)[i].index < (*h)[j].index
	}
	return (*h)[i].weight < (*h)[j].weight
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(server))
}

func (h *MinHeap) Pop() any {
	l := h.Len()
	x := (*h)[l-1]
	*h = (*h)[:l-1]
	return x
}

func (h *MinHeap) peek() server {
	return (*h)[0] // Исправлено: минимум всегда на нулевом индексе
}

func assignTasks(servers []int, tasks []int) []int {
	// Создаем срез с capacity, но длиной 0
	freeSlice := make([]server, 0, len(servers))
	hf := MinHeap(freeSlice)
	for i, s := range servers {
		hf = append(hf, server{weight: s, index: i, time: 0})
	}
	heap.Init(&hf) // Инициализируем кучу разом за O(N)

	hb := MinHeap{}

	res := make([]int, len(tasks))
	time := 0

	for i, t := range tasks {
		// Время не может быть меньше индекса текущей задачи
		time = max(time, i)

		// Если свободных серверов нет, перематываем время
		// к моменту освобождения ближайшего занятого сервера
		if hf.Len() == 0 && hb.Len() > 0 {
			time = max(time, hb.peek().time)
		}

		// Освобождаем ВСЕ серверы, которые закончили работу к текущему моменту
		for hb.Len() > 0 && hb.peek().time <= time {
			s := heap.Pop(&hb).(server)
			s.time = 0 // Обнуляем время для кучи свободных серверов
			heap.Push(&hf, s)
		}

		// Берем лучший свободный сервер и назначаем ему задачу
		s := heap.Pop(&hf).(server)
		res[i] = s.index
		s.time = time + t
		heap.Push(&hb, s)
	}

	return res
}
