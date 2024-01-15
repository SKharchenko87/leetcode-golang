package p2

import (
	"fmt"
	"strings"
)

func getIndexes(s, a string) []int {
	res := []int{}
	i := strings.Index(s, a)
	x := 0
	for i != -1 {
		res = append(res, x+i)
		x = x + i + 1
		s = s[i+1:]
		i = strings.Index(s, a)
	}
	return res
}

func beautifulIndices(s string, a string, b string, k int) []int {
	sai := getIndexes(s, a)
	sbi := getIndexes(s, b)
	fmt.Println(sai)
	fmt.Println(sbi)
	res := []int{}
	for i := 0; i < len(sai); i++ {
		for j := 0; j < len(sbi); j++ {
			if abs(sbi[j]-sai[i]) <= k {
				res = append(res, sai[i])
				break
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
