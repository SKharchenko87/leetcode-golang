package p1684

func countConsistentStrings(allowed string, words []string) int {
	var allowedInt32 int32
	for _, letter := range allowed {
		allowedInt32 |= (1 << (letter - 'a'))
	}
	numberOfConsistent := 0
loop:
	for _, word := range words {
		for _, letter := range word {
			if allowedInt32&(1<<(letter-'a')) == 0 {
				continue loop
			}
		}
		numberOfConsistent++
	}
	return numberOfConsistent
}

func countConsistentStrings0(allowed string, words []string) int {
	allowedArray := [26]bool{}
	for _, letter := range allowed {
		allowedArray[letter-'a'] = true
	}
	numberOfConsistent := 0
loop:
	for _, word := range words {
		for _, letter := range word {
			if !allowedArray[letter-'a'] {
				continue loop
			}
		}
		numberOfConsistent++
	}
	return numberOfConsistent
}
