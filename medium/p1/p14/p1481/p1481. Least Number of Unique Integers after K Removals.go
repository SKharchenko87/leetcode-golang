package p1481

import (
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	l := len(arr)
	if l <= k+1 {
		return l - k
	}
	m := make(map[int]int, l)
	for _, e := range arr {
		m[e]++
	}
	i, lm := 0, len(m)
	counts := make([]int, lm)
	for _, v := range m {
		counts[i] = v
		i++
	}
	sort.Ints(counts)
	for _, cnt := range counts {
		k -= cnt
		if k < 0 {
			break
		}
		lm--
		if k == 0 {
			break
		}
	}
	return lm
}

func findLeastNumOfUniqueInts1(arr []int, k int) int {
	m := make(map[int]int, len(arr))
	keys := []int{}
	for _, e := range arr {
		m[e]++
		if m[e] == 1 {
			keys = append(keys, e)
		}
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})
	for _, key := range keys {
		tmp := k - m[key]
		if tmp < 0 {
			break
		}
		delete(m, key)
		if tmp == 0 {
			break
		} else {
			k = tmp
		}
	}
	return len(m)
}
