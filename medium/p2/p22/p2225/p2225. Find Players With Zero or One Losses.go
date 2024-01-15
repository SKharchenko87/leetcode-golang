package p2225

import (
	"sort"
)

func findWinners(matches [][]int) [][]int {
	gamers := map[int]int{}
	keys := []int{}
	for _, match := range matches {
		winner := match[0]
		if _, ok := gamers[winner]; !ok {
			keys = append(keys, winner)
			gamers[winner] = 0
		}
		looser := match[1]
		if _, ok := gamers[looser]; !ok {
			keys = append(keys, looser)
		}
		gamers[looser]++
	}
	arr0 := []int{}
	arr1 := []int{}
	for k, gamer := range gamers {
		switch gamer {
		case 0:
			arr0 = append(arr0, k)
		case 1:
			arr1 = append(arr1, k)
		}
	}
	sort.Ints(arr0)
	sort.Ints(arr1)
	return [][]int{arr0, arr1}
	//ToDo сделать бейчмарк со следующим варантом и вариантом с массивом [10001]int
	//sort.Ints(keys)
	//for _, k := range keys {
	//	switch gamers[k] {
	//	case 0:
	//		arr0 = append(arr0, k)
	//	case 1:
	//		arr1 = append(arr1, k)
	//	}
	//}
	//return [][]int{arr0, arr1}
}
