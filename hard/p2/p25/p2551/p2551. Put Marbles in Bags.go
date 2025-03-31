package p2551

import "sort"

func putMarbles(weights []int, k int) int64 {
	lastIdx := len(weights) - 1
	for i := 0; i < lastIdx; i++ {
		weights[i] = weights[i] + weights[i+1]
	}
	weights = weights[:lastIdx]

	sort.Slice(weights, func(i, j int) bool {
		return weights[i] > weights[j]
	})

	ans := 0
	for i := 0; i < k-1; i++ {
		ans += weights[i] - weights[lastIdx-i-1]
	}

	return int64(ans)
}
