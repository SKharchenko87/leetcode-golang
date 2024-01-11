package p1026

import (
	. "leetcode/stucture"
	"math"
)

func maxAncestorDiff(root *TreeNode) int {
	var res int
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) (int, int) {
	if root == nil {
		return math.MaxInt, math.MinInt
	}
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	}
	minLeft, maxLeft := dfs(root.Left, res)
	minRight, maxRight := dfs(root.Right, res)
	minCur, maxCur := min(minLeft, minRight, root.Val), max(maxLeft, maxRight, root.Val)
	*res = max(*res, root.Val-minCur, maxCur-root.Val)
	return minCur, maxCur
}

func max(x0, x1, x2 int) int {
	var tmp int
	if x0 < x1 {
		tmp = x1
	} else {
		tmp = x0
	}
	if x2 < tmp {
		return tmp
	} else {
		return x2
	}
}

func min(x0, x1, x2 int) int {
	var tmp int
	if x0 > x1 {
		tmp = x1
	} else {
		tmp = x0
	}
	if x2 > tmp {
		return tmp
	} else {
		return x2
	}
}
