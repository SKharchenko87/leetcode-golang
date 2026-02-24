package p1022

import (
	. "leetcode/stucture"
)

func sumRootToLeaf(root *TreeNode) int {
	var dfs func(root *TreeNode, path uint) uint
	dfs = func(root *TreeNode, path uint) uint {
		if root == nil {
			return 0
		}

		path = (path << 1) | uint(root.Val)

		if root.Left == nil && root.Right == nil {
			return path
		}

		return dfs(root.Left, path) + dfs(root.Right, path)
	}

	return int(dfs(root, 0))
}

func sumRootToLeaf0(root *TreeNode) int {
	res := 0
	var dfs func(root *TreeNode, path int)
	dfs = func(root *TreeNode, path int) {
		if root == nil {
			return
		}

		path <<= 1
		path = path | root.Val

		if root.Left == nil && root.Right == nil {
			res += path
		}

		dfs(root.Left, path)
		dfs(root.Right, path)
	}
	dfs(root, 0)
	return res
}

func run(root []int) int {
	return sumRootToLeaf(SliceToTreeNodeFullMinInt(root))
}
