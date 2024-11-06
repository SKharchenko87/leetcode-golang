package p2914

func minChanges(s string) (res int) {
	for i := 0; i < len(s)-1; i += 2 {
		if s[i] != s[i+1] {
			res++
		}
	}
	return res
}

func minChanges0(s string) int {
	l := len(s)
	res := 0
	for i := 0; i < l-1; i += 2 {
		if s[i] != s[i+1] {
			res++
		}
	}
	return res
}
