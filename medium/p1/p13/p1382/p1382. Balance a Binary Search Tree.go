package p1382

import (
	. "leetcode/stucture"
	. "math"
)

func balanceBST1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	head := &TreeNode{0, nil, root}
	current := head
	for current.Right != nil {
		if current.Right.Left != nil {
			rightRotation(current, current.Right)
		} else {
			current = current.Right
		}
	}

	nodeCount := 0
	current = head.Right
	for current != nil {
		nodeCount++
		current = current.Right
	}

	m := int(Pow(2, Floor(Log2(float64(nodeCount+1)))) - 1)
	makeRotationTree(head, nodeCount-m)
	for m > 1 {
		m /= 2
		makeRotationTree(head, m)
	}
	return head.Right
}

func rightRotation(parent *TreeNode, node *TreeNode) {
	var tmp *TreeNode = node.Left
	node.Left = tmp.Right
	tmp.Right = node
	parent.Right = tmp
}

func leftRotation(parent *TreeNode, node *TreeNode) {
	var tmp *TreeNode = node.Right
	node.Right = tmp.Left
	tmp.Left = node
	parent.Right = tmp
}

func makeRotationTree(head *TreeNode, count int) {
	current := head
	for i := 0; i < count; i++ {
		tmp := current.Right
		leftRotation(current, tmp)
		current = current.Right
	}
}

////////////////////////

func balanceBST(root *TreeNode) *TreeNode {
	arr := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		arr = append(arr, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	var sliceSortToBinaryTree func(arr []int, start, end int) *TreeNode
	sliceSortToBinaryTree = func(arr []int, start, end int) *TreeNode {
		if start > end {
			return nil
		}
		node := &TreeNode{Val: arr[(start+end)/2]}
		node.Left = sliceSortToBinaryTree(arr, start, (start+end)/2-1)
		node.Right = sliceSortToBinaryTree(arr, (start+end)/2+1, end)
		return node
	}
	res := sliceSortToBinaryTree(arr, 0, len(arr)-1)
	return res
}

func helper(root []int) []int {
	tree := SliceToTreeNodeFullMinInt(root)
	resTree := balanceBST(tree)
	res := TreeNodeToSlice(resTree)
	return res
}
