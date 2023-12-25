package main

import "fmt"

func maxVowels(s string, k int) int {
	counter := 0
	max := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u':
			counter++
		}
		if i >= k {
			switch s[i-k] {
			case 'a', 'e', 'i', 'o', 'u':
				counter--
			}
		}
		if counter > max {
			max = counter
		}
	}
	return max
}

func main() {
	//s, k := "leetcode", 3
	//s, k := "abciiidef", 3
	s, k := "weallloveyou", 7
	fmt.Println(maxVowels(s, k))
}
