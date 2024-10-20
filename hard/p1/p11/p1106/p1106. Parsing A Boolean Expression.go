package p1106

type stack struct {
	arr []interface{}
}

func (s *stack) Push(v interface{}) {
	s.arr = append(s.arr, v)
}

func (s *stack) Pop() interface{} {
	l := len(s.arr)
	if l == 0 {
		return nil
	}
	v := s.arr[l-1]
	s.arr = s.arr[:l-1]
	return v
}

func (s *stack) Empty() bool {
	return len(s.arr) == 0
}

func (s *stack) Len() int {
	return len(s.arr)
}

type bools struct {
	f, t int
}

func parseBoolExpr(expression string) bool {
	st := stack{}
	for _, exp := range expression {
		if exp != ',' && exp != '(' {
			if exp == ')' {
				b := bools{}
				var x interface{}
			LOOP:
				for x = st.Pop(); x != nil; x = st.Pop() {
					switch x.(int32) {
					case 't':
						b.t++
					case 'f':
						b.f++
					default:
						break LOOP
					}
				}
				c := x.(int32)
				if c == '&' && b.f == 0 || c == '|' && b.t > 0 || c == '!' && b.f > 0 {
					st.Push('t')
				} else {
					st.Push('f')
				}
			} else {
				st.Push(exp)
			}
		}
	}
	return st.Pop().(int32) == 't'
}
