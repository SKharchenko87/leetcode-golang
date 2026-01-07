package p1339

import . "leetcode/stucture"

const MOD = 1e9 + 7

func maxProduct(root *TreeNode) int {
	var res int64
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		return root.Val + l + r
	}
	totalSum := dfs(root)
	var dfs2 func(root *TreeNode) int
	dfs2 = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := dfs2(root.Left)
		r := dfs2(root.Right)
		res = max(res, int64((totalSum-l)*l), int64((totalSum-r)*r))
		return root.Val + l + r
	}
	dfs2(root)
	return int(res) % MOD
}
