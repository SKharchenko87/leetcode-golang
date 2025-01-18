package p3

// TLE
func longestSpecialPath(edges [][]int, nums []int) []int {
	tree := make(map[int][][2]int)
	for _, edge := range edges {
		u, v, length := edge[0], edge[1], edge[2]
		tree[u] = append(tree[u], [2]int{v, length})
		tree[v] = append(tree[v], [2]int{u, length})
	}

	res := []int{-1, 10001}
	started := make(map[int]bool, len(nums))
	started[0] = true
	var dfs func(node, parent int, length, cnt int, uniq map[int]bool)
	dfs = func(node, parent int, length, cnt int, uniq map[int]bool) {
		color := nums[node]
		if uniq[color] {
			return
		}
		uniq[color] = true
		cnt++
		if length > res[0] {
			res[0] = length
			res[1] = cnt
		} else if length == res[0] && cnt < res[1] {
			res[1] = cnt
		}

		for _, data := range tree[node] {
			if parent != data[0] {
				dfs(data[0], node, length+data[1], cnt, uniq)
				if !started[data[0]] {
					dfs(data[0], node, 0, 0, make(map[int]bool))
					started[data[0]] = true
				}
			}
		}
		delete(uniq, color)
	}
	dfs(0, -1, 0, 0, make(map[int]bool))
	return res
}

type Node struct {
	index          int
	color          int
	childrenNode   []*Node
	childrenLength []int
}

// TLE
func longestSpecialPath0(edges [][]int, nums []int) []int {
	n := len(nums)
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{index: i, color: nums[i]}
	}
	for _, edge := range edges {
		p := min(edge[0], edge[1])
		c := max(edge[0], edge[1])

		if nodes[p].childrenNode == nil {
			nodes[p].childrenNode = []*Node{}
			nodes[p].childrenLength = []int{}
		}
		nodes[p].childrenNode = append(nodes[p].childrenNode, nodes[c])
		nodes[p].childrenLength = append(nodes[p].childrenLength, edge[2])

		if nodes[c].childrenNode == nil {
			nodes[c].childrenNode = []*Node{}
			nodes[c].childrenLength = []int{}
		}
		nodes[c].childrenNode = append(nodes[c].childrenNode, nodes[p])
		nodes[c].childrenLength = append(nodes[c].childrenLength, edge[2])

	}
	res := []int{-1, 10001}
	started := map[*Node]bool{nodes[0]: true}
	var dfs func(node, parent *Node, length, cnt int, uniq map[int]bool)
	dfs = func(node, parent *Node, length, cnt int, uniq map[int]bool) {
		if uniq[node.color] {
			return
		}
		uniq[node.color] = true
		cnt++
		if length > res[0] {
			res[0] = length
			res[1] = cnt
		} else if length == res[0] && cnt < res[1] {
			res[1] = cnt
		}

		for i, child := range node.childrenNode {
			if parent != child {
				dfs(child, node, length+node.childrenLength[i], cnt, uniq)
				if !started[child] {
					dfs(child, node, 0, 0, make(map[int]bool))
					started[child] = true
				}
			}
		}
		delete(uniq, node.color)
	}

	//for _, node := range nodes[0].childrenNode {
	//	dfs(node, &Node{}, 0, 0, make(map[int]bool))
	//}
	dfs(nodes[0], &Node{}, 0, 0, make(map[int]bool))
	return res
}

type Ints interface {
	int | int8 | int16 | int32 | int64
}
type Floats interface {
	float32 | float64
}
type Numeric interface{ Ints | Floats }

func abs[T Numeric](x T) T {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min[T Numeric](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func max[T Numeric](x, y T) T {
	if x > y {
		return x
	}
	return y
}
