package main

import (
	"fmt"
)
import . "leetcode/stucture"

func main() {
	//rootArr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	//rootArr := []int{1, 2, 3, 0, 4, 0, 0, 5, 6, 0, 7}
	rootArr := []int{1, 0, 2, 3, 4, 0, 0, 5, 6, 0, 7, 0, 0, 0, 8}
	//rootArr := []int{6, 9, 7, 3, 0, 2, 8, 5, 8, 9, 7, 3, 9, 9, 4, 2, 10, 0, 5, 4, 3, 10, 10, 9, 4, 1, 2, 0, 0, 6, 5, 0, 0, 0, 0, 9, 0, 9, 6, 5, 0, 5, 0, 0, 7, 7, 4, 0, 1, 0, 0, 3, 7, 0, 9, 0, 0, 0, 0, 0, 0, 0, 0, 9, 9, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6, 8, 7, 0, 0, 0, 3, 10, 0, 0, 0, 0, 0, 1, 0, 1, 2}
	//rootArr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	root := SliceToTreeNode(rootArr, false, 0)
	//fmt.Println(root)
	fmt.Println(longestZigZag(root))
}
func longestZigZag(root *TreeNode) int {
	l := maxDepthZigZag(root.Left, 'L', 1)
	r := maxDepthZigZag(root.Right, 'R', 1)
	return max(l, r) - 1
}

func maxDepthZigZag(root *TreeNode, w byte, cnt int) int {
	if root == nil {
		return cnt
	}
	if w == 'L' {
		return max(maxDepthZigZag(root.Right, 'R', cnt+1), maxDepthZigZag(root.Left, 'L', 1))
	} else {
		return max(maxDepthZigZag(root.Left, 'L', cnt+1), maxDepthZigZag(root.Right, 'R', 1))
	}
}
