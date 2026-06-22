package p1189

var m = map[byte]byte{'b': 0, 'a': 1, 'l': 2, 'o': 3, 'n': 4}

func maxNumberOfBalloons(text string) int {
	cnt := [5]int{}
	for i := 0; i < len(text); i++ {
		if v, exists := m[text[i]]; exists {
			cnt[v]++
		}
	}
	return min(cnt[0], cnt[1], cnt[2]/2, cnt[3]/2, cnt[4])
}
