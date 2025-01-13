package p3223

func minimumLength6(s string) int {
	m := make(map[byte]bool, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]] = !m[s[i]]
	}
	res := 0
	for _, v := range m {
		if v {
			res++
		} else {
			res += 2
		}
	}
	return res
}

func minimumLength(s string) int {
	alphabet := [26]bool{}
	visited := [26]bool{}
	for i := 0; i < len(s); i++ {
		letterIndex := s[i] - 'a'
		alphabet[letterIndex] = !alphabet[letterIndex]
		if !visited[letterIndex] {
			visited[letterIndex] = true
		}
	}
	res := 0
	for i := 0; i < 26; i++ {
		if visited[i] {
			if alphabet[i] {
				res++
			} else {
				res += 2
			}
		}
	}
	return res
}

func minimumLength5(s string) int {
	alphabet := [26]bool{}
	visited := [26]bool{}
	for i := 0; i < len(s); i++ {
		letterIndex := s[i] - 'a'
		alphabet[letterIndex] = !alphabet[letterIndex]
		visited[letterIndex] = true
	}
	res := 0
	for i := 0; i < 26; i++ {
		if visited[i] {
			if alphabet[i] {
				res++
			} else {
				res += 2
			}
		}
	}
	return res
}

func minimumLength4(s string) int {
	alphabet := [26]bool{}
	visited := [26]bool{}
	res := 0
	for i := 0; i < len(s); i++ {
		letterIndex := s[i] - 'a'
		if !visited[letterIndex] {
			visited[letterIndex] = true
			res++
		} else if alphabet[letterIndex] {
			res++
		} else {
			res--
		}
		alphabet[letterIndex] = !alphabet[letterIndex]
	}
	return res
}

func minimumLength3(s string) int {
	alphabet := [26]int{}
	for i := 0; i < len(s); i++ {
		if alphabet[s[i]-'a'] == 2 {
			alphabet[s[i]-'a']--
		} else {
			alphabet[s[i]-'a']++
		}
	}
	res := 0
	for i := 0; i < 26; i++ {
		res += alphabet[i]
	}
	return res
}

func minimumLength2(s string) int {
	alphabet := [26]byte{}
	res := 0
	for i := 0; i < len(s); i++ {
		letterIndex := s[i] - 'a'
		if alphabet[letterIndex] == 0 {
			alphabet[letterIndex] = 1
			res++
		} else if alphabet[letterIndex]&1 == 0 {
			alphabet[letterIndex] = 1
			res--
		} else {
			alphabet[letterIndex] = 2
			res++
		}
	}
	return res
}

func minimumLength1(s string) int {
	alphabet := [26]byte{}
	for i := 0; i < len(s); i++ {
		alphabet[s[i]-'a'] = (alphabet[s[i]-'a']+2)%2 + 1
	}
	res := 0
	for _, b := range alphabet {
		if b != 0 {
			if b&1 == 1 {
				res++
			} else {
				res += 2
			}
		}
	}
	return res
}

func minimumLength0(s string) int {
	alphabet := [26]bool{}
	m := make(map[byte]struct{}, 26)
	for i := 0; i < len(s); i++ {
		alphabet[s[i]-'a'] = !alphabet[s[i]-'a']
		m[s[i]-'a'] = struct{}{}
	}
	res := 0
	for k := range m {
		if alphabet[k] {
			res++
		} else {
			res += 2
		}
	}
	return res
}
