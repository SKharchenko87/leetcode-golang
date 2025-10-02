package p2707

import (
	"slices"
	"sort"
	"strings"
)

type TrieNode struct {
	children map[uint8]*TrieNode
	isEnd    bool
}

func (tr *TrieNode) buildTrie(dictionary []string) {
	for _, word := range dictionary {
		if tr.children[word[0]] == nil {
			tr.children[word[0]] = &TrieNode{make(map[uint8]*TrieNode), false}
		}
		var cur, tmp *TrieNode
		cur = tr.children[word[0]]
		for i := 1; i < len(word); i++ {
			if cur.children[word[i]] == nil {
				tmp = &TrieNode{make(map[uint8]*TrieNode), false}
				cur.children[word[i]] = tmp
			}
			cur = cur.children[word[i]]
		}
		cur.isEnd = true
	}
}

func (tr *TrieNode) searchWord(word string) bool {
	cur := tr
	for _, ch := range word {
		if v, ok := cur.children[uint8(ch)]; !ok {
			return false
		} else {
			cur = v
		}
	}
	return cur.isEnd
}

func minExtraChar2(s string, dictionary []string) int {
	tr := &TrieNode{make(map[uint8]*TrieNode), false}
	tr.buildTrie(dictionary)
	l := len(s)
	dp := make([]int, l+1)
	for i := l - 1; i >= 0; i-- {
		dp[i] = l - i
	}
	for start := l - 1; start >= 0; start-- {
		cur := tr
		depth := 0
		for _, ch := range s[start:] {
			if v, ok := cur.children[uint8(ch)]; !ok {
				break
			} else {
				depth++
				cur = v
				if cur.isEnd {
					dp[start] = min(dp[start+depth], dp[start])
				}
			}
		}
		dp[start] = min(dp[start+1]+1, dp[start])
	}
	return dp[0]
}

func minExtraChar(s string, dictionary []string) int {
	sort.Strings(dictionary) //KlogK
	l := len(s)
	dp := make([]int, l+1)
	for i := l - 1; i >= 0; i-- { //N
		dp[i] = l - i
	}
	for start := l - 1; start >= 0; start-- { //N
		b := l
		index, _ := slices.BinarySearch(dictionary, string(s[start])) // logK
		for i := index; i < len(dictionary); i++ {                    //K
			word := dictionary[i]
			if s[start:] < word {
				break
			}
			if strings.HasPrefix(s[start:], word) { //M
				if start+len(word) == l {
					b = 0
				} else {
					b = min(b, dp[start+len(word)])
				}
			}
		}
		dp[start] = min(dp[start+1]+1, b)
	}
	return dp[0]
}

func minExtraChar1(s string, dictionary []string) int {
	sort.Strings(dictionary)
	var rec func(start int) int
	l := len(s)
	dp := make(map[int]int, l)
	rec = func(start int) int {
		if v, ok := dp[start]; ok {
			return v
		}
		if start >= l {
			return 0
		}
		a := 1 + rec(start+1)
		b := l
		index, _ := slices.BinarySearch(dictionary, string(s[start]))
		for i := index; i < len(dictionary); i++ {
			word := dictionary[i]
			if s[start] < word[0] {
				break
			}
			if strings.HasPrefix(s[start:], word) {
				b = min(b, dp[start+len(word)])
			}
		}
		dp[start] = min(a, b)
		return dp[start]
	}
	return rec(0)
}

func minExtraChar0(s string, dictionary []string) int {
	var rec func(start int) int
	l := len(s)
	dp := make(map[int]int, l)
	rec = func(start int) int {
		if v, ok := dp[start]; ok {
			return v
		}
		if start >= l {
			return 0
		}
		a := 1 + rec(start+1)
		b := l
		for _, word := range dictionary {
			if strings.HasPrefix(s[start:], word) {
				b = min(b, dp[start+len(word)])
			}
		}
		dp[start] = min(a, b)
		return dp[start]
	}
	return rec(0)
}
