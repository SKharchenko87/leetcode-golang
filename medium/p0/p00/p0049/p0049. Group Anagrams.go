package p0049

// ToDo benchmark
type key [26]byte

func getKey(str string) (k key) {
	for _, c := range str {
		k[c-'a']++
	}
	return
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[key][]string)
	for _, v := range strs {
		k := getKey(v)
		m[k] = append(m[k], v)
	}
	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

//func groupAnagrams(strs []string) [][]string {
//	m := make(map[key][]string)
//	for _, v := range strs {
//		k := getKey(v)
//		m[k] = append(m[k], v)
//	}
//	result := make([][]string, 0, len(m))
//	for _, v := range m {
//		result = append(result, v)
//	}
//	return result
//}

//func groupAnagrams(strs []string) [][]string {
//	m := make(map[key]int)
//	res := [][]string{}
//	for _, str := range strs {
//		word_key := getKey(str)
//		if index, ok := m[word_key]; ok {
//			res[index] = append(res[index], str)
//		} else {
//			m[word_key] = len(res)
//			res = append(res, []string{str})
//		}
//	}
//	return res
//}

//func groupAnagrams(strs []string) [][]string {
//	m := make(map[key][]string)
//	for _, str := range strs {
//		word_key := getKey(str)
//		if _, ok := m[word_key]; ok {
//			m[word_key] = append(m[word_key], str)
//		} else {
//			m[word_key] = []string{str}
//		}
//	}
//	res := make([][]string, len(m))
//	i := 0
//	for _, v := range m {
//		res[i] = v
//		i++
//	}
//	return res
//}

//func groupAnagrams(strs []string) [][]string {
//	m := make(map[string][]string)
//
//	for _, str := range strs {
//		word := []byte(str)
//		sort.Slice(word, func(i, j int) bool {
//			return word[i] < word[j]
//		})
//		word_sort := string(word)
//		if _, ok := m[word_sort]; ok {
//			m[word_sort] = append(m[word_sort], str)
//		} else {
//			m[word_sort] = []string{str}
//		}
//	}
//	res := make([][]string, len(m))
//	i := 0
//	for _, v := range m {
//		res[i] = v
//		i++
//	}
//	return res
//}

//func groupAnagrams(strs []string) [][]string {
//	m := make(map[string]int)
//	res := [][]string{}
//	for _, str := range strs {
//		word := []byte(str)
//		sort.Slice(word, func(i, j int) bool {
//			return word[i] < word[j]
//		})
//		word_sort := string(word)
//		if index, ok := m[word_sort]; ok {
//			res[index] = append(res[index], str)
//		} else {
//			m[word_sort] = len(res)
//			res = append(res, []string{str})
//		}
//	}
//	return res
//}
