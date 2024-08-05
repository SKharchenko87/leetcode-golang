package p2678

func countSeniors(details []string) (cnt int) {
	for _, d := range details {
		p := d[11] - '0'
		if p > 6 || p == 6 && d[12]-'0' > 0 {
			cnt++
		}
	}
	return cnt
}
