package p1002

//ToDo banchmark

func commonChars(words []string) []string {
	common := [26]int{}
	cur := [26]int{}
	var min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	var opt = func(word *string) {
		cur = [26]int{}
		for _, ch := range *word {
			if common[ch-'a'] != 0 {
				cur[ch-'a']++
			}
		}
	}
	for _, ch := range words[0] {
		common[ch-'a']++
	}
	l := len(words)
	var j rune
	for i := 1; i < l-1; i++ {
		cur = [26]int{}
		opt(&words[i])
		for j = 0; j < 26; j++ {
			common[j] = min(common[j], cur[j])
		}
	}
	opt(&words[l-1])
	res := []string{}
	for j = 0; j < 26; j++ {
		for k := 0; k < min(common[j], cur[j]); k++ {
			res = append(res, string('a'+j))
		}
	}
	return res
}

func commonChars3(words []string) []string {
	l := len(words)
	common := [26]int{}
	cur := [26]int{}
	res := []string{}
	for i := 0; i < l; i++ {
		cur = [26]int{}
		for _, ch := range words[i] {
			if i == 0 || common[ch-'a'] != 0 {
				cur[ch-'a']++
			}
		}
		if i != 0 {
			var j rune
			for ; j < 26; j++ {
				common[j] = min(common[j], cur[j])
				if i == l-1 {
					for k := 0; k < common[j]; k++ {
						res = append(res, string('a'+j))
					}
				}
			}
		} else {
			common = cur
		}
	}
	return res
}

func commonChars2(words []string) []string {
	l := len(words)
	common := [26]int{}
	cur := [26]int{}
	res := []string{}
	for _, ch := range words[0] {
		common[ch-'a']++
	}
	for i := 1; i < l; i++ {
		cur = [26]int{}
		for _, ch := range words[i] {
			if common[ch-'a'] != 0 {
				cur[ch-'a']++
			}
		}
		var j rune
		for ; j < 26; j++ {
			common[j] = min(common[j], cur[j])
			if i == l-1 {
				for k := 0; k < common[j]; k++ {
					res = append(res, string('a'+j))
				}
			}
		}
	}
	return res
}

func commonChars1(words []string) []string {
	l := len(words)
	arr := make([][26]int, l)
	for i, word := range words {
		for _, ch := range word {
			arr[i][ch-'a']++
		}
	}
	res := []string{}
	var i rune
loop:
	for ; i < 26; i++ {
		minCh := 101
		for j := 0; j < l; j++ {
			if arr[j][i] == 0 {
				minCh = 0
				continue loop
			} else {
				minCh = min(minCh, arr[j][i])
			}
		}
		for k := 0; k < minCh; k++ {
			res = append(res, string('a'+i))
		}
	}
	return res
}

func commonChars0(words []string) []string {
	m := map[rune]int{'a': 100, 'b': 100, 'c': 100, 'd': 100, 'e': 100, 'f': 100, 'g': 100, 'h': 100, 'i': 100, 'j': 100, 'k': 100, 'l': 100, 'm': 100, 'n': 100, 'o': 100, 'p': 100, 'q': 100, 'r': 100, 's': 100, 't': 100, 'u': 100, 'v': 100, 'w': 100, 'x': 100, 'y': 100, 'z': 100}
	l := len(words)
	for i := 0; i < l-1; i++ {
		tmp := make(map[rune]int, 26)
		for _, c := range words[i] {
			tmp[c]++
		}
		minMap(m, tmp)
	}
	res := []string{}
	tmp := make(map[rune]int, 26)
	for _, c := range words[l-1] {
		tmp[c]++
		if v, ok := m[c]; ok && v >= tmp[c] {
			res = append(res, string(c))
		}
	}
	return res
}

func minMap(m1, m2 map[rune]int) {
	for k, v1 := range m1 {
		if v2, ok := m2[k]; ok {
			m1[k] = min(v1, v2)
		} else {
			delete(m1, k)
		}
	}
}
