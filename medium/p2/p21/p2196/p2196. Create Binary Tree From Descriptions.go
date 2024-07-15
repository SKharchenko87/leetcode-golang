package p2196

import . "leetcode/stucture"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func createBinaryTree(descriptions [][]int) *TreeNode {
	m := make(map[int]*TreeNode, len(descriptions))
	children := make(map[int]struct{}, len(descriptions))
	for i := 0; i < len(descriptions); i++ {
		if _, ok := m[descriptions[i][0]]; !ok {
			m[descriptions[i][0]] = &TreeNode{Val: descriptions[i][0]}
		}
		if _, ok := m[descriptions[i][1]]; !ok {
			m[descriptions[i][1]] = &TreeNode{Val: descriptions[i][1]}
		}
		children[descriptions[i][1]] = struct{}{}
		if descriptions[i][2] == 0 {
			m[descriptions[i][0]].Right = m[descriptions[i][1]]
		} else {
			m[descriptions[i][0]].Left = m[descriptions[i][1]]
		}
	}
	for _, node := range m {
		if _, ok := children[node.Val]; !ok {
			return node
		}
	}
	return nil
}

func createBinaryTree0(descriptions [][]int) *TreeNode {
	m := make(map[int]*TreeNode, len(descriptions))
	children := make(map[int]struct{}, len(descriptions))
	for i := 0; i < len(descriptions); i++ {
		var child, parent *TreeNode
		if _, ok := m[descriptions[i][0]]; ok {
			parent = m[descriptions[i][0]]
		} else {
			parent = &TreeNode{Val: descriptions[i][0]}
			m[descriptions[i][0]] = parent
		}
		if _, ok := m[descriptions[i][1]]; ok {
			child = m[descriptions[i][1]]
		} else {
			child = &TreeNode{Val: descriptions[i][1]}
			m[descriptions[i][1]] = child
		}
		children[descriptions[i][1]] = struct{}{}
		if descriptions[i][2] == 0 {
			parent.Right = child
		} else {
			parent.Left = child
		}
	}
	for _, node := range m {
		if _, ok := children[node.Val]; !ok {
			return node
		}
	}
	return nil
}

func run(descriptions [][]int) []int {
	root := createBinaryTree(descriptions)
	return TreeNodeToSlice(root)
}
