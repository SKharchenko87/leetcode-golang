package p3433

import (
	"container/heap"
	"strconv"
)

const (
	OFFLINE byte = iota
	MESSAGE
)

const (
	ALL byte = iota
	HERE
	LIST
)

type Event struct {
	kind      byte
	timestamp int
	mentions  byte
	list      []int
}

func parseMentions(text string) (byte, []int) {
	switch text {
	case "ALL":
		return ALL, nil
	case "HERE":
		return HERE, nil
	default:
		res := []int{}
		var x int
		for _, ch := range text {
			if ch == ' ' {
				res = append(res, x)
				x = 0
			} else if a := ch - '0'; 0 <= a && a <= 9 {
				x = x*10 + int(a)
			}
		}
		res = append(res, x)
		return LIST, res
	}
}

type MinHeap []Event

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].timestamp == h[j].timestamp {
		return h[i].kind < h[j].kind
	}
	return h[i].timestamp < h[j].timestamp
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Event))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func countMentions(numberOfUsers int, events [][]string) []int {
	arr := make([]Event, 0, len(events))
	mh := MinHeap(arr)
	heap.Init(&mh)
	for _, event := range events {
		E := Event{}

		if event[0] == "OFFLINE" {
			E.kind = OFFLINE
		} else {
			E.kind = MESSAGE
		}

		E.timestamp, _ = strconv.Atoi(event[1])

		switch event[2] {
		case "ALL":
			E.mentions = ALL
		case "HERE":
			E.mentions = HERE
		default:
			E.mentions, E.list = parseMentions(event[2])
		}
		heap.Push(&mh, E)
	}
	offline := make([]int, numberOfUsers)
	res := make([]int, numberOfUsers)
	for mh.Len() > 0 {
		e := heap.Pop(&mh).(Event)
		switch e.mentions {
		case ALL:
			for number := 0; number < numberOfUsers; number++ {
				res[number]++
			}
		case HERE:
			for number := 0; number < numberOfUsers; number++ {
				if offline[number] <= e.timestamp {
					res[number]++
				}
			}
		case LIST:
			if e.kind == OFFLINE {
				offline[e.list[0]] = e.timestamp + 60
			} else if e.kind == MESSAGE {
				for _, number := range e.list {
					res[number]++
				}
			}
		}
	}
	return res
}
