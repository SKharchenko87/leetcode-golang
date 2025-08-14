package p0791

func customSortString(order string, s string) string {
	m := make(map[rune]int, 26)
	for _, o := range order {
		m[o] = 0
	}
	tmp := []rune{}
	for _, c := range s {
		if _, ok := m[c]; ok {
			m[c]++
		} else {
			tmp = append(tmp, c)
		}
	}
	res := []rune{}
	for _, k := range order {
		for i := 0; i < m[k]; i++ {
			res = append(res, k)
		}
	}
	res = append(res, tmp...)
	return string(res)
}
