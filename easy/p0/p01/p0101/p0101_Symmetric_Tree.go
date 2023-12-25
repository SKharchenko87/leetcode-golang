package main

import "fmt"
import . "leetcode/stucture"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isSymmetricTree(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil || l.Val != r.Val {
		return false
	}
	return isSymmetricTree(l.Left, r.Right) && isSymmetricTree(l.Right, r.Left)
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricTree(root.Left, root.Right)
}

//func addTreeNode(i int, slice *[]int) *TreeNode {
//	if i >= len(*slice) {
//		return nil
//	}
//	leftIdx, rightIdx := getIndex(i)
//	return &TreeNode{Val: (*slice)[i], Left: addTreeNode(leftIdx, slice), Right: addTreeNode(rightIdx, slice)}
//}
//
//func SliceToTreeNode(slice []int) *TreeNode {
//	return addTreeNode(0, &slice)
//}
//
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
	root := SliceToTreeNodeFullMinInt([]int{1, 2, 2, 3, 4, 3, 4})
	fmt.Println(isSymmetric(root))
}
