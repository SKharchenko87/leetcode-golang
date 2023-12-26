package p1155

import "fmt"

const mod = 1e9 + 7

func getLimitLevel(n int, k int) []int {
	maxLimit := make([]int, n)
	maxLimit[0] = k
	for level := 1; level < n; level++ {
		maxLimit[level] = maxLimit[level-1] + k
	}
	return maxLimit
}

func NumRollsToTarget(n int, k int, target int) int {
	return numRollsToTarget(n, k, target)
}

func numRollsToTarget(n int, k int, target int) int {
	m := map[int]int64{}
	for i := 1; i <= k; i++ {
		curSum := target - i
		if n-1 <= curSum && curSum <= (n-1)*k {
			m[curSum]++
		}
	}
	for level := n - 2; level >= 0; level-- {
		curMap := map[int]int64{}
		for i := 1; i <= k; i++ {
			for newTarget, v := range m {
				curSum := newTarget - i
				if level <= curSum && curSum <= (level+1)*k {
					if cm, ok := curMap[curSum]; ok {
						curMap[curSum] = (cm + v) % mod
					} else {
						curMap[curSum] = v % mod
					}
				}
			}
		}
		m = curMap
	}
	fmt.Println(m)
	if res, ok := m[0]; ok {
		return int(res % mod)
	}
	return 0
}
