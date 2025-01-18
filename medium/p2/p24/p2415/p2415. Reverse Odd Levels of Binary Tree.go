package p2415

import . "leetcode/stucture"

func reverseOddLevels(root *TreeNode) *TreeNode {
	query := []*TreeNode{root}
	index := 0
	for ; index < len(query); index++ {
		if query[index].Left != nil {
			query = append(query, query[index].Left)
		}
		if query[index].Right != nil {
			query = append(query, query[index].Right)
		}
	}
	start := 1
	for i := 1; 1<<i < len(query); i++ {
		if i%2 == 1 {
			for j := 0; j < 1<<(i-1); j++ {
				query[start+1<<i-1-j].Val, query[start+j].Val = query[start+j].Val, query[start+1<<i-1-j].Val
			}
		}
		start += 1 << i
	}
	return root
}

/*
0:0-0
1:1-2    +1
2:3-6    +2
3:7-14   +4
4:15-30  +8
5:31     +16
*/

func run(root []int) []int {
	return TreeNodeToSlice(reverseOddLevels(SliceToTreeNodeFullMinInt(root)))
}
