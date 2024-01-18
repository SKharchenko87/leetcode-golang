package p0070

func climbStairs(n int) int {
	switch n {
	case 0, 1, 2, 3:
		return n
	}
	cur, prev := 3, 2
	for i := 3; i < n; i++ {
		cur, prev = cur+prev, cur
	}
	return cur
}
