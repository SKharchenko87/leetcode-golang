package p0145

import (
	. "leetcode/stucture"
)

func postorderTraversal(root *TreeNode) (res []int) {
	res = make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		if -100 <= root.Val && root.Val <= 100 {
			res = append(res, root.Val)
		}
	}
	dfs(root)
	return
}

func run(arr []int) []int {
	return postorderTraversal(SliceToTreeNodeFullMinInt(arr))
}
