package p0988

import . "leetcode/stucture"

func smallestFromLeaf(root *TreeNode) string {
	arr := make([]rune, 8500)
	return bfs(root, arr, 0)
}

func bfs(root *TreeNode, arr []rune, level int) string {
	level++
	if root.Left == nil && root.Right == nil {
		arr[level] = rune(root.Val) + 'a'
		var s []rune
		for level > 0 {
			s = append(s, arr[level])
			level--
		}
		return string(s)
	}
	arr[level] = rune(root.Val) + 'a'
	if root.Left != nil && root.Right != nil {
		ls := bfs(root.Left, arr, level)
		rs := bfs(root.Right, arr, level)
		if ls < rs {
			return ls
		} else {
			return rs
		}
	} else if root.Left != nil {
		return bfs(root.Left, arr, level)
	} else {
		return bfs(root.Right, arr, level)
	}
}
