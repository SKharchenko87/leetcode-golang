package p1261

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SliceToTree(slice []interface{}) *TreeNode {
	l := len(slice)
	if l == 0 {
		return nil
	}
	if slice[0] == nil {
		return &TreeNode{}
	}
	root := &TreeNode{Val: slice[0].(int)}
	prevLevelNodeIndex := 0
	prevLevelNodes := []*TreeNode{root}
	newLevelNodes := []*TreeNode{}
	for i := 1; i < l; i++ {
		if (prevLevelNodeIndex)/2 == len(prevLevelNodes) {
			prevLevelNodeIndex = 0
			prevLevelNodes, newLevelNodes = newLevelNodes, prevLevelNodes
			newLevelNodes = newLevelNodes[:0]
		}
		if slice[i] != nil {
			cur := &TreeNode{Val: slice[i].(int)}
			newLevelNodes = append(newLevelNodes, cur)
			if prevLevelNodeIndex%2 == 0 {
				prevLevelNodes[prevLevelNodeIndex/2].Left = cur
			} else {
				prevLevelNodes[prevLevelNodeIndex/2].Right = cur
			}
		}
		prevLevelNodeIndex++
	}
	return root
}

type FindElements struct {
	m map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	res := FindElements{make(map[int]bool)}
	root.Val = 0
	res.m[root.Val] = true
	arr := []*TreeNode{root}
	var newArr []*TreeNode
	for len(arr) != 0 {
		for i, node := range arr {
			if node.Left != nil {
				v := 2*arr[i].Val + 1
				res.m[v] = true
				node.Left.Val = v
				newArr = append(newArr, node.Left)
			}
			if node.Right != nil {
				v := 2*arr[i].Val + 2
				res.m[v] = true
				node.Right.Val = v
				newArr = append(newArr, node.Right)
			}
		}
		arr, newArr = newArr, arr
		newArr = newArr[:0]
	}

	return res
}

func (this *FindElements) Find(target int) bool {
	return this.m[target]
}

func run(commands []string, params [][]interface{}) []interface{} {
	var fe FindElements
	res := make([]interface{}, len(commands))
	var tree *TreeNode
	for i, command := range commands {
		switch command {
		case "FindElements":
			tree = SliceToTree(params[i])
			fe = Constructor(tree)
		case "find":
			res[i] = fe.Find(params[i][0].(int))
		}
	}
	return res
}
