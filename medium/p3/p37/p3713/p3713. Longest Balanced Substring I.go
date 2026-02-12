package p3713

import (
	"math"
)

func longestBalanced(s string) int {
	n := len(s)
	total := make([]int, 26)
	for i := 0; i < n; i++ {
		total[s[i]-'a']++
	}

	cur := make([]int, 26)
	res := 0
	for i := 0; i < n; i++ {
		if res > n-i {
			break
		}
		copy(cur, total)
		for j := n - 1; j >= 0; j-- {
			if res >= j-i+1 {
				break
			}
			if check(cur) {
				res = j - i + 1
				break
			}
			cur[s[j]-'a']--
		}
		total[s[i]-'a']--
	}

	return res
}

func check(cur []int) bool {
	i := 0
	for ; i < len(cur) && cur[i] == 0; i++ {

	}

	if i == len(cur) {
		return false
	}

	x := cur[i]
	for ; i < len(cur); i++ {
		if cur[i] > 0 && cur[i] != x {
			return false
		}
	}
	return true
}

func minMax(num []int) (int, int) {
	mx, mn := math.MinInt, math.MaxInt
	for i := 0; i < len(num); i++ {
		if num[i] > 0 {
			if num[i] > mx {
				mx = num[i]
			}
			if num[i] < mn {
				mn = num[i]
			}
		}
	}
	return mn, mx
}
