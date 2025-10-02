package p0796

func rotateString(s string, goal string) bool {
	l := len(s)
	if l != len(goal) {
		return false
	}
	for i := 0; i < l; i++ {
		if s[i:] == goal[:l-i] && s[:i] == goal[l-i:] {
			return true
		}
	}
	return false
}
