package p4

import (
	"fmt"
	"strings"
)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

type tree struct {
	children map[uint8]*tree
	isEnd    bool
}

func minStartingIndex0(s string, pattern string) int {
	ls := len(s)
	lp := len(pattern)
LOOP:
	for i := 0; i < ls-lp+1; i++ {
		err := 1
		for j := 0; j < lp; j++ {
			if s[j+i] != pattern[j] {
				err--
				if err < 0 {
					continue LOOP
				}
			}
		}
		return i
	}
	return -1
}

func minStartingIndex1(s string, pattern string) int {
	lp := len(pattern)
	ls := len(s)
	if lp > ls {
		return -1
	}
	if lp == 1 {
		return 0
	}

	hp := make([]int, lp)
	hp[0] = int(pattern[0] - 'a')
	for i := 1; i < lp; i++ {
		hp[i] = hp[i-1] + int(pattern[i]-'a')
	}

	hash := make([]int, ls)
	hash[0] = int(s[0] - 'a')
	i := 1
	for ; i < lp; i++ {
		hash[i] = hash[i-1] + int(s[i]-'a')
	}

	if hash[i-1]-26 <= hp[lp-1] && hp[lp-1] <= hash[i-1]+26 {
		err := 1
		j := 0
		for ; j < lp; j++ {
			if s[j+i-lp+1-1] != pattern[j] {
				err--
				if err < 0 {
					break
				}
			}
		}
		if j == lp {
			return i - lp + 1 - 1
		}
	}

LOOP:
	for ; i < ls; i++ {
		hash[i] = hash[i-1] + int(s[i]-'a')
		if hash[i]-26-hash[i-lp] <= hp[lp-1] && hp[lp-1] <= hash[i]+26-hash[i-lp] {
			err := 1
			for j := 0; j < lp; j++ {
				if s[j+i-lp+1] != pattern[j] {
					err--
					if err < 0 {
						continue LOOP
					}
				}
			}
			return i - lp + 1
		}
	}
	return -1
}

func minStartingIndex(s string, pattern string) int {
	lp := len(pattern)
	ls := len(s)
	if lp > ls {
		return -1
	}
	if lp == 1 {
		return 0
	}
	p1 := pattern[:lp/2]
	arr:=strings.Split(s, p1)
	for i, s2 := range arr {
		if len(s2)<
	}
	p2 := pattern[lp/2:]
	fmt.Println(p1, p2)
	return -1
}
