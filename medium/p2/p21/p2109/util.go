package p2109

import "strings"

func getIndexSpaces(text string) []int {
	result := make([]int, 0)
	index := 0
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			result = append(result, index)
		} else {
			index++
		}
	}
	return result
}

func deleteSpace(text string) string {
	return strings.ReplaceAll(text, " ", "")
}
