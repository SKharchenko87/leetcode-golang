package p2458

import . "leetcode/stucture"

type LevelLength struct {
	node0, node1     int
	length0, length1 int
}

func treeQueries(root *TreeNode, queries []int) []int {
	nodeLevel := make([]int, 100_001)
	levelMaxLength := []LevelLength{}
	var dfs func(root *TreeNode, level int) int
	dfs = func(root *TreeNode, level int) int {
		if root == nil {
			return 1
		}
		l := max(dfs(root.Left, level+1), dfs(root.Right, level+1))
		nodeLevel[root.Val] = level
		if level+1 >= len(levelMaxLength) {
			for i := len(levelMaxLength); i <= level+1; i++ {
				levelMaxLength = append(levelMaxLength, LevelLength{0, 0, 0, 0})
			}
		}
		if l-1 >= levelMaxLength[level].length0 {
			levelMaxLength[level].length1, levelMaxLength[level].length0 = levelMaxLength[level].length0, l-1
			levelMaxLength[level].node1, levelMaxLength[level].node0 = levelMaxLength[level].node0, root.Val
		} else if l-1 >= levelMaxLength[level].length1 {
			levelMaxLength[level].length1 = l - 1
			levelMaxLength[level].node1 = root.Val
		}
		return l + 1
	}
	dfs(root, 0)
	res := make([]int, len(queries))
	for i, query := range queries {
		level := nodeLevel[query]
		if levelMaxLength[level].node0 == query {
			if levelMaxLength[level].node1 != 0 {
				res[i] = level + levelMaxLength[level].length1
			} else {
				res[i] = level - 1
			}
		} else {
			if levelMaxLength[level].node1 != 0 {
				res[i] = level + levelMaxLength[level].length0
			} else {
				res[i] = level - 1
			}
		}
	}
	return res
}

func run(root []int, queries []int) []int {
	return treeQueries(SliceToTreeNodeFullMinInt(root), queries)
}
