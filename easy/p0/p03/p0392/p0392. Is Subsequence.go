package main

import "fmt"

func isSubsequence(s string, t string) bool {
	i := 0
	ls := len(s)
	if ls == 0 {
		return true
	}
	for j := i; j < len(t); j++ {
		if s[i] == t[j] {
			i++
			if i == ls {
				return true
			}
		}
	}
	return false
}

func main() {
	s, t := "abc", "ahxgdc"
	fmt.Println(isSubsequence(s, t))
}
