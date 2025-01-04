package p1930

import (
	"math/bits"
	"slices"
)

func countPalindromicSubsequence(s string) int {
	l := len(s)
	left, right := make([]uint32, l), make([]uint32, l)
	left[0] = 1 << (s[0] - 'a')
	for i := 1; i < l; i++ {
		left[i] = left[i-1] | 1<<(s[i]-'a')
	}
	right[l-1] = 1 << (s[l-1] - 'a')
	for i := l - 2; i >= 1; i-- {
		right[i] = right[i+1] | 1<<(s[i]-'a')
	}
	arr := [26]uint32{}
	for i := 1; i < l-1; i++ {
		arr[s[i]-'a'] = arr[s[i]-'a'] | (left[i-1] & right[i+1])
	}

	res := 0
	for _, u := range arr {
		res += bits.OnesCount32(u)
	}
	return res
}

func countPalindromicSubsequence2(s string) int {
	m := make(map[byte][]int, 26)
	l := len(s)
	for i := 0; i < l; i++ {
		c := s[i] - 'a'
		m[c] = append(m[c], i)
	}
	var c0, c1 byte
	res := 0
	for c0 = 0; c0 < 26; c0++ {
		if len(m[c0]) > 0 {
			index0 := m[c0][0]
			index2 := m[c0][len(m[c0])-1]
			for c1 = 0; c1 < 26; c1++ {
				if len(m[c1]) > 0 {
					tmp, _ := slices.BinarySearch(m[c1], index0+1)
					if tmp < len(m[c1]) && m[c1][tmp] < index2 {
						res++
					}
				}
			}
		}
	}
	return res
}

func countPalindromicSubsequence0(s string) int {
	right := [26]int{}
	for i := 1; i < len(s); i++ {
		right[s[i]-'a']++
	}
	left := [26]bool{}
	left[s[0]-'a'] = true
	resMap := make(map[[2]byte]struct{}, 676)
	for i := 1; i < len(s)-1; i++ {
		c := s[i] - 'a'
		right[c]--
		for c1, b := range left {
			if b && right[c1] > 0 {
				if _, ok := resMap[[2]byte{c, byte(c1)}]; !ok {
					resMap[[2]byte{c, byte(c1)}] = struct{}{}
				}
			}
		}
		left[c] = true
	}
	return len(resMap)
}

func countPalindromicSubsequence1(s string) int {
	right := [26]int{}
	for i := 1; i < len(s); i++ {
		right[s[i]-'a']++
	}
	left := [26]bool{}
	left[s[0]-'a'] = true
	resMap := make(map[[2]byte]struct{}, 676)
	for i := 1; i < len(s)-1; i++ {
		c := s[i] - 'a'
		right[c]--
		for c1, b := range left {
			if b && right[c1] > 0 {
				resMap[[2]byte{c, byte(c1)}] = struct{}{}
			}
		}
		left[c] = true
	}
	return len(resMap)
}
