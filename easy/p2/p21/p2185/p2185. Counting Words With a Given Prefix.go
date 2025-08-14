package p2185

import "strings"

func prefixCount(words []string, pref string) int {
	result := 0
	lenPref := len(pref)
	for _, word := range words {
		if lenPref <= len(word) && word[:lenPref] == pref {
			result++
		}
	}
	return result
}

func prefixCount0(words []string, pref string) int {
	result := 0
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			result++
		}
	}
	return result
}

func prefixCount1(words []string, pref string) int {
	result := 0
	lenPref := len(pref)
LOOP:
	for _, word := range words {
		if lenPref <= len(word) {
			for i := range pref {
				if word[i] != pref[i] {
					continue LOOP
				}
			}
			result++
		}
	}
	return result
}
