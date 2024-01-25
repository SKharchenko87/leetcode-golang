package p3

import "sort"

func minimumPushes(word string) int {
	arr := make([][2]int, 26)
	m := map[byte]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0, 'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0}
	l := len(word)
	for i := 0; i < l; i++ {
		m[word[i]]++
		arr[word[i]-'a'] = [2]int{int(word[i]), m[word[i]]}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	count := 0
	for i, v := range arr {
		k := i/8 + 1
		count = count + v[1]*k
	}
	return count
}
