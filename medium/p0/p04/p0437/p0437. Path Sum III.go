package main

import (
	"fmt"
)

func main() {
	//rootArr, targetSum := []int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1}, 8
	//rootArr, targetSum := []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 5, 1}, 22
	//rootArr, targetSum := []int{1}, 0
	//rootArr, targetSum := []int{1}, 1
	rootArr, targetSum := []int{1, 2}, 1
	//rootArr, targetSum := []int{-2, 0, -3}, -5
	//rootArr, targetSum := []int{-2, 0, -3}, -3
	//rootArr, targetSum := []int{1, -2, -3}, -2
	//rootArr, targetSum := []int{0, 1, 1}, 1
	root := SliceToTreeNode(rootArr)
	fmt.Println(pathSum(root, targetSum))
}

func pathSumRecur(root *TreeNode, targetSum int, path []int) int {
	res := 0
	path = append(path, root.Val)
	l := len(path)
	sum := 0
	for i := l - 1; i >= 0; i-- {
		sum += path[i]
		if sum == targetSum {
			res++
		}
	}
	left := root.Left
	right := root.Right
	if right != nil {
		rpath := make([]int, l)
		copy(rpath, path)
		res = res + pathSumRecur(right, targetSum, rpath)
	}
	if left != nil {
		res = res + pathSumRecur(left, targetSum, path)
	}
	return res
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	path := []int{}
	return pathSumRecur(root, targetSum, path)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
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
