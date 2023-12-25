package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs) && len(prefix) > 0; i++ {
		for len(prefix) > 0 && (strings.Index(strs[i], prefix) == -1 || strings.Index(strs[i], prefix) != 0) {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//str := "asfasdfasd"
	//str2 := str[:len(str)-1]
	//fmt.Println(str2)

}
