package p0590

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	res := []int{}
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, child := range node.Children {
			dfs(child)
		}
		if 0 <= node.Val && node.Val <= 10000 {
			res = append(res, node.Val)
		}
	}
	dfs(root)
	return res
}
