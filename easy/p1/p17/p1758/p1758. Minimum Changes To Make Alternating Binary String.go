package p1758

func minOperations(s string) int {
	firstOne := 0
	firstZero := 0
	for i := 0; i < len(s); i++ {
		x := int(s[i] - '0')
		if i%2 == 0 {
			firstZero += x
			firstOne += 1 - x
		} else {
			firstZero += 1 - x
			firstOne += x
		}
	}
	return min(firstOne, firstZero)
}
