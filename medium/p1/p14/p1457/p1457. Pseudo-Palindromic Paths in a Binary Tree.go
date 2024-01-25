package p1457

import . "leetcode/stucture"

func pseudoPalindromicPaths(root *TreeNode) int {
	path := [9]int{}
	res := 0
	countOdd := 0
	var changeCountOdd = func(i int, countOdd *int, cmd rune) {
		switch cmd {
		case '+':
			path[i]++
		case '-':
			path[i]--
		}
		if path[i]&1 == 1 {
			*countOdd++
		} else {
			*countOdd--
		}
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		changeCountOdd(root.Val-1, &countOdd, '+')
		if root.Left == nil && root.Right == nil {
			if countOdd <= 1 {
				res++
			}
			changeCountOdd(root.Val-1, &countOdd, '-')
			return
		}
		if root.Left != nil {
			dfs(root.Left)
		}
		if root.Right != nil {
			dfs(root.Right)
		}
		changeCountOdd(root.Val-1, &countOdd, '-')
	}
	dfs(root)
	return res
}
