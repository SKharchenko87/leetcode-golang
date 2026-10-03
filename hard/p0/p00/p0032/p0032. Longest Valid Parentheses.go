package p0032

func longestValidParentheses(s string) int {
	n := len(s)
	m := map[int][]int{}
	cur := 0
	m[0] = []int{-1}
	for i := 0; i < n; i++ {
		if s[i] == ')' {
			cur--
		} else {
			cur++
		}
		if v, ok := m[cur]; ok {
			m[cur] = append(v, i)
		} else {
			m[cur] = []int{i}
		}
	}
	res := 0
	for _, v := range m {
		for j := len(v) - 1; j > 0; j-- {
			for i := 0; i < len(v) && i < j && (v[j]-v[i]) > res; i++ {
				if isValid(s[v[i]+1 : v[j]+1]) {
					res = v[j] - v[i]
				}
			}
		}
	}
	return res
}

func isValid(s string) bool {
	cur := 0
	if s[0] == ')' || s[len(s)-1] == '(' {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			cur--
		} else {
			cur++
		}
		if cur < 0 {
			return false
		}
	}
	return cur == 0
}
