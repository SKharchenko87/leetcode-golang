package main

import "fmt"
import . "leetcode/stucture"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	p := SliceToTreeNodeFullMinInt([]int{1, 2})
	q := SliceToTreeNodeFullMinInt([]int{1, 2, 3})
	fmt.Println(isSameTree(p, q))

}
