package p2375

type Stack []byte

func (s *Stack) Push(val byte) {
	*s = append(*s, val)
}
func (s *Stack) Pop() byte {
	l := len(*s)
	res := (*s)[l-1]
	*s = (*s)[:l-1]
	return res
}
func (s Stack) Empty() bool {
	l := len(s)
	return l == 0
}

func smallestNumber(pattern string) string {
	l := len(pattern)
	res := make([]byte, 0, l)
	stack := Stack(make([]byte, 0, l+1))
	num := byte(1)
	for i := 0; i < l; i++ {
		stack.Push(num)
		num++
		if pattern[i] == 'I' {
			for !stack.Empty() {
				res = append(res, stack.Pop()+'0')
			}
		}
	}
	stack.Push(num)
	for !stack.Empty() {
		res = append(res, stack.Pop()+'0')
	}
	return string(res)
}

func smallestNumber0(pattern string) string {
	length := len(pattern)
	res := make([]byte, length+1)
	visited := [10]bool{}
	var dfs func(index int) bool
	dfs = func(index int) bool {
		if index == length+1 {
			return true
		}
		if pattern[index-1] == 'I' {
			for i := res[index-1] - '0' + 1; i < 10; i++ {
				if visited[i] {
					continue
				}
				visited[i] = true
				res[index] = i + '0'
				if dfs(index + 1) {
					return true
				}
				visited[i] = false
			}
		} else {
			for i := res[index-1] - '0' - 1; i > 0; i-- {
				if visited[i] {
					continue
				}
				visited[i] = true
				res[index] = i + '0'
				if dfs(index + 1) {
					return true
				}
				visited[i] = false
			}
		}
		return false
	}
	for i := 1; i < 10; i++ {
		visited[i] = true
		res[0] = byte(i) + '0'
		if dfs(1) {
			return string(res)
		}
		visited[i] = false
	}
	return string(res)
}
