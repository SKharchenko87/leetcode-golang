package main

import "fmt"
import . "leetcode/stucture"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

//func addTreeNode(i int, slice *[]int) *TreeNode {
//	if i >= len(*slice) {
//		return nil
//	}
//	leftIdx, rightIdx := getIndex(i)
//	return &TreeNode{Val: (*slice)[i], Left: addTreeNode(leftIdx, slice), Right: addTreeNode(rightIdx, slice)}
//}

//func SliceToTreeNode(slice []int) *TreeNode {
//	return addTreeNode(0, &slice)
//}

//func getIndex(x int) (left, right int) {
//	if x == 0 {
//		return 1, 2
//	}
//	var x_level int
//	for x>>x_level > 0 {
//		x_level++
//	}
//	pos_in_level := x - (1 << (x_level - 1)) + 1
//	child_start_of_level := (1 << (x_level)) - 1
//	left = child_start_of_level + pos_in_level*2
//	right = child_start_of_level + pos_in_level*2 + 1
//	return left, right
//}

func main() {
	root := SliceToTreeNodeFullMinInt([]int{3, 9, 20, NULL, NULL, 15, 7})
	fmt.Println(maxDepth(root))
}
