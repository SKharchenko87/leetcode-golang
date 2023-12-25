package main

import (
	"fmt"
	. "leetcode/stucture"
)

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	if root.Left != nil {
		res = append(res, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, preorderTraversal(root.Right)...)
	}
	return res
}

func main() {
	//a := 1
	//b := 2
	//c := 3
	//m := [](*int){&a, nil, &b, &c}
	//for i, v := range m {
	//	if v != nil {
	//		fmt.Println(i, *v)
	//	} else {
	//		fmt.Println(i, v)
	//	}
	//}
	x := []int{1, 0, 2, 3}
	ss := SliceToTreeNode(x)
	fmt.Println(preorderTraversal(ss))
}
