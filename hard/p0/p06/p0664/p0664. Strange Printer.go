package p0664

func strangePrinter(s string) int {
	s = removeDuplicates(s)
	memo := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		memo[i] = make([]int, len(s))
	}
	return minimumTurns(s, 0, len(s)-1, memo)
}

func removeDuplicates(s string) string {
	if len(s) <= 1 {
		return s
	}
	deduplicate := []byte{s[0]}
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			deduplicate = append(deduplicate, s[i])
		}
	}
	return string(deduplicate)
}

func minimumTurns(s string, start, end int, memo [][]int) int {
	if start > end {
		return 0
	}
	if memo[start][end] != 0 {
		return memo[start][end]
	}
	minTurns := 1 + minimumTurns(s, start+1, end, memo)
	for k := start + 1; k <= end; k++ {
		if s[k] == s[start] {
			turnsWithMatch := minimumTurns(s, start, k-1, memo) + minimumTurns(s, k+1, end, memo)
			minTurns = min(minTurns, turnsWithMatch)
		}
	}
	memo[start][end] = minTurns
	return minTurns
}
