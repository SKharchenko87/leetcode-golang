package p2385

import (
	"fmt"
	. "leetcode/stucture"
)

func maxDepth(root *TreeNode, start int, res *int) (int, bool) {
	if root == nil {
		return 0, false
	}

	maxDepthLeft, foundLeft := maxDepth(root.Left, start, res)
	maxDepthRight, foundRight := maxDepth(root.Right, start, res)

	if root.Val == start {
		*res = max(maxDepthLeft, maxDepthRight)
		return 1, true
	}

	if foundRight {
		*res = max(*res, maxDepthLeft+maxDepthRight)
		return maxDepthRight + 1, true
	}
	if foundLeft {
		*res = max(*res, maxDepthLeft+maxDepthRight)
		return maxDepthLeft + 1, true
	}
	return 1 + max(maxDepthLeft, maxDepthRight), false
}

func amountOfTime(root *TreeNode, start int) int {
	res := 0
	maxDepth(root, start, &res)
	fmt.Println(res)
	return res
}
