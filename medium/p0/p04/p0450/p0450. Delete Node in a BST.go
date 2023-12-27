package p0450

import . "leetcode/stucture"

func searchBST(root *TreeNode, val int) *TreeNode {
	if root.Val == val {
		return root
	} else if root.Val < val && root.Right != nil {
		return searchBST(root.Right, val)
	} else if root.Val > val && root.Left != nil {
		return searchBST(root.Left, val)
	}
	return nil
}

func searchMinInRight(root *TreeNode, prev *TreeNode, min int) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, prev
	}
	if min > root.Val {
		min = root.Val
		if root.Left != nil {
			searchMinInRight(root.Left, root, min)
		} else {
			return root, prev
		}
	}
	return nil, prev
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	x := searchBST(root, key)
	if x == nil {
		return nil
	}
	prevMinNode, minNode := searchMinInRight(x.Right, root, root.Val)
	prevMinNode.Left = minNode.Right
	return root
}
