package p2410

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	lp, lt := len(players), len(trainers)
	i, j := 0, 0
	result := 0
	for i < lp && j < lt {
		if players[i] <= trainers[j] {
			result++
			i++
		}
		j++
	}
	return result
}
