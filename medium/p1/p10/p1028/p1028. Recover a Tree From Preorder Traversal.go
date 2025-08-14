package p1028

import (
	. "leetcode/stucture"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Stack []any

func (s *Stack) Push(v any) {
	*s = append(*s, v)
}
func (s *Stack) Pop() any {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}
func (s *Stack) Top() any {
	l := len(*s)
	v := (*s)[l-1]
	return v
}
func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func recoverFromPreorder(traversal string) *TreeNode {
	level := 0
	numStartIndex, numEndIndex := -1, -1
	levels, vals := make([]int, 0), make([]int, 0)
	for i, c := range traversal {
		if c == '-' {
			if numEndIndex != -1 {
				val, _ := strconv.Atoi(traversal[numStartIndex : numEndIndex+1])
				vals = append(vals, val)
				numStartIndex = -1
				numEndIndex = -1
			}
			level++
		} else if numStartIndex == -1 {
			levels = append(levels, level)
			level = 0
			numStartIndex = i
			numEndIndex = i
		} else {
			numEndIndex = i
		}
	}
	val, _ := strconv.Atoi(traversal[numStartIndex : numEndIndex+1])
	vals = append(vals, val)

	root := &TreeNode{Val: vals[0]}
	valTreeStack := Stack{root}
	levelStack := Stack{0}
	for i := 1; i < len(levels); i++ {
		node := &TreeNode{Val: vals[i]}
		for levelStack.Top().(int)+1 != levels[i] {
			levelStack.Pop()
			valTreeStack.Pop()
		}

		if valTreeStack.Top().(*TreeNode).Left == nil {
			valTreeStack.Top().(*TreeNode).Left = node
		} else {
			valTreeStack.Pop().(*TreeNode).Right = node
			levelStack.Pop()
		}
		levelStack.Push(levels[i])
		valTreeStack.Push(node)

	}
	return root
}

func run(traversal string) []interface{} {
	return TreeToSlice(recoverFromPreorder(traversal))
}
