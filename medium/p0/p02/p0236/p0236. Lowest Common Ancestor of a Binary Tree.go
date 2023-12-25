package main

import (
	"fmt"
	. "leetcode/stucture"
	"math"
)

func main() {
	rootArr, pArr, qArr := []int{3, 5, 1, 6, 2, math.MinInt, 8, math.MinInt, math.MinInt, 7, 4}, 5, 1
	//rootArr, pArr, qArr := []int{3, 5, 1, 6, 2, math.MinInt, 8, math.MinInt, math.MinInt, 7, 4}, 5, 4
	//rootArr, pArr, qArr := []int{1, 2}, 1, 2
	//rootArr, pArr, qArr := []int{1, 2}, 2, 1
	//rootArr, pArr, qArr := []int{1, 2, 3, math.MinInt, 4}, 4, 3
	root := SliceToTreeNode(rootArr, false, math.MinInt)
	//fmt.Println(root)
	p := SliceToTreeNode([]int{pArr}, false, math.MinInt)
	q := SliceToTreeNode([]int{qArr}, false, math.MinInt)
	fmt.Println(lowestCommonAncestor(root, p, q))
}

func recursive(root, x *TreeNode, slice *[]TreeNode) {
	if root == nil {
		return
	}
	if root.Val == x.Val {
		*slice = append(*slice, *root)
		return
	}
	if root.Left != nil && len(*slice) == 0 {
		recursive(root.Left, x, slice)
	}
	if root.Right != nil && len(*slice) == 0 {
		recursive(root.Right, x, slice)
	}
	if len(*slice) > 0 {
		*slice = append(*slice, *root)
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parr := []TreeNode{}
	qarr := []TreeNode{}
	//ToDo
	//chp:= make(chan []TreeNode)
	//chq:= make(chan []TreeNode)
	recursive(root, p, &parr)
	recursive(root, q, &qarr)
	lenp := len(parr)
	lenq := len(qarr)
	l := min(lenp, lenq)

	i := 0
	for ; i < l; i++ {
		if parr[lenp-1-i].Val != qarr[lenq-1-i].Val {
			return &parr[lenp-i]
		}
	}
	if i == l {
		return &parr[lenp-l]
	}
	return &parr[lenp-1-i]
}
