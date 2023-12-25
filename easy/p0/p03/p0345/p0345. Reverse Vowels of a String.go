package main

import "fmt"

func reverseVowels(s string) string {
	m := map[byte]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1}
	x := []byte(s)
	l := len(x)
	i, j := 0, l-1
	for ; i < l; i++ {
		if _, ok := m[x[i]]; ok {
			for ; j >= i; j-- {
				if _, ok := m[x[j]]; ok {
					x[i], x[j] = x[j], x[i]
					j--
					break
				}
			}
		}
		if i >= j {
			break
		}
	}
	return string(x)
}

func main() {
	//fmt.Println(reverseVowels("leetcode"))
	//fmt.Println(reverseVowels("ai"))
	//fmt.Println(reverseVowels("aA"))
	fmt.Println(reverseVowels("race car"))
}
