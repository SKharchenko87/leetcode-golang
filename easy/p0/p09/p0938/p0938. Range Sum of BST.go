package p0938

import . "leetcode/stucture"

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	tmp := 0
	if low <= root.Val && root.Val <= high {
		tmp = root.Val
	}
	return tmp + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
