package p0796

import "strings"

func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}

func rotateString1(s string, goal string) bool {
	ls, lg := len(s), len(goal)
	if ls == lg {
		for i := 0; i < ls; i++ {
			if s[:i] == goal[lg-i:] && s[i:] == goal[:lg-i] {
				return true
			}
		}
	}
	return false
}

func rotateString0(s string, goal string) bool {
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
