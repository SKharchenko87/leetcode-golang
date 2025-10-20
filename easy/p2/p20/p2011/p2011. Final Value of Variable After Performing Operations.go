package p2011

func finalValueAfterOperations(operations []string) int {
	res := 0
	for _, operation := range operations {
		if operation[1] == '+' {
			res++
		} else {
			res--
		}
	}
	return res
}
