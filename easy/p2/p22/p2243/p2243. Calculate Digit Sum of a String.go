package p2243

import (
	"strconv"
	"strings"
)

func digitSum(s string, k int) string {
	if len(s) <= k {
		return s
	}
	sb := strings.Builder{}
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i] - '0')
		if i%k == k-1 {
			sb.WriteString(strconv.Itoa(sum))
			sum = 0
		}
	}
	if len(s)%k > 0 {
		sb.WriteString(strconv.Itoa(sum))
	}
	return digitSum(sb.String(), k)
}
