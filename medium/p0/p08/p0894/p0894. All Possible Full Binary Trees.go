package p0894

import (
	. "leetcode/stucture"
)

var factorial = [20]int{}

func init() {
	factorial[0] = 1
	for i := 1; i < 20; i++ {
		factorial[i] = factorial[i-1] * i
	}
}

func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}
	dp := make([][]*TreeNode, n+1)
	for i := 3; i <= n; i++ {
		x := (i - 3)
		y := (i - 3) / 2
		a := factorial[x]
		b := factorial[y]
		dp[i] = make([]*TreeNode, 0, a/b/b)
	}
	dp[1] = []*TreeNode{{0, nil, nil}}
	for count := 3; count <= n; count += 2 {
		for i := 1; i < count-1; i += 2 {
			j := count - 1 - i
			for _, left := range dp[i] {
				for _, right := range dp[j] {
					root := &TreeNode{Val: 0, Left: left, Right: right}
					dp[count] = append(dp[count], root)
				}
			}
		}
	}
	return dp[n]
}

func run(n int) [][]any {
	res := [][]any{}
	for _, tree := range allPossibleFBT(n) {
		res = append(res, TreeToSlice(tree))
	}
	return res
}
