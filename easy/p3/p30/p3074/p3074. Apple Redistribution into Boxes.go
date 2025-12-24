package p3074

import "sort"

// O(a)+O(clogc) | O(1)
func minimumBoxes(apple []int, capacity []int) int {
	sum := 0
	for i := 0; i < len(apple); i++ {
		sum += apple[i]
	}
	sort.Ints(capacity)
	i := len(capacity) - 1
	for ; i >= 0 && sum > 0; i-- {
		sum -= capacity[i]
	}
	return len(capacity) - i - 1
}

// O(a)+O(clogc) | O(c)
func minimumBoxes0(apple []int, capacity []int) int {
	sum := 0
	for i := 0; i < len(apple); i++ {
		sum += apple[i]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(capacity)))
	i := 0
	for ; i < len(capacity) && sum > 0; i++ {
		sum -= capacity[i]
	}
	return i
}
