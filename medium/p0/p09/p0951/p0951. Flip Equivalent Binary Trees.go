package p0951

import . "leetcode/stucture"

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	nodes1 := make([]*TreeNode, 0, 100)
	nodes2 := make([]*TreeNode, 0, 100)
	nodes1 = append(nodes1, root1)
	nodes2 = append(nodes2, root2)
	for len(nodes2) > 0 {
		node1 := nodes1[0]
		node2 := nodes2[0]
		nodes1 = nodes1[1:]
		nodes2 = nodes2[1:]
		if compareNodes(node1, node2) {
			if node1 != nil && node2 != nil {
				if compareNodes(node1.Left, node2.Left) && compareNodes(node1.Right, node2.Right) {

				} else if compareNodes(node1.Left, node2.Right) && compareNodes(node1.Right, node2.Left) {
					node1.Left, node1.Right = node1.Right, node1.Left
				} else {
					return false
				}
				nodes1 = append(nodes1, node1.Left, node1.Right)
				nodes2 = append(nodes2, node2.Left, node2.Right)
			}
		} else {
			return false
		}
	}
	return true
}

// dfs
func flipEquiv1(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(root1 *TreeNode, root2 *TreeNode) bool
	dfs = func(root1 *TreeNode, root2 *TreeNode) bool {
		if compareNodes(root1, root2) {
			if root1 != nil && root2 != nil {
				if compareNodes(root1.Left, root2.Left) && compareNodes(root1.Right, root2.Right) {

				} else if compareNodes(root1.Left, root2.Right) && compareNodes(root1.Right, root2.Left) {
					root1.Left, root1.Right = root1.Right, root1.Left
				} else {
					return false
				}
				return dfs(root1.Left, root2.Left) && dfs(root1.Right, root2.Right)
			}
		} else {
			return false
		}
		return true
	}
	return dfs(root1, root2)
}

// bfs
func flipEquiv0(root1 *TreeNode, root2 *TreeNode) bool {
	nodes1 := make([]*TreeNode, 0, 100)
	nodes2 := make([]*TreeNode, 0, 100)
	nodes1 = append(nodes1, root1)
	nodes2 = append(nodes2, root2)
	index := 0
	for len(nodes2) > index {
		node1 := nodes1[index]
		node2 := nodes2[index]

		if compareNodes(node1, node2) {
			if node1 != nil && node2 != nil {
				if compareNodes(node1.Left, node2.Left) && compareNodes(node1.Right, node2.Right) {

				} else if compareNodes(node1.Left, node2.Right) && compareNodes(node1.Right, node2.Left) {
					node1.Left, node1.Right = node1.Right, node1.Left
				} else {
					return false
				}
				nodes1 = append(nodes1, node1.Left, node1.Right)
				nodes2 = append(nodes2, node2.Left, node2.Right)
			}
		} else {
			return false
		}

		index++
	}
	return true
}

func compareNodes(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 != nil && root2 != nil && root1.Val == root2.Val {
		return true
	}
	return false
}

func run(root1Slice, root2Slice []int) bool {
	return flipEquiv(SliceToTreeNodeFullMinInt(root1Slice), SliceToTreeNodeFullMinInt(root2Slice))
}
