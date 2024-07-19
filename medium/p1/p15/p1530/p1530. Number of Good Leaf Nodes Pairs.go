package p1530

import (
	. "leetcode/stucture"
)

// #1: l = ?  r = ?
// #2: l = 1 0 0 0 0 r = ?
// #4: не заходим
// #5: l = 1 0 0 0 0 r = 1 0 0 0 0
// #2: l = 1 0 0 0 0 r = 0 2 0 0 0
// #1: l = 0 1 2 0 0 r = ?
// #3: l = 1 0 0 0 0 r = 1 0 0 0 0
// #1: l = 0 1 2 0 0 r = 0 2 0 0 0
// => 0 0 3 2 0

func countPairs(root *TreeNode, distance int) int {
	var dfs func(node *TreeNode, distance int) [12]int
	dfs = func(node *TreeNode, distance int) [12]int {
		if node == nil {
			return [12]int{}
		} else if node.Left == nil && node.Right == nil {
			current := [12]int{}
			current[0] = 1
			return current
		}
		left := dfs(node.Left, distance)
		right := dfs(node.Right, distance)
		current := [12]int{}
		for i := 0; i < 10; i++ {
			current[i+1] += left[i] + right[i]
		}
		current[11] = left[11] + right[11]
		for i := 0; i <= distance; i++ {
			for j := 0; j <= distance; j++ {
				if 2+i+j <= distance {
					current[11] += left[i] * right[j]
				}
			}
		}
		return current
	}
	return dfs(root, distance)[11]
}

// Approach 1: Graph Conversion + BFS
func countPairs2(root *TreeNode, distance int) int {
	graph := make(map[*TreeNode][]*TreeNode)
	leafs := make(map[*TreeNode]bool)
	var dfs func(node *TreeNode, parent *TreeNode)
	dfs = func(node *TreeNode, parent *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			leafs[node] = true
		}
		if parent != nil {
			if _, ok := graph[parent]; !ok {
				graph[parent] = make([]*TreeNode, 0)
			}
			graph[parent] = append(graph[parent], node)

			if _, ok := graph[node]; !ok {
				graph[node] = make([]*TreeNode, 0)
			}
			graph[node] = append(graph[node], parent)

		}
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)
	res := 0

	for leaf, _ := range leafs {
		bfsQueue := make([]*TreeNode, 0)
		seen := make(map[*TreeNode]bool)
		bfsQueue = append(bfsQueue, leaf)
		seen[leaf] = true
		for i := 0; i <= distance; i++ {
			size := len(bfsQueue)
			for j := 0; j < size; j++ {
				currNode := bfsQueue[0]
				bfsQueue = bfsQueue[1:]
				if _, ok := leafs[currNode]; ok && currNode != leaf {
					res++
				}
				for _, neighbor := range graph[currNode] {
					if !seen[neighbor] {
						bfsQueue = append(bfsQueue, neighbor)
						seen[neighbor] = true
					}
				}
			}
		}
	}
	return res / 2
}

type mers struct {
	parent *TreeNode
	left   *TreeNode
	right  *TreeNode
}

type neighbor struct {
	node     *TreeNode
	distance int
}

func countPairs1(root *TreeNode, distance int) int {
	graph := make(map[*TreeNode]mers)
	leafs := make(map[*TreeNode]bool)
	var dfs func(node *TreeNode, parent *TreeNode)
	dfs = func(node *TreeNode, parent *TreeNode) {
		//if node == nil {
		//	return
		//}
		if node.Left == nil && node.Right == nil {
			leafs[node] = true
		}
		if _, ok := graph[node]; !ok {
			graph[node] = mers{}
		}
		v := graph[node]
		v.parent = parent
		v.right = node.Right
		v.left = node.Left
		graph[node] = v
		if node.Right != nil {
			dfs(node.Right, node)
		}
		if node.Left != nil {
			dfs(node.Left, node)
		}
	}
	dfs(root, nil)
	res := 0
	for leaf, _ := range leafs {
		path := map[*TreeNode]bool{}
		path[leaf] = true
		//currNode := leaf
		neighbors := []neighbor{neighbor{leaf, 0}}

		for j := 0; j < len(neighbors) && neighbors[j].distance < distance; j++ {
			currNode := neighbors[j].node
			i := neighbors[j].distance
			if graph[currNode].parent != nil && !path[graph[currNode].parent] {
				if _, ok := leafs[graph[currNode].parent]; ok {
					res++
				}
				neighbors = append(neighbors, neighbor{graph[currNode].parent, i + 1})
				path[graph[currNode].parent] = true
			}

			if graph[currNode].left != nil && !path[graph[currNode].left] {
				if _, ok := leafs[graph[currNode].left]; ok {
					res++
				}
				neighbors = append(neighbors, neighbor{graph[currNode].left, i + 1})
				path[graph[currNode].left] = true
			}

			if graph[currNode].right != nil && !path[graph[currNode].right] {
				if _, ok := leafs[graph[currNode].right]; ok {
					res++
				}
				neighbors = append(neighbors, neighbor{graph[currNode].right, i + 1})
				path[graph[currNode].right] = true
			}
		}

	}
	return res / 2
}

func countPairs0(root *TreeNode, distance int) int {
	leafs := [][]byte{}
	path := make([]byte, 1024)
	var dfs func(node *TreeNode, direction byte, level int)
	dfs = func(node *TreeNode, direction byte, level int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			y := make([]byte, level)
			copy(y, path[:level])
			leafs = append(leafs, y)
		}
		path[level] = 'L'
		dfs(node.Left, 'L', level+1)
		path[level] = 'R'
		dfs(node.Right, 'R', level+1)
	}
	dfs(root, ' ', 0)
	res := 0
	for i, leaf := range leafs {
		//println(string(leaf))
		for j := i + 1; j < len(leafs); j++ {
			levelI, levelJ := len(leaf), len(leafs[j])
			levels := levelI + levelJ
			minLevel := min(levelI, levelJ)
			k := 0
			for ; k < minLevel && leaf[k] == leafs[j][k]; k++ {
			}
			if levels-2*k <= distance {
				res++
			}
		}
	}
	return res
}

func run(slice []int, distance int) int {
	return countPairs(SliceToTreeNodeFullMinInt(slice), distance)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
