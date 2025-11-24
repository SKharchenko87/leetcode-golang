package p2785

import "sort"

var m = map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {}}

func sortVowels(s string) string {
	l := len(s)
	t := make([]byte, l)
	indexes := make([]int, 0, l)
	vowels := make([]byte, 0, l)
	for i := 0; i < l; i++ {
		if _, ok := m[s[i]]; ok {
			indexes = append(indexes, i)
			vowels = append(vowels, s[i])
		} else {
			t[i] = s[i]
		}
	}
	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})
	for i0, i1 := range indexes {
		t[i1] = vowels[i0]
	}
	return string(t)
}
