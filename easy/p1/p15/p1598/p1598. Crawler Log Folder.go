package p1598

func minOperations(logs []string) int {
	res := 0
	for i := 0; i < len(logs); i++ {
		switch logs[i] {
		case "../":
			if res > 0 {
				res--
			}
		case "./":
		default:
			res++
		}
	}
	return res
}
