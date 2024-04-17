package p0623

import . "leetcode/stucture"

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		v := TreeNode{Val: val}
		v.Left = root
		return &v
	} else {
		bfs(root, val, depth, 1)
		return root
	}
}

func bfs(root *TreeNode, val int, depth int, level int) {
	level++
	if root == nil {
		return
	}
	if depth == level {
		vl, vr := TreeNode{Val: val}, TreeNode{Val: val}
		vl.Left = root.Left
		vr.Right = root.Right
		root.Left = &vl
		root.Right = &vr
		return
	} else {
		bfs(root.Left, val, depth, level)
		bfs(root.Right, val, depth, level)
	}
}
