package p3016

import "sort"

func minimumPushes(word string) int {
	alphabet := make([]int, 26)
	for i := 0; i < len(word); i++ {
		alphabet[int(word[i])-'a']++
	}
	sort.Ints(alphabet)
	res := 0
	for i := 25; i >= 0; i-- {
		res += alphabet[i] * ((25-i)/8 + 1)
	}
	return res
}
