package p1347

func minSteps(s string, t string) int {
	res := 0
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	for _, v := range m {
		if v > 0 {
			res += v
		}
	}
	return res
}
