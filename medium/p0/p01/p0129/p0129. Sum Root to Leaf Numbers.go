package p0129

import . "leetcode/stucture"

func sumNumbers(root *TreeNode) int {
	return sumN(root, 0)
}

func sumN(root *TreeNode, prev int) int {
	if root.Left == nil && root.Right == nil {
		return prev*10 + root.Val
	}
	res := 0
	if root.Left != nil {
		res = res + sumN(root.Left, prev*10+root.Val)
	}
	if root.Right != nil {
		res = res + sumN(root.Right, prev*10+root.Val)
	}
	return res
}
