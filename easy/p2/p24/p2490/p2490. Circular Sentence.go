package p2490

func isCircularSentence(sentence string) bool {
	l := len(sentence)
	if sentence[0] != sentence[l-1] {
		return false
	}
	for i := 1; i < l-1; i++ {
		if sentence[i] == ' ' {
			if sentence[i-1] != sentence[i+1] {
				return false
			}
		}
	}
	return true
}
