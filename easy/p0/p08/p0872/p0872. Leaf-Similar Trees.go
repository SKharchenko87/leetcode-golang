package p0872

import (
	. "leetcode/stucture"
)

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go goRun(root1, ch1)
	go goRun(root2, ch2)
	for {
		var1, ok1 := <-ch1
		var2, ok2 := <-ch2
		if ok1 != ok2 || var1 != var2 {
			return false
		}
		if !(ok1 && ok2) {
			return true
		}
	}
}

func goRun(root *TreeNode, ch chan int) {
	dfs(root, ch)
	close(ch)
}

func dfs(root *TreeNode, ch chan int) {
	if root.Left == nil && root.Right == nil {
		ch <- root.Val
	} else {
		if root.Left != nil {
			dfs(root.Left, ch)
		}
		if root.Right != nil {
			dfs(root.Right, ch)
		}
	}
}
