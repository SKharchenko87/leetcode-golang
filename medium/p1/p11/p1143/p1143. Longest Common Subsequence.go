package p1143

import "strings"

func toMapChar(text string) map[int32]bool {
	res := make(map[int32]bool, 26)
	for _, c := range text {
		res[c] = true
	}
	return res
}

func longestCommonSubsequence(text1 string, text2 string) int {
	//l1 := len(text1)
	//l2 := len(text2)
	m1 := toMapChar(text1)
	m2 := toMapChar(text2)

	sb1 := strings.Builder{}
	for _, c := range text1 {
		if m2[c] {
			sb1.WriteByte(byte(c))
		}
	}
	sb2 := strings.Builder{}
	for _, c := range text2 {
		if m1[c] {
			sb2.WriteByte(byte(c))
		}
	}
	
	return 0
}
