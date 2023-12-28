package p0450

import (
	"fmt"
	. "leetcode/stucture"
	"math"
)

func searchBST(root *TreeNode, val int, prev *TreeNode) (*TreeNode, *TreeNode) {
	if root.Val == val {
		return root, prev
	} else if root.Val < val && root.Right != nil {
		return searchBST(root.Right, val, root)
	} else if root.Val > val && root.Left != nil {
		return searchBST(root.Left, val, root)
	}
	return nil, prev
}

func searchMinInRight(root *TreeNode, prev *TreeNode, min int) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, prev
	}
	if min > root.Val {
		min = root.Val
		if root.Left != nil {
			return searchMinInRight(root.Left, root, min)
		} else {
			return root, prev
		}
	}
	return root, prev
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key && root.Right == nil && root.Left == nil {
		return nil
	}
	x, prev := searchBST(root, key, nil)
	if x == nil {
		return root
	}
	if x.Right == nil && x.Left == nil {
		if prev.Left == x {
			prev.Left = nil
			return root
		}
		if prev.Right == x {
			prev.Right = nil
			return root
		}
	}
	minInt := math.MaxInt
	minNode, prevMinNode := searchMinInRight(x.Right, x, minInt)

	fmt.Println(x.Val, prevMinNode.Val, minNode.Val)
	if x == prevMinNode {
		x.Val = x.Right.Val
		x.Right = x.Right.Right
	} else if prevMinNode.Left != nil {
		x.Val = minNode.Val
		prevMinNode.Left = prevMinNode.Left.Right
	} else if prevMinNode.Right != nil {
		x.Val = minNode.Val
		prevMinNode.Val = minNode.Val
		prevMinNode.Left = minNode.Left
		prevMinNode.Right = minNode.Right
	}
	return root
}
