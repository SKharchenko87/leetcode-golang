package p1190

func reverseParentheses(s string) string {
	res := []byte(s)
	openParentheses := make([]int, 0)
	for i := 0; i < len(res); i++ {
		if res[i] == ')' {
			start := openParentheses[len(openParentheses)-1]
			openParentheses = openParentheses[:len(openParentheses)-1]
			reverseSubSlice(res, start, i)
			res = append(res[:i], res[i+1:]...)
			res = append(res[:start], res[start+1:]...)
			i -= 2
		} else if res[i] == '(' {
			openParentheses = append(openParentheses, i)
		}
	}
	return string(res)
}

type stack []interface{}

func (s *stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *stack) Pop() interface{} {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}

func reverseParentheses2(s string) string {
	res := []byte(s)
	openParentheses := stack{}
	for i := 0; i < len(res); i++ {
		if res[i] == ')' {
			start := openParentheses.Pop().(int)
			reverseSubSlice(res, start, i)
			res = append(res[:i], res[i+1:]...)
			res = append(res[:start], res[start+1:]...)
			i -= 2
		} else if res[i] == '(' {
			openParentheses.Push(i)
		}
	}
	return string(res)
}

func reverseSubSlice[T any](slice []T, start, end int) {
	for j := 0; j < (end-start)/2; j++ {
		slice[j+start+1], slice[end-1-j] = slice[end-1-j], slice[j+start+1]
	}
}

// Wormhole
func reverseParentheses1(s string) string {
	wormhole := map[int]int{}
	openParentheses := stack{}
	for i, c := range s {
		if c == '(' {
			openParentheses.Push(i)
		} else if c == ')' {
			v := openParentheses.Pop().(int)
			wormhole[v] = i
			wormhole[i] = v
		}
	}
	direction := 1
	curIndex := 0
	res := []byte{}
	for i := 0; i < len(s); i++ {
		if v, ok := wormhole[curIndex]; ok {
			curIndex = v
			direction = -direction
		} else {
			res = append(res, s[curIndex])
		}
		curIndex += direction
	}
	return string(res)
}

func reverseParentheses0(s string) string {
	res := []byte(s)
	openParentheses := make([]int, 0)
	for i := 0; i < len(res); i++ {
		if res[i] == ')' {
			start := openParentheses[len(openParentheses)-1]
			for j := 0; j < (i-start)/2; j++ {
				res[j+start+1], res[i-1-j] = res[i-1-j], res[j+start+1]
			}
			openParentheses = openParentheses[:len(openParentheses)-1]
			res = append(res[:i], res[i+1:]...)
			res = append(res[:start], res[start+1:]...)
			i -= 2
		} else if res[i] == '(' {
			openParentheses = append(openParentheses, i)
		}
	}
	return string(res)
}
