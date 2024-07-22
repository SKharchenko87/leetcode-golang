package p2096

import (
	"bytes"
	. "leetcode/stucture"
	"slices"
)

func getDirections(root *TreeNode, startValue int, destValue int) string {
	var findPath func(root *TreeNode, findVal *int, path *[]byte, direction byte) bool
	findPath = func(root *TreeNode, findVal *int, path *[]byte, direction byte) bool {
		if root == nil {
			return false
		}
		if root.Val == *findVal {
			return true
		}
		var isFindLeft, isFindRight bool
		if isFindLeft = findPath(root.Left, findVal, path, 'L'); isFindLeft {
			*path = append(*path, 'L')
		} else if isFindRight = findPath(root.Right, findVal, path, 'R'); isFindRight {
			*path = append(*path, 'R')
		}
		return isFindLeft || isFindRight
	}
	startPath := []byte{}
	destPath := []byte{}

	findPath(root, &startValue, &startPath, ' ')
	findPath(root, &destValue, &destPath, ' ')

	i := 0
	minLengthPath := min(len(startPath), len(destPath))
	for i < minLengthPath && startPath[len(startPath)-1-i] == destPath[len(destPath)-1-i] {
		i++
	}
	prefix := bytes.Repeat([]byte{'U'}, len(startPath)-i)
	slices.Reverse(destPath)
	return string(prefix) + string(destPath[i:])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getDirections0(root *TreeNode, startValue int, destValue int) string {
	var findPath func(root *TreeNode, findVal int, path *[]*TreeNode) bool
	findPath = func(root *TreeNode, findVal int, path *[]*TreeNode) bool {
		if root == nil {
			return false
		}
		if root.Val == findVal {
			*path = append(*path, root)
			return true
		}
		isFind := findPath(root.Left, findVal, path) || findPath(root.Right, findVal, path)
		if isFind {
			*path = append(*path, root)
		}
		return isFind
	}
	startPath := []*TreeNode{}
	destPath := []*TreeNode{}
	findPath(root, startValue, &startPath)
	findPath(root, destValue, &destPath)
	i := 0
	for ; i < min(len(startPath), len(destPath)); i++ {
		if startPath[len(startPath)-i-1] != destPath[len(destPath)-i-1] {
			break
		}
	}
	prefix := bytes.Repeat([]byte{'U'}, len(startPath)-i)
	suffix := make([]byte, len(destPath)-i)
	j := 0
	for ; i < len(destPath); i++ {
		if destPath[len(destPath)-i].Right == destPath[len(destPath)-i-1] {
			suffix[j] = 'R'
		} else if destPath[len(destPath)-i].Left == destPath[len(destPath)-i-1] {
			suffix[j] = 'L'
		}
		j++
	}
	return string(prefix) + string(suffix)
}

func run(slice []int, startValue int, destValue int) string {
	root := SliceToTreeNodeFullMinInt(slice)
	return getDirections(root, startValue, destValue)
}
