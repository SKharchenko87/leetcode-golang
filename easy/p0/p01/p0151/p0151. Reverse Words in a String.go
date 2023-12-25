package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	ss := strings.Fields(s)
	sb := strings.Builder{}
	for i := len(ss) - 1; i >= 0; i-- {
		if sb.Len() > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(ss[i])
	}
	return sb.String()
}

func main() {
	fmt.Println("-=", reverseWords("the sky is blue"), "=-")
	fmt.Println("-=", reverseWords("  hello world  "), "=-")
	fmt.Println("-=", reverseWords("a good   example"), "=-")
}
