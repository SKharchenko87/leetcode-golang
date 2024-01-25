package p1

import "sort"

func minimumPushes(word string) int {
	arr := make([]int, 26)
	l := len(word)
	for i := 0; i < l; i++ {
		arr[word[i]-'a']++
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	count := 0
	for i, v := range arr {
		k := i/8 + 1
		count = count + v*k
	}
	return count
}
