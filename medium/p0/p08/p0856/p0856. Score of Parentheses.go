package p0856

func scoreOfParentheses(s string) int {
	n := len(s)
	stack := make([]int, 0, n/2)
	res := 0
	cur := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, cur)
			cur = 0
		} else {
			if cur == 0 {
				cur++
			} else {
				cur *= 2
			}
			cur += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res += cur
			cur = 0
		}
	}
	res += cur
	return res
}
