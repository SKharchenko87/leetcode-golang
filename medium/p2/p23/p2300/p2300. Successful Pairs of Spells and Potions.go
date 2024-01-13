package p2300

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	n, m := len(spells), len(potions)
	res := make([]int, n)
	sort.Ints(potions)
	for i, spell := range spells {
		l, r := 0, m
		for r-l > 1 {
			j := (l + r) / 2
			if int64(spell)*int64(potions[j]) >= success {
				r = j
			} else {
				l = j
			}
		}
		if int64(spell)*int64(potions[l]) >= success {
			res[i] = m - l
		} else {
			res[i] = m - r
		}
	}
	return res
}
