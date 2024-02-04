package p0076

import "math"

// ToDo разобраться
func minWindow(s string, t string) string {

	result := ""
	minLength := math.MaxInt

	var have = make(map[byte]int)
	var need = make(map[byte]int)

	if t == "" || len(s) < len(t) {
		return result
	}

	for i := 0; i < len(t); i++ {
		need[t[i]] = 1 + need[t[i]]
		have[t[i]] = 0
	}
	haveCount := 0
	needCount := len(need)

	l, r := 0, 0
	for r < len(s) {
		if _, exists := have[s[r]]; exists {
			have[s[r]]++
			if have[s[r]] == need[s[r]] {
				haveCount++
			}
		}
		for haveCount == needCount {
			if minLength > (r - l) {
				result = s[l : r+1]
				minLength = r - l + 1
			}
			if _, exists := have[s[l]]; exists {
				have[s[l]]--
				if have[s[l]] < need[s[l]] {
					haveCount--
				}
			}
			l++
		}
		r++
	}
	return result
}
