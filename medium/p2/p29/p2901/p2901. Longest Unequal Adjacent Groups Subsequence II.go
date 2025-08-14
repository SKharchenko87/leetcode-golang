package p2901

import "slices"

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	l := len(words)
	dp := make([]int, l)
	prev := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = 1
		prev[i] = -1
	}
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if HammingDistance(words[i], words[j]) != -1 && groups[i] != groups[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
	}
	bestWordIndex := -1
	best := -1
	for i := 0; i < l; i++ {
		if best < dp[i] {
			best = dp[i]
			bestWordIndex = i
		}
	}

	res := make([]string, 0, best+1)
	index := bestWordIndex
	for index != -1 {
		res = append(res, words[index])
		index = prev[index]
	}
	slices.Reverse(res)
	return res
}

func HammingDistance(a, b string) int {
	position := -1
	if len(a) != len(b) {
		return -1
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if position != -1 {
				return -1
			}
			position = i
		}
	}
	return position
}
