package p0948

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	left, right, maxTotal, total := 0, len(tokens)-1, 0, 0
	for left <= right {
		if tokens[left] <= power {
			power, total, left = power-tokens[left], total+1, left+1
			if maxTotal < total {
				maxTotal = total
			}
		} else if total == 0 {
			break
		} else {
			power, right, total = power+tokens[right], right-1, total-1
		}
	}
	return maxTotal
}
