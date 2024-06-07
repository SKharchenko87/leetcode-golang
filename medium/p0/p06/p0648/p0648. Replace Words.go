package p0648

import (
	"sort"
	"strings"
)

// //Merge join
func replaceWords(dictionary []string, sentence string) string {
	sort.Strings(dictionary)
	for i := 1; i < len(dictionary); i++ {
		if strings.Index(dictionary[i], dictionary[i-1]) == 0 {
			dictionary = append(dictionary[:i], dictionary[i+1:]...)
			i--
		}
	}
	words := strings.Split(sentence, " ")
	positionMap := map[string][]int{}
	for i, word := range words {
		if arr, ok := positionMap[word]; ok {
			positionMap[word] = append(arr, i)
		} else {
			positionMap[word] = []int{i}
		}
	}
	i, j := 0, 0
	li, lj := len(words), len(dictionary)
	resArr := make([]string, li)
	sort.Strings(words)
	for i < li && j < lj {
		if strings.HasPrefix(words[i], dictionary[j]) {
			for _, position := range positionMap[words[i]] {
				resArr[position] = dictionary[j]
			}
			i++
		} else if words[i] > dictionary[j] {
			j++
		} else {
			for _, position := range positionMap[words[i]] {
				resArr[position] = words[i]
			}
			i++
		}
	}
	for ; i < li; i++ {
		for _, position := range positionMap[words[i]] {
			resArr[position] = words[i]
		}
	}
	return strings.Join(resArr, " ")
}

// Hash join
func replaceWords1(dictionary []string, sentence string) string {
	dl := len(dictionary)
	hash, lengths, lengthsMap := make(map[string]struct{}, dl), []int{}, make(map[int]struct{}, dl)
	for _, word := range dictionary {
		l := len(word)
		hash[word] = struct{}{}
		lengthsMap[l] = struct{}{}
	}
	for k, _ := range lengthsMap {
		lengths = append(lengths, k)
	}
	sort.Ints(lengths)
	resArr := strings.Split(sentence, " ")
loop:
	for i, word := range resArr {
		l := len(word)
		for _, v := range lengths {
			if v > l {
				break
			}
			if _, ok := hash[word[:v]]; ok {
				resArr[i] = word[:v]
				continue loop
			}
		}
		resArr[i] = word
	}
	return strings.Join(resArr, " ")
}

type FindMap struct {
	m map[uint8]*FindMap
}

func replaceWords0v1(dictionary []string, sentence string) string {
	m := map[uint8]*FindMap{}
	for _, word := range dictionary {
		var pointer *FindMap
		if v, ok := m[word[0]]; ok {
			if v.m == nil {
				continue
			}
			if 0 == len(word)-1 {
				v.m = nil
				continue
			}
			pointer = v
		} else {
			pointer = &FindMap{nil}
			m[word[0]] = pointer
		}
		for i := 1; i < len(word); i++ {
			if pointer.m == nil {
				pointer.m = map[uint8]*FindMap{}
			} else {
				if v, ok := pointer.m[word[i]]; ok {
					if v.m == nil {
						break
					}
					if i == len(word)-1 {
						v.m = nil
						break
					}
					pointer = v
					continue
				}
			}
			newPointer := &FindMap{nil}
			pointer.m[word[i]] = newPointer
			pointer = newPointer
		}
	}
	resArr := strings.Split(sentence, " ")
loop:
	for i, word := range resArr {
		if v, ok := m[word[0]]; ok {
			k := 0
			for j := 1; j < len(word) && v.m != nil; j++ {
				if nextV, nextOk := v.m[word[j]]; nextOk {
					k++
					if nextV.m == nil {
						break
					}
					v = nextV
				} else {
					resArr[i] = word
					continue loop
				}
			}
			resArr[i] = word[:k+1]
		} else {
			resArr[i] = word
		}
	}
	return strings.Join(resArr, " ")
}

func replaceWords0(dictionary []string, sentence string) string {

	m := map[uint8]*FindMap{}

	var dfs func(word string, index int, prev *FindMap)
	dfs = func(word string, index int, prev *FindMap) {
		if index == len(word) {
			prev.m = nil
			return
		}
		cur := FindMap{nil}
		if index == 0 {
			if v, ok := m[word[index]]; ok {
				if v.m != nil {
					dfs(word, index+1, v)
				}
			} else {
				m[word[index]] = &cur
				dfs(word, index+1, &cur)
			}
		} else {
			if prev.m == nil {
				prev.m = map[uint8]*FindMap{word[index]: &cur}
				dfs(word, index+1, &cur)
			} else if v, ok := prev.m[word[index]]; ok {
				if v.m != nil {
					dfs(word, index+1, v)
				}
			} else {
				prev.m[word[index]] = &cur
				dfs(word, index+1, &cur)
			}
		}
	}

	for _, word := range dictionary {
		dfs(word, 0, &FindMap{})
	}
	resArr := strings.Split(sentence, " ")
loop:
	for i, word := range resArr {
		if v, ok := m[word[0]]; ok {
			k := 0
			for j := 1; j < len(word) && v.m != nil; j++ {
				if nextV, nextOk := v.m[word[j]]; nextOk {
					k++
					if nextV.m == nil {
						break
					}
					v = nextV
				} else {
					resArr[i] = word
					continue loop
				}
			}
			resArr[i] = word[:k+1]
		} else {
			resArr[i] = word
		}
	}
	return strings.Join(resArr, " ")
}
