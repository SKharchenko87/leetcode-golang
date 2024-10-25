package p2641

import (
	. "leetcode/stucture"
)

func replaceValueInTree(root *TreeNode) *TreeNode {
	stackNode := []*TreeNode{}
	stackLevel := []int{}
	levels := []int{}
	index := 0
	stackNode = append(stackNode, root)
	stackLevel = append(stackLevel, 0)

	levels = append(levels, 0)

	for index != len(stackNode) {
		if index > 1 && stackLevel[index] != stackLevel[index-1] {
			i := 0
			level := stackLevel[index-1]
			for level == stackLevel[index-1-i] {
				stackNode[index-1-i].Val = levels[level] - stackNode[index-1-i].Val
				i++
			}
		}
		currentNode := stackNode[index]
		currentLevel := stackLevel[index]

		if len(levels)-1 < currentLevel+1 {
			levels = append(levels, 0)
		}

		if ln := currentNode.Left; ln != nil {
			levels[currentLevel+1] += ln.Val
			if currentNode.Right != nil {
				ln.Val += currentNode.Right.Val
			}
			stackNode = append(stackNode, ln)
			stackLevel = append(stackLevel, currentLevel+1)
		}
		if rn := currentNode.Right; rn != nil {
			levels[currentLevel+1] += rn.Val
			if currentNode.Left != nil {
				rn.Val = currentNode.Left.Val
			}
			stackNode = append(stackNode, rn)
			stackLevel = append(stackLevel, currentLevel+1)
		}
		index++

	}

	if index-1 > 1 {
		i := 0
		level := stackLevel[index-1]
		for level == stackLevel[index-1-i] {
			stackNode[index-1-i].Val = levels[level] - stackNode[index-1-i].Val
			i++
		}
	}

	if root != nil {
		root.Val = 0
		if root.Left != nil {
			root.Left.Val = 0
		}
		if root.Right != nil {
			root.Right.Val = 0
		}
	}

	return root
}

func run(arr []int) []int {
	root := SliceToTreeNodeFullMinInt(arr)
	tree := replaceValueInTree(root)
	return TreeNodeToSlice(tree)
}
