package p0150

import "strconv"

type stack[T interface{}] struct {
	arr   []T
	point int
}

func (s *stack[T]) pop() T {
	s.point--
	return s.arr[s.point]
}

func (s *stack[T]) peek() T {
	return s.arr[s.point-1]
}

func (s *stack[T]) push(x T) {
	if s.point == len(s.arr) {
		s.arr = append(s.arr, x)
	} else {
		s.arr[s.point] = x
	}
	s.point++
}

func evalRPN(tokens []string) int {
	st := stack[int]{arr: make([]int, 0), point: 0}
	for _, s := range tokens {
		switch s {
		case "+", "-", "*", "/":
			b := st.pop()
			a := st.pop()
			switch s {
			case "+":
				st.push(a + b)
			case "-":
				st.push(a - b)
			case "*":
				st.push(a * b)
			case "/":
				st.push(a / b)
			}
		default:
			x, _ := strconv.Atoi(s)
			st.push(x)
		}
	}
	return st.pop()
}
