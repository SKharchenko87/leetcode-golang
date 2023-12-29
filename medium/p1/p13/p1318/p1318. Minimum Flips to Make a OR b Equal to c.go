package p1318

func minFlips(a int, b int, c int) int {
	if a+b == c || c-a+c-b == 0 {
		return 0
	}
	cnt := 0
	for i := 1 << 32; i > 0; i = i >> 1 {
		abit, bbit, cbit := a/i, b/i, c/i
		if cbit == 1 && abit == 0 && bbit == 0 {
			cnt++
		} else {
			if cbit == 0 && abit == 1 {
				cnt++
			}
			if cbit == 0 && bbit == 1 {
				cnt++
			}
		}
		a, b, c = a%i, b%i, c%i
	}
	return cnt
}
