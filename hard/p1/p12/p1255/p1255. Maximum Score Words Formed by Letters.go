package p1255

func maxScoreWords(words []string, letters []byte, score []int) int {
	l := len(words)
	result := 0
	letterArr := [26]int{}
	for _, letter := range letters {
		letterArr[letter-'a']++
	}

	var recursive func(index int, letterArray [26]int, res int) int
	recursive = func(index int, letterArray [26]int, res int) int {
		if index == l {
			return res
		}
		scoreWord := 0
		newLetterArray := [26]int{}
		copy26(&newLetterArray, &letterArray)
		for _, letter := range words[index] {
			letterIndex := letter - 'a'
			letterArray[letterIndex]--
			if letterArray[letterIndex] < 0 {
				scoreWord = 0
				break
			}
			scoreWord += score[letterIndex]
		}
		return max(res, recursive(index+1, letterArray, res+scoreWord), recursive(index+1, newLetterArray, res))
	}

	for i := 0; i < l; i++ {
		newLetterArray := [26]int{}
		copy26(&newLetterArray, &letterArr)
		res := recursive(i, newLetterArray, 0)
		if res > result {
			result = res
		}
	}
	return result
}

func copy26(dst, src *[26]int) {
	for i := 0; i < 26; i++ {
		dst[i] = src[i]
	}
}
