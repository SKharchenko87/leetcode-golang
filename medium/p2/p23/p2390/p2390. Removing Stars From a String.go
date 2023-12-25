package main

import "fmt"

func main() {
	s := "leet**cod****e"
	fmt.Println(removeStars(s))
}
func removeStars(s string) string {
	res := make([]uint8, len(s))
	counterStart := 0
	j := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '*' {
			counterStart++
		} else {
			if counterStart != 0 {
				counterStart--
			} else {
				res[j] = s[i]
				j--
			}
		}
	}
	return string(res[j+1:])
}
