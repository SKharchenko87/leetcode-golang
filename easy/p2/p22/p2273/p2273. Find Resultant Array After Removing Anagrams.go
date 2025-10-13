package p2273

import "sort"

func removeAnagrams(words []string) []string {
	l := len(words)
	ans := make([]string, 0, l)
	prevA, prevB, prevC := digit(words[0])
	ans = append(ans, words[0])
	for i := 1; i < l; i++ {
		word := words[i]
		currA, currB, currC := digit(word)
		if !(currA == prevA && currB == prevB && currC == prevC) {
			ans = append(ans, word)
		}
		prevA, prevB, prevC = currA, currB, currC
	}
	return ans
}

func digit(s string) (ax int, bx int, cx int) {
	pow := [9]int{1, 10, 100, 1000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000}
	for _, c := range s {
		x := c - 'a'
		p := x % 9
		m := x / 9
		switch m {
		case 0:
			ax += pow[p]
		case 1:
			bx += pow[p]
		case 2:
			cx += pow[p]
		}
	}
	return
}

func removeAnagrams0(words []string) []string {
	l := len(words)
	ans := make([]string, 0, l)
	prev := sortString(words[0])
	ans = append(ans, words[0])
	for i := 1; i < l; i++ {
		word := words[i]
		curr := sortString(word)
		if curr != prev {
			ans = append(ans, word)
		}
		prev = curr
	}
	return ans
}

func sortString(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
