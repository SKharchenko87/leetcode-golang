package p2583

import (
	. "leetcode/stucture"
	"sort"
)

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	largestLevelSums := []int64{}
	stack := []*TreeNode{}
	levels := []int{}
	stack = append(stack, root)
	levels = append(levels, 0)
	index := 0
	for len(stack) > 0 {
		node := stack[index]
		level := levels[index]
		if node.Left != nil {
			stack = append(stack, node.Left)
			levels = append(levels, level+1)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			levels = append(levels, level+1)
		}
		if len(largestLevelSums) <= level {
			largestLevelSums = append(largestLevelSums, int64(node.Val))
		} else {
			largestLevelSums[level] += int64(node.Val)
		}
		stack = stack[1:]
		levels = levels[1:]
	}
	sort.Slice(largestLevelSums, func(i, j int) bool {
		return largestLevelSums[i] > largestLevelSums[j]
	})
	if len(largestLevelSums) < k {
		return -1
	}
	return largestLevelSums[k-1]
}

func kthLargestLevelSum1(root *TreeNode, k int) int64 {
	largestLevelSums := []int64{}
	stack := []*TreeNode{}
	levels := []int{}
	stack = append(stack, root)
	levels = append(levels, 0)
	index := 0
	for index != len(stack) {
		node := stack[index]
		level := levels[index]
		if node.Left != nil {
			stack = append(stack, node.Left)
			levels = append(levels, level+1)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			levels = append(levels, level+1)
		}
		if len(largestLevelSums) <= level {
			largestLevelSums = append(largestLevelSums, int64(node.Val))
		} else {
			largestLevelSums[level] += int64(node.Val)
		}
		index++
	}
	sort.Slice(largestLevelSums, func(i, j int) bool {
		return largestLevelSums[i] > largestLevelSums[j]
	})
	if len(largestLevelSums) < k {
		return -1
	}
	return largestLevelSums[k-1]
}

func kthLargestLevelSum0(root *TreeNode, k int) int64 {
	largestLevelSums := []int64{}
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(largestLevelSums) <= level {
			largestLevelSums = append(largestLevelSums, int64(root.Val))
		} else {
			largestLevelSums[level] += int64(root.Val)
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	sort.Slice(largestLevelSums, func(i, j int) bool {
		return largestLevelSums[i] > largestLevelSums[j]
	})
	if len(largestLevelSums) < k {
		return -1
	}
	return largestLevelSums[k-1]
}

func run(arr []int, k int) int64 {
	root := SliceToTreeNodeFullMinInt(arr)
	return kthLargestLevelSum(root, k)
}
