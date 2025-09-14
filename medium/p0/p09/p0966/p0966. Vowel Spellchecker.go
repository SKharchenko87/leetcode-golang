package p0966

import "strings"

func spellchecker(wordlist []string, queries []string) []string {
	l := len(queries)
	m1 := make(map[string]int, l)
	m2 := make(map[string]int, l)
	m3 := make(map[string]int, l)
	r := strings.NewReplacer("a", "*", "e", "*", "i", "*", "o", "*", "u", "*")
	for i, s := range wordlist {
		if _, ok := m1[s]; !ok {
			m1[s] = i
		}
		sl := strings.ToLower(s)
		if _, ok := m2[sl]; !ok {
			m2[sl] = i
		}
		sc := r.Replace(sl)
		if _, ok := m3[sc]; !ok {
			m3[sc] = i
		}
	}

	result := make([]string, l)
	for i := 0; i < l; i++ {
		if v, ok := m1[queries[i]]; ok {
			result[i] = wordlist[v]
		} else if v, ok := m2[strings.ToLower(queries[i])]; ok {
			result[i] = wordlist[v]
		} else if v, ok := m3[r.Replace(strings.ToLower(queries[i]))]; ok {
			result[i] = wordlist[v]
		}
	}
	return result
}
