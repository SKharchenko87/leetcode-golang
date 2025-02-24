package p0889

import . "leetcode/stucture"

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// The first element in preorder is the root.
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}

	// Find the index of the left subtree's root in the postorder array.
	leftRootVal := preorder[1]
	leftRootIndex := 0
	for i, val := range postorder {
		if val == leftRootVal {
			leftRootIndex = i
			break
		}
	}

	// Recursively construct the left and right subtrees.
	root.Left = constructFromPrePost(preorder[1:leftRootIndex+2], postorder[:leftRootIndex+1])
	root.Right = constructFromPrePost(preorder[leftRootIndex+2:], postorder[leftRootIndex+1:len(postorder)-1])

	return root
}
func run(preorder []int, postorder []int) []any {
	return TreeToSlice(constructFromPrePost(preorder, postorder))
}
