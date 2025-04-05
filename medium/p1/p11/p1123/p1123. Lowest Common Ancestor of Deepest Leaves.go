package p1123

import . "leetcode/stucture"

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	lca, _ := dfs(root, 0)
	return lca
}

func dfs(node *TreeNode, depth int) (*TreeNode, int) {
	if node == nil {
		return nil, depth - 1
	}

	left, leftDepth := dfs(node.Left, depth+1)
	right, rightDepth := dfs(node.Right, depth+1)

	if leftDepth > rightDepth {
		return left, leftDepth
	}
	if leftDepth < rightDepth {
		return right, rightDepth
	}

	// doesnt matter to return left or right depth since both are equal here
	return node, leftDepth
}

func run(root []any) []any {
	arr := SliceToTree(root)
	tree := lcaDeepestLeaves(arr)
	return TreeToSlice(tree)
}
