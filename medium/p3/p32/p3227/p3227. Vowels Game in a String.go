package p3227

func doesAliceWin(s string) bool {
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			return true
		}
	}
	return false

}

func doesAliceWin0(s string) bool {
	cntVowels := 0
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			cntVowels++
		}
	}
	if cntVowels == 0 {
		return false
	}
	return true
}

//   leetcode
// A       de - leetco
// B        e - d
// A          - e

//   leetcod
// A         - leetcod
