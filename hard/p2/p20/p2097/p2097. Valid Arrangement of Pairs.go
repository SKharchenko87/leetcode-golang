package p2097

type Node int

func validArrangement(pairs [][]int) [][]int {
	l := len(pairs)
	p := map[int]int{}
	relations := map[int][]int{}
	for i := 0; i < l; i++ {
		start, end := pairs[i][0], pairs[i][1]
		if _, ok := relations[start]; ok {
			relations[start] = append(relations[start], end)
		} else {
			relations[start] = []int{end}
		}

		p[start]++
		if p[start] == 0 {
			delete(p, start)
		}
		p[end]--
		if p[end] == 0 {
			delete(p, end)
		}
	}

	start := pairs[0][0]
	for pairIndex, cnt := range p {
		if cnt > 0 {
			start = pairIndex
			break
		}
	}

	stack := []int{start}
	result := make([][]int, l+2)
	indexResult := l
	result[indexResult+1] = []int{0, 0}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		n := len(relations[top])
		if n == 0 {
			result[indexResult+1][0] = int(top)
			result[indexResult] = []int{0, int(top)}
			indexResult--
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, relations[top][n-1])
			relations[top] = relations[top][:n-1]
		}
	}
	return result[1 : l+1]
}

func validArrangement0(pairs [][]int) [][]int {
	l := len(pairs)
	p := map[Node]int{}
	relations := map[Node][]Node{}
	for i := 0; i < l; i++ {
		start, end := Node(pairs[i][0]), Node(pairs[i][1])
		if _, ok := relations[start]; ok {
			relations[start] = append(relations[start], end)
		} else {
			relations[start] = []Node{end}
		}

		p[start]++
		if p[start] == 0 {
			delete(p, start)
		}
		p[end]--
		if p[end] == 0 {
			delete(p, end)
		}
	}

	start := Node(pairs[0][0])
	for pairIndex, cnt := range p {
		if cnt > 0 {
			start = pairIndex
			break
		}
	}

	path := []Node{}
	stack := []Node{start}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		n := len(relations[top])
		if n == 0 {
			path = append(path, top)
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, relations[top][n-1])
			relations[top] = relations[top][:n-1]
		}
	}

	m := len(path)
	result := make([][]int, l)
	for i := m - 1; i > 0; i-- {
		result[m-1-i] = []int{int(path[i]), int(path[i-1])}
	}
	return result
}
