package p1609

import . "leetcode/stucture"

func isEvenOddTree(root *TreeNode) bool {
	var bfs func(level int, root *TreeNode) bool
	arr := make([]int, 2000)
	bfs = func(level int, root *TreeNode) bool {
		if root == nil {
			return true
		}
		if level%2 == root.Val%2 {
			return false
		}
		if len(arr) <= level {
			arr = append(arr, 0)
		}
		if arr[level] != 0 && (level%2 == 0 && arr[level] >= root.Val || level%2 == 1 && arr[level] <= root.Val) {
			return false
		}
		arr[level] = root.Val
		return bfs(level+1, root.Left) && bfs(level+1, root.Right)
	}
	return bfs(0, root)
}
