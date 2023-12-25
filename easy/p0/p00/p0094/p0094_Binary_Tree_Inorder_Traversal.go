package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getVal(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		res = append(res, getVal(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, getVal(root.Right)...)
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	res := getVal(root)
	return res
}

func addTreeNode(i int, slice *[]int) *TreeNode {
	if i >= len(*slice) {
		return nil
	}
	leftIdx, rightIdx := getIndex(i)
	return &TreeNode{Val: (*slice)[i], Left: addTreeNode(leftIdx, slice), Right: addTreeNode(rightIdx, slice)}
}

func SliceToTreeNode(slice []int) *TreeNode {
	return addTreeNode(0, &slice)
}

func getIndex(x int) (left, right int) {
	if x == 0 {
		return 1, 2
	}
	var x_level int
	for x>>x_level > 0 {
		x_level++
	}
	pos_in_level := x - (1 << (x_level - 1)) + 1
	child_start_of_level := (1 << (x_level)) - 1
	left = child_start_of_level + pos_in_level*2
	right = child_start_of_level + pos_in_level*2 + 1
	return left, right
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	s := SliceToTreeNode(arr)

	fmt.Println(s)

	fmt.Println(len(arr))

	fmt.Println(getIndex(0))
	fmt.Println(getIndex(5))
	fmt.Println(getIndex(8))

	fmt.Println(inorderTraversal(s))

}
