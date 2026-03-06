package p1689

import (
	"strings"
)

func minPartitions(n string) int {
	for i := '9'; i > '1'; i-- {
		if strings.IndexRune(n, i) >= 0 {
			return int(i - '0')
		}
	}
	return 1
}

func minPartitions0(n string) int {
	res := 0
	for _, c := range n {
		res = max(res, int(c-'0'))
		if res == 9 {
			break
		}
	}
	return res
}
