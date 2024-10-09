package p1963

func minSwaps(s string) int {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	if len(stack) == 0 {
		return 0
	}
	return (len(stack)/2-1)/2 + 1
}

func minSwaps1(s string) int {
	stackSize := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			stackSize++
		} else if stackSize > 0 {
			stackSize--
		}
	}
	return (stackSize + 1) / 2
}

func minSwaps0(s string) int {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	if len(stack) == 0 {
		return 0
	}
	return (len(stack)/2-1)/2 + 1
}
