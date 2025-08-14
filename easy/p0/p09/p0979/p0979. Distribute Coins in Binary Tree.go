package p0979

import . "leetcode/stucture"

func distributeCoins(root *TreeNode) int {
	movement, _ := dfs(root)
	return movement
}

func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	leftMovement, leftBalance := dfs(node.Left)
	rightMovement, rightBalance := dfs(node.Right)
	totalBalance := leftBalance + rightBalance + node.Val - 1
	return leftMovement + rightMovement + abs(totalBalance), totalBalance
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
