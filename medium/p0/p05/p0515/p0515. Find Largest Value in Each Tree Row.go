package p0515

import . "leetcode/stucture"

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queueNode := []*TreeNode{}
	queueLevel := []int{}
	var result []int

	queueNode = append(queueNode, root)
	queueLevel = append(queueLevel, 0)
	for len(queueNode) > 0 {
		node := queueNode[0]
		curLevel := queueLevel[0]
		if curLevel >= len(result) {
			result = append(result, node.Val)
		} else if result[curLevel] < node.Val {
			result[curLevel] = node.Val
		}
		if node.Left != nil {
			queueNode = append(queueNode, node.Left)
			queueLevel = append(queueLevel, curLevel+1)
		}
		if node.Right != nil {
			queueNode = append(queueNode, node.Right)
			queueLevel = append(queueLevel, curLevel+1)
		}
		queueNode = queueNode[1:]
		queueLevel = queueLevel[1:]
	}
	return result
}

func run(root []int) []int {
	return largestValues(SliceToTreeNodeFullMinInt(root))
}
