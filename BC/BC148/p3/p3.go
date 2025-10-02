package p3

import "fmt"

type Stack []interface{}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}
func (s *Stack) Pop() interface{} {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}

func (s *Stack) Top() interface{} {
	l := len(*s)
	v := (*s)[l-1]
	return v
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func longestSpecialPath(edges [][]int, nums []int) []int {
	n := len(nums)
	tree := make(map[int][][2]int)
	for _, edge := range edges {
		u, v, length := edge[0], edge[1], edge[2]
		tree[u] = append(tree[u], [2]int{v, length})
		tree[v] = append(tree[v], [2]int{u, length})
	}
	colors := make([]*Stack, 50001)
	sum := make([]int, n)
	var maxLength, minNodes int
	maxLength = -1
	var dfs func(cur, parent, index, indexBreak, length int)
	dfs = func(cur, parent, index, indexBreak, length int) {
		sum[index] = length
		newIndexBreak := indexBreak
		if colors[nums[cur]] == nil {
			colors[nums[cur]] = new(Stack)
		}
		color := colors[nums[cur]]
		if !color.Empty() {
			v := color.Top().(int)
			if v > indexBreak {
				newIndexBreak = v
				curLength, curNodes := sum[index-1]-sum[indexBreak+1], index-indexBreak-1
				if maxLength < curLength {
					maxLength = curLength
					minNodes = curNodes
				} else if maxLength == curLength && minNodes > curNodes {
					minNodes = curNodes
				}
			}
		}
		color.Push(index)

		if len(tree[cur]) > 1 || cur == 0 {
			for _, child := range tree[cur] {
				if child[0] != parent {
					dfs(child[0], cur, index+1, newIndexBreak, length+child[1])
				}
			}
		} else {
			curLength, curNodes := length-sum[newIndexBreak+1], index-newIndexBreak
			if maxLength < curLength {
				maxLength = curLength
				minNodes = curNodes
			} else if maxLength == curLength && minNodes > curNodes {
				minNodes = curNodes
			}
		}

		color.Pop()

	}
	dfs(0, -1, 0, -1, 0)
	return []int{maxLength, minNodes}
}

func longestSpecialPath4(edges [][]int, nums []int) []int {
	n := len(nums)
	tree := make(map[int][][2]int)
	for _, edge := range edges {
		u, v, length := edge[0], edge[1], edge[2]
		tree[u] = append(tree[u], [2]int{v, length})
		tree[v] = append(tree[v], [2]int{u, length})
	}
	colors := map[int]*Stack{}
	sum := make([]int, n)
	var maxLength, minNodes int
	maxLength = -1
	var dfs func(cur, parent, index, indexBreak, length int)
	dfs = func(cur, parent, index, indexBreak, length int) {
		sum[index] = length
		newIndexBreak := indexBreak
		if colors[nums[cur]] == nil {
			colors[nums[cur]] = new(Stack)
		}
		color := colors[nums[cur]]
		if !color.Empty() {
			v := color.Top().(int)
			if v > indexBreak {
				newIndexBreak = v
				curLength, curNodes := sum[index-1]-sum[indexBreak+1], index-indexBreak-1
				if maxLength < curLength {
					maxLength = curLength
					minNodes = curNodes
				} else if maxLength == curLength && minNodes > curNodes {
					minNodes = curNodes
				}
			}
		}
		color.Push(index)

		if len(tree[cur]) > 1 || cur == 0 {
			for _, child := range tree[cur] {
				if child[0] != parent {
					dfs(child[0], cur, index+1, newIndexBreak, length+child[1])
				}
			}
		} else {
			curLength, curNodes := length-sum[newIndexBreak+1], index-newIndexBreak
			if maxLength < curLength {
				maxLength = curLength
				minNodes = curNodes
			} else if maxLength == curLength && minNodes > curNodes {
				minNodes = curNodes
			}
		}

		color.Pop()

	}
	dfs(0, -1, 0, -1, 0)
	return []int{maxLength, minNodes}
}

// TLE
func longestSpecialPath3(edges [][]int, nums []int) []int {
	n := len(nums)
	path := make([][2]int, n)
	indexPath := 0
	calculatedPath := make([]bool, n)
	//indexCalculatedPath := 0
	tree := make(map[int][][2]int)
	for _, edge := range edges {
		u, v, length := edge[0], edge[1], edge[2]
		tree[u] = append(tree[u], [2]int{v, length})
		tree[v] = append(tree[v], [2]int{u, length})
	}

	var maxLength, minNodes int
	minNodes = 10001
	var dfs func(node, parent, length int)
	dfs = func(node, parent, length int) {
		path[indexPath][0] = node
		path[indexPath][1] = length
		if len(tree[node]) == 1 && node != 0 {
			//fmt.Println(path[:indexPath+1])
			computing(path, nums, indexPath, &maxLength, &minNodes, &calculatedPath)
			indexPath++
		} else {
			indexPath++
			for _, v := range tree[node] {
				if v[0] != parent {
					dfs(v[0], node, v[1])
				}
			}
		}
		indexPath--
	}
	dfs(0, -1, 0)

	return []int{maxLength, minNodes}
}

func computing(path [][2]int, nums []int, lastIndexPath int, maxLength, minNodes *int, calculatedPath *[]bool) {
	colors := make(map[int]bool, lastIndexPath)
	colors[nums[path[lastIndexPath][0]]] = true
	j := lastIndexPath
	i := lastIndexPath - 1
	sum := path[lastIndexPath][1]
	for ; i >= 0; i-- {
		colorI := nums[path[i][0]]
		if colors[colorI] {

			if *maxLength < sum-path[i+1][1] {
				*maxLength = sum - path[i+1][1]
				*minNodes = j - i
			} else if *maxLength == sum-path[i+1][1] && *minNodes > j-i {
				*minNodes = j - i
			}

			colorJ := nums[path[j][0]]
			for colorJ != colorI {
				colors[colorJ] = false
				sum -= path[j][1] // tree[path[j]][path[j-1]]
				j--
				if (*calculatedPath)[path[j][0]] {
					return
				}
				(*calculatedPath)[path[j][0]] = true
				colorJ = nums[path[j][0]]
			}

			sum -= path[j][1] // tree[path[j]][path[j-1]]
			j--
			if (*calculatedPath)[path[j][0]] {
				return
			}
			(*calculatedPath)[path[j][0]] = true
			sum += path[i][1] // tree[path[i]][path[i+1]]
			colors[colorI] = true

		} else {
			colors[colorI] = true
			sum += path[i][1] // tree[path[i]][path[i+1]]
		}
	}

	if *maxLength < sum {
		*maxLength = sum
		*minNodes = j - i
	} else if *maxLength == sum && *minNodes > j-i {
		*minNodes = j - i
	}
}

// TLE
func longestSpecialPath2(edges [][]int, nums []int) []int {
	n := len(nums)
	path := make([]int, n)
	indexPath := 0
	calculatedPath := make([]bool, n)
	//indexCalculatedPath := 0
	tree := make(map[int]map[int]int)
	for _, edge := range edges {
		u, v, length := edge[0], edge[1], edge[2]
		if _, ok := tree[u]; !ok {
			tree[u] = make(map[int]int)
		}
		if _, ok := tree[v]; !ok {
			tree[v] = make(map[int]int)
		}
		tree[u][v] = length
		tree[v][u] = length
	}
	var maxLength, minNodes int
	minNodes = 10001
	var dfs func(node, parent int)
	dfs = func(node, parent int) {
		path[indexPath] = node
		if len(tree[node]) == 1 && node != 0 {
			//fmt.Println(path[:indexPath+1])
			calculate(path, nums, indexPath, tree, &maxLength, &minNodes, &calculatedPath)
			indexPath++
		} else {
			indexPath++
			for v, _ := range tree[node] {
				if v != parent {
					dfs(v, node)
				}
			}
		}
		indexPath--
	}
	dfs(0, -1)

	return []int{maxLength, minNodes}
}

func calculate(path, nums []int, lastIndexPath int, tree map[int]map[int]int, maxLength, minNodes *int, calculatedPath *[]bool) {
	colors := make(map[int]bool, lastIndexPath)
	colors[nums[path[lastIndexPath]]] = true
	j := lastIndexPath
	i := lastIndexPath - 1
	sum := 0
	for ; i >= 0; i-- {
		colorI := nums[path[i]]
		if colors[colorI] {

			if *maxLength < sum {
				*maxLength = sum
				*minNodes = j - i
			} else if *maxLength == sum && *minNodes > j-i {
				*minNodes = j - i
			}

			colorJ := nums[path[j]]
			for colorJ != colorI {
				colors[colorJ] = false
				sum -= tree[path[j]][path[j-1]]
				j--
				if (*calculatedPath)[path[j]] {
					return
				}
				(*calculatedPath)[path[j]] = true
				colorJ = nums[path[j]]
			}

			sum -= tree[path[j]][path[j-1]]
			j--
			if (*calculatedPath)[path[j]] {
				return
			}
			(*calculatedPath)[path[j]] = true
			sum += tree[path[i]][path[i+1]]
			colors[colorI] = true

		} else {
			colors[colorI] = true
			sum += tree[path[i]][path[i+1]]
		}
	}

	if *maxLength < sum {
		*maxLength = sum
		*minNodes = j - i
	} else if *maxLength == sum && *minNodes > j-i {
		*minNodes = j - i
	}

}

func calculate1(path, nums []int, lastIndexPath int, tree map[int]map[int]int, maxLength, minNodes *int) (res [][3]int) {
	colors := make(map[int]bool, lastIndexPath)
	colors[nums[path[lastIndexPath]]] = true
	j := lastIndexPath
	i := lastIndexPath - 1
	sum := 0
	for ; i >= 0; i-- {
		colorI := nums[path[i]]
		if colors[colorI] {
			res = append(res, [3]int{i + 1, j, sum})
			if *maxLength < sum {
				*maxLength = sum
				*minNodes = j - i
			} else if *maxLength == sum && *minNodes > j-i {
				*minNodes = j - i
			}

			colorJ := nums[path[j]]
			for colorJ != colorI && j > 0 {
				colors[colorJ] = false
				sum -= tree[path[j]][path[j-1]]
				j--
			}
			if j > 0 {
				sum -= tree[path[j]][path[j-1]]
				j--
				sum += tree[path[i]][path[i+1]]
				colors[colorI] = true
			}
		} else {
			colors[colorI] = true
			sum += tree[path[i]][path[i+1]]
		}
	}
	res = append(res, [3]int{i + 1, j, sum})
	if *maxLength < sum {
		*maxLength = sum
		*minNodes = j - i
	} else if *maxLength == sum && *minNodes > j-i {
		*minNodes = j - i
	}
	fmt.Println(res)
	return res
}

func calculate0(path, nums []int, lastIndexPath int) (res [][2]int) {
	colors := make(map[int]bool, lastIndexPath)
	colors[nums[path[lastIndexPath]]] = true
	j := lastIndexPath
	i := lastIndexPath - 1
	for ; i >= 0; i-- {
		colorI := nums[path[i]]
		if colors[colorI] {
			res = append(res, [2]int{i + 1, j})
			colorJ := nums[path[j]]
			for colorJ != colorI {
				colors[colorJ] = false
				j--
				colorJ = nums[path[j]]
			}
			j--
		} else {
			colors[colorI] = true
		}
	}
	res = append(res, [2]int{i + 1, j})
	fmt.Println(res)
	return res
}

// TLE
func longestSpecialPath1(edges [][]int, nums []int) []int {
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

	//for _, nodes := range nodes[0].childrenNode {
	//	dfs(nodes, &Node{}, 0, 0, make(map[int]bool))
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
