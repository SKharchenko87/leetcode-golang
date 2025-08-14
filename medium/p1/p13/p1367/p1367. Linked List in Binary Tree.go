package p1367

import (
	. "leetcode/stucture"
	"slices"
	"strconv"
	"strings"
)

func isSubPath2(head *ListNode, root *TreeNode) bool {
	stacks := [][]*ListNode{}
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		for i, stack := range stacks {
			if stack[len(stack)-1] != nil && stack[len(stack)-1].Val == root.Val {
				if stack[len(stack)-1].Next == nil {
					return true
				}
				stacks[i] = append(stack, stack[len(stack)-1].Next)
			} else {
				stacks[i] = append(stack, nil)
			}
		}
		if head.Val == root.Val {
			if head.Next == nil {
				return true
			}
			stacks = append(stacks, []*ListNode{head.Next})
		}
		res := dfs(root.Left) || dfs(root.Right)
		for i, stack := range stacks {
			stacks[i] = stack[:len(stack)-1]
			l := len(stacks[i])
			if l == 0 {
				stacks = slices.Delete(stacks, i, i+1)
			}
		}
		return res
	}
	return dfs(root)
}

func isSubPath1(head *ListNode, root *TreeNode) bool {
	b := strings.Builder{}
	for head != nil {
		b.WriteRune(rune(head.Val))
		head = head.Next
	}
	targetPath := b.String()
	var dfs func(root *TreeNode, path string) bool
	dfs = func(root *TreeNode, path string) bool {
		if root == nil {
			return false
		}
		path += string(rune(root.Val))
		if strings.Contains(path, targetPath) {
			return true
		}
		return dfs(root.Left, path) || dfs(root.Right, path)
	}
	return dfs(root, "")
}

func isSubPath0(head *ListNode, root *TreeNode) bool {
	b := strings.Builder{}
	for head != nil {
		b.WriteRune('_')
		b.WriteString(strconv.Itoa(head.Val))
		b.WriteRune('_')
		head = head.Next
	}
	targetPath := b.String()
	var dfs func(root *TreeNode, path string) bool
	dfs = func(root *TreeNode, path string) bool {
		if root == nil {
			return false
		}
		path += "_" + strconv.Itoa(root.Val) + "_"
		if strings.Contains(path, targetPath) {
			return true
		}
		return dfs(root.Left, path) || dfs(root.Right, path)
	}
	return dfs(root, "")
}

func run(head []int, root []int) bool {
	return isSubPath(ArrToList(head), SliceToTreeNodeFullMinInt(root))
}
