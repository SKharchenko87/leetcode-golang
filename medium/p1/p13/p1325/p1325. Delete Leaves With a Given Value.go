package p1325

import . "leetcode/stucture"

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if dfs(root, target) {
		return nil
	} else {
		return root
	}
}

func dfs(root *TreeNode, target int) bool {
	if root.Left != nil && dfs(root.Left, target) {
		root.Left = nil
	}
	if root.Right != nil && dfs(root.Right, target) {
		root.Right = nil
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == target {
			return true
		}
	}
	return false
}
