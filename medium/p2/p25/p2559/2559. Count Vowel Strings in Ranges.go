package p2559

var m = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

func vowelStrings(words []string, queries [][]int) []int {
	n := len(words)
	prefSum := make([]int, n)
	cnt := 0
	for i := 0; i < n; i++ {
		if m[words[i][0]] && m[words[i][len(words[i])-1]] {
			cnt++
		}
		prefSum[i] = cnt
	}
	q := len(queries)
	ans := make([]int, q)
	for i := 0; i < q; i++ {
		l, r := queries[i][0], queries[i][1]
		ans[i] = prefSum[r] - prefSum[l]
		if m[words[l][0]] && m[words[l][len(words[l])-1]] {
			ans[i]++
		}
	}
	return ans
}
