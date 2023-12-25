package main

func main() {
	rootArr := []int{3, 1, 4, 3, 0, 1, 5}
	root1 := SliceToTreeNode(rootArr)
	print(goodNodes(root1))
}

func pathMax(root *TreeNode, m int) int {
	i := 0
	if root.Val >= m {
		i++
		m = root.Val
	}
	if root.Left != nil {
		i += pathMax(root.Left, m)
	}
	if root.Right != nil {
		i += pathMax(root.Right, m)
	}
	return i

}

func goodNodes(root *TreeNode) int {
	return pathMax(root, -10001)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getVal(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		res = append(res, getVal(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, getVal(root.Right)...)
	}
	return res
}

func addTreeNode(i int, slice *[]int) *TreeNode {
	if i >= len(*slice) {
		return nil
	}
	leftIdx, rightIdx := getIndex(i)
	return &TreeNode{Val: (*slice)[i], Left: addTreeNode(leftIdx, slice), Right: addTreeNode(rightIdx, slice)}
}

func SliceToTreeNode(slice []int) *TreeNode {
	return addTreeNode(0, &slice)
}

func getIndex(x int) (left, right int) {
	if x == 0 {
		return 1, 2
	}
	var x_level int
	for x>>x_level > 0 {
		x_level++
	}
	pos_in_level := x - (1 << (x_level - 1)) + 1
	child_start_of_level := (1 << (x_level)) - 1
	left = child_start_of_level + pos_in_level*2
	right = child_start_of_level + pos_in_level*2 + 1
	return left, right
}
