package p0763

func partitionLabels(s string) []int {
	// Если пустая строка возвращаем пустой срез
	l := len(s)
	if l == 0 {
		return []int{}
	}

	// Определяем последнию позицию буквы в строке
	lastLetters := [26]int{}
	for i, ch := range s {
		lastLetters[ch-'a'] = i
	}

	// Определяем длины партиций
	firstIndexPart := 0
	lastIndexPart := lastLetters[s[0]-'a']
	res := []int{}
	for i := 0; i < l; i++ {
		lastIndexPart = max(lastIndexPart, lastLetters[s[i]-'a'])
		if i == lastIndexPart {
			res = append(res, lastIndexPart-firstIndexPart+1)
			firstIndexPart = lastIndexPart + 1
		}
	}

	return res
}

func partitionLabels0(s string) []int {
	lastLetters := [26]int{}
	for i, ch := range s {
		lastLetters[ch-'a'] = i
	}
	res := []int{}
	for i := 0; i < len(s); i++ {
		lastIndexPart := lastLetters[s[i]-'a']
		j := i
		for ; j <= lastIndexPart; j++ {
			lastIndexPart = max(lastIndexPart, lastLetters[s[j]-'a'])
		}
		res = append(res, j-i)
		i = j - 1
	}
	return res
}
