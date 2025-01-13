package p1408

import (
	"cmp"
	"slices"
	"strings"
)

func stringMatching(words []string) []string {
	slices.SortFunc(words, func(a, b string) int {
		if len(a) == len(b) {
			return cmp.Compare(a, b)
		}
		return len(a) - len(b)
	})
	l := len(words)
	result := make([]string, 0, l-1)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if strings.Contains(words[j], words[i]) {
				result = append(result, words[i])
				break
			}
		}
	}
	return result
}
