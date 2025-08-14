package p2

import "slices"

type TreeNode struct {
	Val      int
	ch       byte
	children []*TreeNode
}

func findSubtreeSizes(parent []int, s string) []int {
	arr := make([]*TreeNode, len(parent))
	for i := 0; i < len(parent); i++ {
		arr[i] = &TreeNode{ch: s[i], children: make([]*TreeNode, 0)}
	}
	var root *TreeNode
	for i := 0; i < len(parent); i++ {
		if parent[i] == -1 {
			root = arr[i]
		} else {
			arr[parent[i]].children = append(arr[parent[i]].children, arr[i])
		}

	}
	var dfs func(root *TreeNode, parent *TreeNode, i int) int
	m := [26]*TreeNode{}
	dfs = func(root *TreeNode, parent *TreeNode, i int) int {
		if root == nil {
			return 0
		}
		if m[root.ch-'a'] != nil {
			parent.children = append(parent.children, root.children...)
			root.children = []*TreeNode{}
			slices.Delete(parent.children, i, i)
			return -1
		} else {
			m[root.ch-'a'] = root
		}
		for j := 0; j < len(root.children); j++ {
			j += dfs(root.children[j], root, j)
		}
		return 0
	}
	for j := 0; j < len(root.children); j++ {
		m[root.ch-'a'] = root
		j += dfs(root.children[j], nil, j)
	}

	return parent
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
