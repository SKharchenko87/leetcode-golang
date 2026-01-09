package p0865

import . "leetcode/stucture"

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var left, right, path []*TreeNode
	maxLevel := -1
	var dfs func(root *TreeNode, path []*TreeNode)
	dfs = func(root *TreeNode, path []*TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root)
		level := len(path)
		if root.Left == nil && root.Right == nil {
			if level > maxLevel {
				left = append([]*TreeNode(nil), path...)
				right = left
				maxLevel = level
			} else if level == maxLevel {
				right = append([]*TreeNode(nil), path...)
			}
		}
		dfs(root.Left, path)
		dfs(root.Right, path)
		path = path[:level-1]
		return
	}
	dfs(root, path)
	for i := maxLevel - 1; i >= 0; i-- {
		if left[i] == right[i] {
			return left[i]
		}
	}
	return root
}
