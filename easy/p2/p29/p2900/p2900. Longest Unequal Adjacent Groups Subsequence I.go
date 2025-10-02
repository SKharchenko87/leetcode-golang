package p2900

func getLongestSubsequence(words []string, groups []int) []string {
	res := []string{words[0]}
	prev := groups[0]
	for i := 1; i < len(words); i++ {
		if prev != groups[i] {
			res = append(res, words[i])
			prev = groups[i]
		}
	}
	return res
}
