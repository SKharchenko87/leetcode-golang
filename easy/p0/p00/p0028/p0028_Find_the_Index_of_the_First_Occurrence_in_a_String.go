package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	//return strings.Index(haystack, needle)
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			return -1
		}
		if haystack[i] == needle[0] {
			j := 1
			for j < len(needle) {
				if haystack[i+j] != needle[j] {
					break
				}
				j++
			}
			if j == len(needle) {
				return i
			}
		}
	}
	return -1
}

func main() {
	haystack := "aaa"
	needle := "aaaa"

	//haystack := "mississippi"
	//needle := "issipi"

	haystack = "assadbutsad"
	needle = "sad"

	fmt.Println(strStr(haystack, needle))
}
