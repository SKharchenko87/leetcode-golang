package main

import (
	"slices"
	"sort"
)

func main() {
	root1arr, root2arr := []int{1, 2, 3}, []int{1, 3, 2}
	root1, root2 := SliceToTreeNode(root1arr), SliceToTreeNode(root2arr)
	print(leafSimilar(root1, root2))

}

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

func traversal(root *TreeNode, res *[]int) {
	if root.Left == nil && root.Right == nil {
		*res = append(*res, root.Val)
	}
	if root.Left != nil {
		traversal(root.Left, res)
	}
	if root.Right != nil {
		traversal(root.Right, res)
	}
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	res1 := []int{}
	res2 := []int{}
	traversal(root1, &res1)
	traversal(root2, &res2)
	sort.Ints(res1)
	sort.Ints(res2)
	return slices.Equal(res1, res2)
}
