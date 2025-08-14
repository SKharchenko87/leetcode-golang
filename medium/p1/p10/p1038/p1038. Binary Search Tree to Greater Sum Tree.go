package p1038

import . "leetcode/stucture"

func bstToGst(root *TreeNode) *TreeNode {
	var sum int
	node := root
	for node != nil {
		if node.Right == nil {
			sum += node.Val
			node.Val = sum
			node = node.Left
		} else {
			succ := getSuccessor(node)
			if succ.Left == nil {
				succ.Left = node
				node = node.Right
			} else {
				succ.Left = nil
				sum += node.Val
				node.Val = sum
				node = node.Left
			}

		}
	}
	return root
}

func getSuccessor(node *TreeNode) *TreeNode {
	succ := node.Right
	for succ.Left != nil && succ.Left != node {
		succ = succ.Left
	}
	return succ
}

func bstToGst0(root *TreeNode) *TreeNode {
	var sum int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		sum += root.Val
		root.Val = sum
		dfs(root.Left)
	}
	dfs(root)
	return root
}

func helper(arr []int) []int {
	root := SliceToTreeNodeFullMinInt(arr)
	tmp := bstToGst(root)
	res := TreeNodeToSlice(tmp)
	return res
}
