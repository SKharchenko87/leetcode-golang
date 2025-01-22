package p1455

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	cnt := 1
	isFind := true
	for i := 0; i < len(sentence)-len(searchWord)+1; i++ {
		if sentence[i] == ' ' {
			cnt++
			isFind = true
		} else if isFind {
			if sentence[i:i+len(searchWord)] == searchWord {
				return cnt
			} else {
				isFind = false
			}
		}
	}
	return -1
}

func isPrefixOfWord2(sentence string, searchWord string) int {
	cnt := 1
	isFind := true
	index := 0
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' {
			index = 0
			cnt++
			isFind = true
		} else if isFind {
			if searchWord[index] == sentence[i] {
				index++
				if index == len(searchWord) {
					return cnt
				}
			} else {
				isFind = false
			}
		}
	}
	return -1
}

func isPrefixOfWord1(sentence string, searchWord string) int {
	arr := strings.Split(sentence, " ")
	for i, s := range arr {
		if strings.HasPrefix(s, searchWord) {
			return i + 1
		}
	}
	return -1
}

func isPrefixOfWord0(sentence string, searchWord string) int {
	sentence = " " + sentence
	searchWord = " " + searchWord
	i := strings.Index(sentence, searchWord)
	if i == -1 {
		return -1
	}
	arr := strings.Split(sentence[:i], " ")
	return len(arr)
}
