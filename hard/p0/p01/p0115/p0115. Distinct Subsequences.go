package p0115

func numDistinct(s string, t string) int {
	l := len(t)
	dp := make([]int, l)
	m := [52][]int{}
	for i := 0; i < 52; i++ {
		m[i] = []int{}
	}
	for i := 0; i < l; i++ {
		v := t[i]
		if v <= 'Z' {
			m[v-'A'] = append(m[v-'A'], i)
		} else {
			m[v-'a'+26] = append(m[v-'a'+26], i)
		}
	}
	for i := 0; i < len(s); i++ {
		var v byte
		if s[i] <= 'Z' {
			v = s[i] - 'A'
		} else {
			v = s[i] - 'a' + 26
		}
		for j := len(m[v]) - 1; j >= 0; j-- {
			u := m[v][j]
			if u == 0 {
				dp[u]++
			} else {
				dp[u] += dp[u-1]
			}
		}
	}
	return dp[l-1]
}
