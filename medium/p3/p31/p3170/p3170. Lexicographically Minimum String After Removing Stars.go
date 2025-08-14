package p3170

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}

func (s Stack) Empty() bool {
	return len(s) == 0
}

func (s Stack) LastVal() int {
	return s[len(s)-1]
}

func clearStars(s string) string {
	l := len(s)
	alphabet := [26]Stack{}
	cntStars := 0
	for i := 0; i < l; i++ {
		if s[i] == '*' {
			cntStars++
			for j := 0; j < 26; j++ {
				if !alphabet[j].Empty() {
					alphabet[j].Pop()
					break
				}
			}
		} else {
			ch := s[i] - 'a'
			alphabet[ch].Push(i)
		}
	}

	resBytes := make([]byte, l-cntStars)
	index := l - cntStars - 1
	for i := l - 1; i >= 0; i-- {
		if s[i] == '*' {
		} else {
			ch := s[i] - 'a'
			if !alphabet[ch].Empty() {
				v := alphabet[ch].Pop()
				if v == i {
					resBytes[index] = s[i]
					index--
				} else {
					alphabet[ch].Push(v)
				}
			}
		}
	}
	return string(resBytes[index+1:])
}

func clearStars0(s string) string {
	l := len(s)
	alphabet := [26]Stack{}
	cntStars := 0
	for i := 0; i < l; i++ {
		if s[i] == '*' {
			cntStars++
			for j := 0; j < 26; j++ {
				if !alphabet[j].Empty() {
					alphabet[j].Pop()
					break
				}
			}
		} else {
			ch := s[i] - 'a'
			alphabet[ch].Push(i)
		}
	}

	resBytes := make([]byte, l-cntStars)
	index := l - cntStars - 1
	for i := l - 1; i >= 0; i-- {
		if s[i] == '*' {
		} else {
			ch := s[i] - 'a'
			if !alphabet[ch].Empty() {
				v := alphabet[ch].LastVal()
				if v == i {
					resBytes[index] = s[i]
					index--
					alphabet[ch].Pop()
				}
			}
		}
	}
	return string(resBytes[index+1:])
}
