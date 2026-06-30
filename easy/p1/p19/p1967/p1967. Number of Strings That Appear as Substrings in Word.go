package p1967

import "strings"

func numOfStrings(patterns []string, word string) int {
	res := 0
	for i := 0; i < len(patterns); i++ {
		if strings.Contains(word, patterns[i]) {
			res++
		}
	}
	return res
}
