package p2486

func appendCharacters(s string, t string) int {
	sIndex, ls := 0, len(s)
	tIndex, lt := 0, len(t)
	for sIndex < ls && tIndex < lt {
		if t[tIndex] == s[sIndex] {
			tIndex++
		}
		sIndex++
	}
	return lt - tIndex
}
