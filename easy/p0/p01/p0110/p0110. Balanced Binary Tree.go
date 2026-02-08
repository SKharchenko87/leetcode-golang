package p0110

import . "leetcode/stucture"

func dfs(root *TreeNode) (int, bool) {
	if root == nil {
		return 1, true
	}
	left, bl := dfs(root.Left)
	right, br := dfs(root.Right)
	return max(left, right) + 1, bl && br && abs(left-right) < 2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isBalanced(root *TreeNode) bool {
	_, res := dfs(root)
	return res
}

func run(root []int) bool {
	return isBalanced(SliceToTreeNodeFullMinInt(root))
}
