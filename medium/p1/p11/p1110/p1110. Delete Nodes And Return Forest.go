package p1110

import (
	. "leetcode/stucture"
	"sort"
)

type queue[T any] []T

func (q *queue[T]) Len() int {
	return len(*q)
}
func (q *queue[T]) Push(x T) {
	*q = append(*q, x)
}
func (q *queue[T]) PopLeft() T {
	if len(*q) == 0 {
		panic("stack is empty")
	}
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func delNodes(root *TreeNode, toDelete []int) []*TreeNode {

	deleteMap := make(map[int]bool)
	for i := 0; i < len(toDelete); i++ {
		deleteMap[toDelete[i]] = true
	}

	forest := make([]*TreeNode, 0)
	if !deleteMap[root.Val] {
		forest = append(forest, root)
	}

	q := queue[*TreeNode]{root}
	for {
		if q.Len() == 0 {
			break
		}
		cur := q.PopLeft()

		if cur.Left != nil {
			q.Push(cur.Left)
			if deleteMap[cur.Left.Val] {
				cur.Left = nil
			}
		}
		if cur.Right != nil {
			q.Push(cur.Right)
			if deleteMap[cur.Right.Val] {
				cur.Right = nil
			}
		}

		if deleteMap[cur.Val] {
			if cur.Left != nil {
				forest = append(forest, cur.Left)
				cur.Left = nil
			}
			if cur.Right != nil {
				forest = append(forest, cur.Right)
				cur.Left = nil
			}
		}
	}
	return forest
}

type stack[T any] []T

func (s *stack[T]) Push(v T) {
	*s = append(*s, v)
}
func (s *stack[T]) Pop() T {
	l := len(*s)
	if l == 0 {
		panic("stack is empty")
	}
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}
func (s *stack[T]) Peek() T {
	l := len(*s)
	if l == 0 {
		panic("stack is empty")
	}
	return (*s)[l-1]
}
func (s *stack[T]) Len() int {
	return len(*s)
}

func delNodes4(root *TreeNode, toDelete []int) []*TreeNode {
	deleteMap := make(map[int]bool)
	for i := 0; i < len(toDelete); i++ {
		deleteMap[toDelete[i]] = true
	}
	result := make([]*TreeNode, 0)
	if !deleteMap[root.Val] {
		result = append(result, root)
	}
	forest := stack[*TreeNode]{}
	cur := root
	var isDel = false
	for cur != nil || forest.Len() > 0 {
		if forest.Len() > 0 {
			cur = forest.Pop()

			curRight := cur.Right

			print(cur.Val)
			if cur.Left != nil && deleteMap[cur.Left.Val] {
				cur.Left = nil
			}
			if cur.Right != nil && deleteMap[cur.Right.Val] {
				cur.Right = nil
			}
			isDel = deleteMap[cur.Val]
			if isDel && cur.Left != nil {
				result = append(result, cur.Left)
				cur.Left = nil
			}
			if isDel && cur.Right != nil {
				result = append(result, cur.Right)
				cur.Left = nil
			}
			cur = curRight
		}
		for cur != nil {
			forest.Push(cur)
			cur = cur.Left
		}
	}
	return result
}

func delNodes3(root *TreeNode, toDelete []int) []*TreeNode {
	deleteMap := [1000]bool{}
	for i := 0; i < len(toDelete); i++ {
		deleteMap[toDelete[i]-1] = true
	}
	result := make([]*TreeNode, 0)
	if !deleteMap[root.Val-1] {
		result = append(result, root)
	}
	forest := stack[*TreeNode]{}
	cur := root
	var isDel = false
	for cur != nil || forest.Len() > 0 {
		if forest.Len() > 0 {
			cur = forest.Pop()

			curRight := cur.Right

			print(cur.Val)
			if cur.Left != nil && deleteMap[cur.Left.Val-1] {
				cur.Left = nil
			}
			if cur.Right != nil && deleteMap[cur.Right.Val-1] {
				cur.Right = nil
			}
			isDel = deleteMap[cur.Val-1]
			if isDel && cur.Left != nil {
				result = append(result, cur.Left)
				cur.Left = nil
			}
			if isDel && cur.Right != nil {
				result = append(result, cur.Right)
				cur.Left = nil
			}
			cur = curRight
		}
		for cur != nil {
			forest.Push(cur)
			cur = cur.Left
		}
	}
	return result
}

func delNodes2(root *TreeNode, toDelete []int) []*TreeNode {
	deleteMap := [1000]bool{}
	for i := 0; i < len(toDelete); i++ {
		deleteMap[toDelete[i]-1] = true
	}
	result := make([]*TreeNode, 0)
	if !deleteMap[root.Val-1] {
		result = append(result, root)
	}
	var (
		isDel bool
		dfs   func(root *TreeNode)
	)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if isDel {
			root.Left = nil
		}
		dfs(root.Right)
		if isDel {
			root.Right = nil
		}
		isDel = deleteMap[root.Val-1]
		if isDel && root.Left != nil {
			result = append(result, root.Left)
			root.Left = nil
		}
		if isDel && root.Right != nil {
			result = append(result, root.Right)
			root.Left = nil
		}
	}
	dfs(root)
	return result
}

func delNodes1(root *TreeNode, toDelete []int) []*TreeNode {
	deleteMap := make(map[int]bool)
	for i := 0; i < len(toDelete); i++ {
		deleteMap[toDelete[i]] = true
	}
	result := make([]*TreeNode, 0)
	if !deleteMap[root.Val] {
		result = append(result, root)
	}
	var isDel bool
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if isDel {
			root.Left = nil
		}
		dfs(root.Right)
		if isDel {
			root.Right = nil
		}
		isDel = deleteMap[root.Val]
		if isDel && root.Left != nil {
			result = append(result, root.Left)
			root.Left = nil
		}
		if isDel && root.Right != nil {
			result = append(result, root.Right)
			root.Left = nil
		}
	}
	dfs(root)
	return result
}

type nodeValLevel struct {
	val   int
	level int
}

func delNodes0(root *TreeNode, toDelete []int) []*TreeNode {
	deleteMap := make(map[int]int)
	for i := 0; i < len(toDelete); i++ {
		deleteMap[toDelete[i]] = -1
	}
	parents := map[int]*TreeNode{}
	nodeValLevels := []nodeValLevel{}
	dfsGetParents(root, nil, &parents, 0, &nodeValLevels, &deleteMap)

	sort.Slice(nodeValLevels, func(i, j int) bool {
		return nodeValLevels[i].level > nodeValLevels[j].level
	})

	result := []*TreeNode{}
	if _, ok := deleteMap[root.Val]; !ok {
		result = append(result, root)
	}
	for _, vl := range nodeValLevels {
		parent := parents[vl.val]

		var current *TreeNode
		if parent != nil {
			if parent.Left != nil && parent.Left.Val == vl.val {
				current = parent.Left
				parent.Left = nil
			} else {
				current = parent.Right
				parent.Right = nil
			}
		} else {
			current = root
		}
		if current.Left != nil {
			result = append(result, current.Left)
			current.Left = nil
		}
		if current.Right != nil {
			result = append(result, current.Right)
			current.Right = nil
		}

	}
	return result
}

func dfsGetParents(root *TreeNode, parent *TreeNode, parents *map[int]*TreeNode, level int, nodeValLevels *[]nodeValLevel, deleteMap *map[int]int) {
	if root == nil {
		return
	}
	(*parents)[root.Val] = parent
	if _, ok := (*deleteMap)[root.Val]; ok {
		*nodeValLevels = append(*nodeValLevels, nodeValLevel{root.Val, level})
	}
	dfsGetParents(root.Left, root, parents, level+1, nodeValLevels, deleteMap)
	dfsGetParents(root.Right, root, parents, level+1, nodeValLevels, deleteMap)
}

func run(slice, toDelete []int) [][]int {
	forest := delNodes(SliceToTreeNodeFullMinInt(slice), toDelete)
	res := make([][]int, len(forest))
	for i, node := range forest {
		res[i] = TreeNodeToSlice(node)
	}
	return res
}
