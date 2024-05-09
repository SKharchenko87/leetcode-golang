package p0506

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) (res []string) {
	l := len(score)
	res = make([]string, l)
	tmp := make([][2]int, l)
	for i, s := range score {
		tmp[i] = [2]int{i, s}
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i][1] > tmp[j][1]
	})
	for i, s := range tmp {
		switch i {
		case 0:
			res[s[0]] = "Gold Medal"
		case 1:
			res[s[0]] = "Silver Medal"
		case 2:
			res[s[0]] = "Bronze Medal"
		default:
			res[s[0]] = strconv.Itoa(i + 1)
		}
	}
	return
}
