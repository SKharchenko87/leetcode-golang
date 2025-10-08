package p2300

import (
	"math"
	"slices"
	"sort"
)

type dd struct {
	indexes []int
	spells  []int
}

func (d dd) Len() int {
	return len(d.indexes)
}

func (d dd) Less(i, j int) bool {
	return d.spells[i] > d.spells[j]
}

func (d dd) Swap(i, j int) {
	d.indexes[i], d.indexes[j] = d.indexes[j], d.indexes[i]
	d.spells[i], d.spells[j] = d.spells[j], d.spells[i]
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	n := len(spells)
	m := len(potions)
	result := make([]int, n)
	sort.Ints(potions)

	if n >= 10 {
		for i := 0; i < n; i++ {
			x := int(math.Ceil(float64(success) / float64(spells[i])))
			index, _ := slices.BinarySearch(potions, x)
			result[i] = m - index
		}
	} else {
		indexes := make([]int, n)
		for i := 0; i < n; i++ {
			indexes[i] = i
		}
		d := dd{indexes: indexes, spells: spells}
		sort.Sort(d)
		j := 0
		for i := 0; i < n; i++ {
			for ; j < m && int64(potions[j])*int64(spells[i]) < success; j++ {
			}
			result[indexes[i]] = m - j
		}
	}
	return result
}

func successfulPairs2(spells []int, potions []int, success int64) []int {
	n := len(spells)
	m := len(potions)
	indexes := make([]int, n)
	for i := 0; i < n; i++ {
		indexes[i] = i
	}
	d := dd{indexes: indexes, spells: spells}
	sort.Sort(d)
	sort.Ints(potions)
	j := 0
	result := make([]int, n)
	for i := 0; i < n; i++ {
		for ; j < m && int64(potions[j])*int64(spells[i]) < success; j++ {
		}
		result[indexes[i]] = m - j
	}
	return result

}

func successfulPairs1(spells []int, potions []int, success int64) []int {
	n := len(spells)
	m := len(potions)
	sort.Ints(potions)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		x := int(math.Ceil(float64(success) / float64(spells[i])))
		index, _ := slices.BinarySearch(potions, x)
		result[i] = m - index
	}
	return result
}

func successfulPairs0(spells []int, potions []int, success int64) []int {
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
