package p0865

import . "leetcode/stucture"

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	left := make([]*TreeNode, 0)
	right := make([]*TreeNode, 0)
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
			//path = path[:len(path)-1]
			//return
		}
		dfs(root.Left, path)
		dfs(root.Right, path)
		path = path[:level-1]
		return
	}
	path := make([]*TreeNode, 0)
	dfs(root, path)
	for i := maxLevel - 1; i >= 0; i-- {
		if left[i] == right[i] {
			return left[i]
		}
	}
	return root
}
