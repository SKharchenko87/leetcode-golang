package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	len1 := len(word1)
	len2 := len(word2)
	str := make([]uint8, len1+len2)
	l := min(len1, len2)
	x := max(len1, len2)
	for i := 0; i < l; i++ {
		str[2*i] = word1[i]
		str[2*i+1] = word2[i]
	}
	for i := l; i < x; i++ {
		var tmp uint8
		if l == len1 {
			tmp = word2[i]
		} else {
			tmp = word1[i]
		}
		str[l+i] = tmp
	}
	return string(str)
}

func main() {
	word1, word2 := "abcd", "pq"
	fmt.Println(mergeAlternately(word1, word2))
}
