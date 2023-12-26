package p0199

import (
	. "leetcode/stucture"
)

func recursing(root *TreeNode, res *map[int]int, level int) *TreeNode {
	if root == nil {
		return nil
	}
	if _, ok := (*res)[level]; !ok {
		(*res)[level] = root.Val
	}
	if root.Right != nil {
		recursing(root.Right, res, level+1)
	}
	if root.Left != nil {
		recursing(root.Left, res, level+1)
	}
	return nil
}

func rightSideView(root *TreeNode) []int {
	res := map[int]int{}
	if root != nil {
		recursing(root, &res, 0)
	}
	r := make([]int, len(res))
	for k, v := range res {
		r[k] = v
	}
	return r
}
