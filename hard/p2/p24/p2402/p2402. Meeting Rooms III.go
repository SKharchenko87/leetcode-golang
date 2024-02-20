//ToDo разобраться

package p2402

import "slices"

func getEarliestEmptyRoom(roomsOccupiedUntil []int, neededAt int) int {
	earliestEmptyRoom := 0
	for i := range roomsOccupiedUntil {
		if roomsOccupiedUntil[i] <= neededAt {
			return i
		}
		if roomsOccupiedUntil[i] < roomsOccupiedUntil[earliestEmptyRoom] {
			earliestEmptyRoom = i
		}
	}
	return earliestEmptyRoom
}

func mostBooked(n int, meetings [][]int) int {
	slices.SortFunc[[][]int](meetings, func(a, b []int) int {
		return a[0] - b[0]
	})

	roomsOccupiedUntil := make([]int, n)
	roomsMeetingCount := make([]int, n)

	var emptyRoom, duration int

	for _, meeting := range meetings {
		emptyRoom = getEarliestEmptyRoom(roomsOccupiedUntil, meeting[0])
		if roomsOccupiedUntil[emptyRoom] <= meeting[0] {
			roomsOccupiedUntil[emptyRoom] = meeting[1]
		} else {
			duration = meeting[1] - meeting[0]
			roomsOccupiedUntil[emptyRoom] += duration
		}
		roomsMeetingCount[emptyRoom]++
		//fmt.Println(meeting, roomsOccupiedUntil)
	}
	maxMeetingsRoom := 0
	for i := range roomsMeetingCount {
		if roomsMeetingCount[i] > roomsMeetingCount[maxMeetingsRoom] {
			maxMeetingsRoom = i
		}
	}
	//fmt.Println(roomsMeetingCount)
	return maxMeetingsRoom
}
