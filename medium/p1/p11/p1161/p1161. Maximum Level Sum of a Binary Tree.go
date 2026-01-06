package p1161

import (
	. "leetcode/stucture"
	"math"
)

func maxLevelSum(root *TreeNode) int {
	var curNodes = []*TreeNode{}

	resSum := math.MinInt32
	res := 0
	level := 0
	curNodes = append(curNodes, root)
	size := 1
	i := 0
	for {
		level++
		levelSum := 0
		newSize := 0
		for ; i < size; i++ {
			levelSum += curNodes[i].Val
			if curNodes[i].Left != nil {
				curNodes = append(curNodes, curNodes[i].Left)
				newSize++
			}
			if curNodes[i].Right != nil {
				curNodes = append(curNodes, curNodes[i].Right)
				newSize++
			}
		}
		if levelSum > resSum {
			resSum = levelSum
			res = level
		}
		if newSize == 0 {
			break
		}
		size += newSize
	}
	return res
}

func maxLevelSum1(root *TreeNode) int {
	curNodes := make([]*TreeNode, 0, 120000)
	nextNodes := make([]*TreeNode, 0, 120000)

	resSum := math.MinInt64
	res := 0
	level := 0
	curNodes = append(curNodes, root)
	for {
		level++
		levelSum := 0
		for i := 0; i < len(curNodes); i++ {
			levelSum += curNodes[i].Val
			if curNodes[i].Left != nil {
				nextNodes = append(nextNodes, curNodes[i].Left)
			}
			if curNodes[i].Right != nil {
				nextNodes = append(nextNodes, curNodes[i].Right)
			}
		}
		if levelSum > resSum {
			resSum = levelSum
			res = level
		}
		if len(nextNodes) == 0 {
			break
		}
		curNodes, nextNodes = nextNodes, curNodes
		nextNodes = nextNodes[:0]
	}
	return res
}

/*
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
++++++++++
*/
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

func maxLevelSum0(root *TreeNode) int {
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
