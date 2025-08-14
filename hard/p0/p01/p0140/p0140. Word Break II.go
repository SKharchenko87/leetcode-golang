package p0140

import (
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	l := len(s)
	p := make([][]int, l)
	for _, word := range wordDict {
		lenSubstring := len(word)
		prevIndex := -1
		index := strings.Index(s, word)
		for prevIndex != index {
			if p[index] == nil {
				p[index] = []int{}
			}
			p[index] = append(p[index], index+lenSubstring)
			prevIndex = index
			index = prevIndex + strings.Index(s[prevIndex+1:], word) + 1
		}
	}
	res := []string{}
	var recursive func(start int, candidate string)
	recursive = func(start int, candidate string) {
		for i := 0; i < len(p[start]); i++ {
			if p[start] != nil {
				newStart := p[start][i]
				if newStart == l {
					res = append(res, strings.TrimSpace(candidate+" "+s[start:newStart]))
				} else {
					recursive(newStart, candidate+" "+s[start:newStart])
				}
			}
		}
	}
	recursive(0, "")
	return res
}

func findPositions(s, substring string) []int {
	res := []int{}
	prevIndex := -1
	index := strings.Index(s, substring)
	for prevIndex != index {
		res = append(res, index)
		prevIndex = index
		index = prevIndex + strings.Index(s[prevIndex+1:], substring) + 1
	}
	return res
}
