package p0884

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	m := make(map[string]int)
	for _, word := range strings.Split(s1, " ") {
		m[word]++
	}
	for _, word := range strings.Split(s2, " ") {
		m[word]++
	}
	res := make([]string, 0, len(m))
	for s, i := range m {
		if i < 2 {
			res = append(res, s)
		}
	}
	return res
}
