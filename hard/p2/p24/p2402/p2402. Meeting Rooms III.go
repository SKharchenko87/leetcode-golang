package p2402

import (
	"container/heap"
	"sort"
)

type Room struct {
	index    int
	freeTime int
}

type Rooms []Room

func (rooms Rooms) Len() int {
	return len(rooms)
}
func (rooms Rooms) Less(i, j int) bool {
	if rooms[i].freeTime == rooms[j].freeTime {
		return rooms[i].index < rooms[j].index
	}
	return rooms[i].freeTime < rooms[j].freeTime
}
func (rooms Rooms) Swap(i, j int) {
	rooms[i], rooms[j] = rooms[j], rooms[i]
}
func (rooms *Rooms) Push(x interface{}) {
	*rooms = append(*rooms, x.(Room))
}
func (rooms *Rooms) Pop() interface{} {
	old := *rooms
	n := len(old)
	room := old[n-1]
	*rooms = old[0 : n-1]
	return room
}

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	var rooms, freeRooms Rooms
	rooms = make([]Room, 0, n)
	freeRooms = make([]Room, n)

	cnt := make([]int, n)
	l := len(meetings)
	for i := 0; i < n; i++ {
		freeRooms[i] = Room{i, 0}
	}
	heap.Init(&rooms)
	heap.Init(&freeRooms)

	var room Room
	i := 0
LOOP:
	for i < l {
		start := meetings[i][0]
		end := meetings[i][1]
		/* Если нет текущих встреч сразу берем из свободных комнат и переходим к селдующему событию*/
		if rooms.Len() == 0 {
			room = heap.Pop(&freeRooms).(Room)
			room.freeTime = end
			heap.Push(&rooms, room)
			cnt[room.index]++
			i++
			continue
		}

		/* Проверяем занятые комнаты сравнивая их освобождение и старта текущего события,
		если комната освобождается до старта, то помещаем ее в список свободных комнат.
		Таким образом первая в очереди будет комната которая освободится первой после старта события
		*/
		room = heap.Pop(&rooms).(Room)
		for room.freeTime <= start {
			room.freeTime = 0
			heap.Push(&freeRooms, room)
			if rooms.Len() == 0 {
				continue LOOP
			}
			room = heap.Pop(&rooms).(Room)
		}

		if freeRooms.Len() > 0 {
			/* Есть свободные комнаты. Возращаем в список занятых первую из очереди.
			И берем комнату из очереди свободных*/
			heap.Push(&rooms, room)
			room = heap.Pop(&freeRooms).(Room)
			room.freeTime = end
			heap.Push(&rooms, room)
			cnt[room.index]++
		} else {
			/* Нет свободных комнат, значит первая из очереди.
			Устанавливаем время от момента когда комната освободится
			*/
			room.freeTime += end - start
			heap.Push(&rooms, room)
			cnt[room.index]++
		}

		i++
	}
	res := 0
	resMax := cnt[0]
	for i := 1; i < n; i++ {
		if resMax < cnt[i] {
			resMax = cnt[i]
			res = i
		}
	}
	return res
}
