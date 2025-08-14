package p0513

import (
	. "leetcode/stucture"
)

func findBottomLeftValue(root *TreeNode) int {
	res := root.Val
	maxLevel := 0
	var dfs func(level int, root *TreeNode)
	dfs = func(level int, root *TreeNode) {
		if root == nil {
			return
		}
		if maxLevel < level {
			maxLevel = level
			res = root.Val
		}
		dfs(level+1, root.Left)
		dfs(level+1, root.Right)
	}
	dfs(1, root.Left)
	dfs(1, root.Right)
	return res
}
