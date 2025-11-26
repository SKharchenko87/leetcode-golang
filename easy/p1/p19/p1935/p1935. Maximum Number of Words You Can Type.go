package p1935

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Fields(text)
	res := len(words)
	for _, word := range words {
		if strings.ContainsAny(word, brokenLetters) {
			res--
		}
	}
	return res
}
