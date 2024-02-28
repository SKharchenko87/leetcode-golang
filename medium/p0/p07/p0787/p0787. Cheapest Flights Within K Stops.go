package p0787

import (
	"fmt"
	"math"
)

// ToDo разобраться Bellman-Ford
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	prev := []int{}
	now := make([]int, n)

	for i := range now {
		now[i] = math.MaxInt32
	}

	now[src] = 0

	for k >= 0 {
		prev = now
		now = append([]int{}, prev...)

		for _, flight := range flights {
			now[flight[1]] = min(now[flight[1]], prev[flight[0]]+flight[2])
		}

		k--
	}

	if now[dst] == math.MaxInt32 {
		return -1
	}

	return now[dst]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func findCheapestPrice1(n int, flights [][]int, src int, dst int, k int) int {
	m := make(map[int]map[int]int, n)
	for _, flight := range flights {
		if _, ok := m[flight[0]]; !ok {
			m[flight[0]] = map[int]int{flight[1]: flight[2]}
		} else if m[flight[0]][flight[1]] == 0 || m[flight[0]][flight[1]] > flight[2] {
			m[flight[0]][flight[1]] = flight[2]
		}
	}
	fmt.Println(m)
	return 0
	//prev_arr := m[src]
	//arr := [][2]int{}
	//res := math.MaxInt
	//for i := 0; i < k; i++ {
	//	for j := range prev_arr {
	//		if prev_arr[j][0] == dst {
	//			arr
	//		}
	//	}
	//}
}
