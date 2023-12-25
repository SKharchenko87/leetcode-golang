package p1161

import (
	. "leetcode/stucture"
	"math"
)

func rec(root *TreeNode, arr *[]int, level int) {
	curArr := (*arr)
	if root == nil {
		return
	}
	if len(curArr)-1 < level {
		*arr = append(curArr, root.Val)
	} else {
		(*arr)[level] += root.Val
	}
	rec(root.Left, arr, level+1)
	rec(root.Right, arr, level+1)
}

func maxLevelSum(root *TreeNode) int {
	arr := []int{}
	rec(root, &arr, 0)
	var maxLevel int
	mLS := math.MinInt32
	for i, v := range arr {
		if mLS < v {
			mLS = v
			maxLevel = i
		}
	}
	return maxLevel + 1
}
