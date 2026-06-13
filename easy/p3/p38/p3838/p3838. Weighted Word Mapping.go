package p3838

func mapWordWeights(words []string, weights []int) string {
	l := len(words)
	res := make([]byte, l)
	var s int
	for i := 0; i < l; i++ {
		s = 0
		for j := 0; j < len(words[i]); j++ {
			s += weights[words[i][j]-'a']
		}
		res[i] = byte('z' - s%26)
	}
	return string(res)
}
