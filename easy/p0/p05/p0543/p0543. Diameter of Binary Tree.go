package p0543

import . "leetcode/stucture"

func diameterOfBinaryTree(root *TreeNode) int {
	var res int = 0
	radiusOfBinaryTree(&res, root)
	return res
}

func radiusOfBinaryTree(res *int, root *TreeNode) int {
	if root == nil {
		return 0
	}
	r := radiusOfBinaryTree(res, root.Right)
	l := radiusOfBinaryTree(res, root.Left)
	*res = max(*res, r+l)
	return max(r, l) + 1
}
