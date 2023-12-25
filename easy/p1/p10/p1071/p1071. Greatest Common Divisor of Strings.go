package main

import (
	"fmt"
	"sort"
	"strings"
)

func dividers(x int) []int {
	m := map[int]bool{1: true, x: true}
	if x > 2 && x%2 == 0 {
		m[2] = true
		m[x/2] = true
	}
	if x > 3 && x%3 == 0 {
		m[3] = true
		m[x/3] = true
	}
	if x > 4 && x%4 == 0 {
		m[4] = true
		m[x/4] = true
	}
	for i := 5; i < x/5; i++ {
		if x%i == 0 {
			m[i] = true
			m[x/i] = true
		}
	}
	res := make([]int, len(m))
	i := 0
	for k, _ := range m {
		res[i] = k
		i++
	}
	sort.Slice(res, func(i, j int) bool { return res[i] > res[j] })
	return res
}

func gcdOfStrings(str1 string, str2 string) string {
	l2 := len(str2)
	m := dividers(l2)
	str := ""
	for _, v := range m {
		x := str2[:v]
		str = strings.Repeat(x, l2/v)
		if str == str2 {
			cnt := strings.Count(str1, x)
			if cnt != 0 && len(str1)%cnt == 0 {
				if str1 == strings.Repeat(x, cnt) {
					return x
				}
			}
		}
	}
	return ""
}

func main() {
	//str1, str2 := "ABABAB", "AB"
	str1, str2 := "ABABABAB", "ABAB"
	//str1, str2 := "ABCABC", "ABC"
	//str1, str2 := "ABABAB", "ABAB"
	//str1, str2 := "LEET", "CODE"
	//str1, str2 := "ABCDEF", "ABC"
	fmt.Println(gcdOfStrings(str1, str2))
}
