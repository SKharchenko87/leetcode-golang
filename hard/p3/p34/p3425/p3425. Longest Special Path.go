package p3425

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
