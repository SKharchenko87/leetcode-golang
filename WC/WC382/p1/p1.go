package p1

import (
	"strings"
)

func countKeyChanges(s string) int {
	sl := strings.ToLower(s)
	count := 0
	for i := 1; i < len(sl); i++ {
		if sl[i] != sl[i-1] {
			count++
		}
	}
	return count
}
