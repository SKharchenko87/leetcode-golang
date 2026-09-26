package p1807

import (
	"iter"
	"maps"
	"strings"
)

/*TLE*/
func evaluate0(s string, knowledge [][]string) string {
	for _, row := range knowledge {
		s = strings.ReplaceAll(s, "("+row[0]+")", row[1])
	}
	i := strings.IndexRune(s, '(')
	for i >= 0 {
		j := strings.IndexRune(s[i:], ')')
		s = strings.ReplaceAll(s, s[i:i+j+1], "?")
		i = strings.IndexRune(s, '(')
	}
	return s
}

func evaluate(s string, knowledge [][]string) string {
	res := make([]string, 0, len(s))
	knowledgeMap := maps.Collect(pairs(knowledge))
	wordStart := -1
	bracketStart := -1
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			res = append(res, s[wordStart+1:i])
			bracketStart = i
		case ')':
			wordStart = i
			if v, ok := knowledgeMap[s[bracketStart+1:i]]; ok {
				res = append(res, v)
			} else {
				res = append(res, "?")
			}
		}
	}
	if wordStart < len(s) {
		res = append(res, s[wordStart+1:])
	}
	return strings.Join(res, "")
}

func pairs(s [][]string) iter.Seq2[string, string] {
	return func(yield func(k, v string) bool) {
		for i := 0; i < len(s); i++ {
			if !yield(s[i][0], s[i][1]) {
				return
			}
		}
	}
}
