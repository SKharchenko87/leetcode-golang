package p2696

func minLength(s string) int {
	stack := []byte{s[0]}
	for i := 1; i < len(s); i++ {
		l := len(stack)
		if s[i] == 'B' || s[i] == 'D' {
			if l > 0 && stack[l-1] == s[i]-1 {
				stack = stack[:l-1]
				continue
			}
		}
		stack = append(stack, s[i])
	}
	return len(stack)
}

func minLength1(s string) int {
	index := -1
	stack := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == 'B' || s[i] == 'D' {
			if index >= 0 && s[stack[index]] == s[i]-1 {
				index--
				continue
			}
		}
		index++
		stack[index] = byte(i)
	}
	return index + 1
}

type stack []int

func (s *stack) Push(v int) {
	*s = append(*s, v)
}
func (s *stack) Pop() int {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}

func (s stack) Empty() bool {
	return len(s) == 0
}

func (s stack) Len() int {
	return len(s)
}

func (s stack) Peek() int {
	if len(s) == 0 {
		return -1
	}
	l := len(s)
	return s[l-1]
}

func minLength0(s string) int {
	var tmp stack

	for i := 0; i < len(s); i++ {

		c := s[i]
		if c == 'B' || c == 'D' {
			index := tmp.Peek()
			if index >= 0 && s[index] == c-1 {
				tmp.Pop()
				continue
			}
		}
		tmp.Push(i)
	}
	return tmp.Len()
}
