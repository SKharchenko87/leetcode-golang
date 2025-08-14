package p2942

func findWordsContaining(words []string, x byte) []int {
	result := make([]int, 0, len(words))
	for i, word := range words {
		for j := 0; j < len(word); j++ {
			if word[j] == x {
				result = append(result, i)
				break
			}
		}
	}
	return result
}
