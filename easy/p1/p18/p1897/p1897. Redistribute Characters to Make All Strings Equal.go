package p1897

func makeEqual(words []string) bool {
	m := map[int32]int{}
	for _, word := range words {
		for _, c := range word {
			m[c]++
		}
	}
	l := len(words)
	for _, v := range m {
		if v%l != 0 {
			return false
		}
	}
	return true
}
