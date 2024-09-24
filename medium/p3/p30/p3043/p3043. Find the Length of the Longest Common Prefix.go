package p3043

import (
	"slices"
)

type NodeTreeMap struct {
	m map[int]*NodeTreeMap
}

type NodeTree [10]*NodeTree

var pow9 = []int{0, 9, 99, 999, 9_999, 99_999, 999_999, 9_999_999, 99_999_999, 999_999_999}

func longestCommonPrefix3(arr1 []int, arr2 []int) int {
	tree := &NodeTreeMap{}
	for _, i := range arr1 {
		parseMap(tree, i)
	}
	maxPrefix := 0
	for _, i := range arr2 {
		cur := tree
		candidateMaxPrefix := 0
		index, _ := slices.BinarySearch(pow9, i)
		for j := index; j > 0; j-- {
			x := i / (pow9[j-1] + 1)
			if cur.m[x] == nil {
				break
			} else {
				candidateMaxPrefix++
				cur = cur.m[x]
			}
			i %= (pow9[j-1] + 1)
		}
		maxPrefix = max(maxPrefix, candidateMaxPrefix)
	}
	return maxPrefix
}

func parseMap(cur *NodeTreeMap, i int) {
	index, _ := slices.BinarySearch(pow9, i)
	for j := index; j > 0; j-- {
		x := i / (pow9[j-1] + 1)
		if cur == nil {
			cur = &NodeTreeMap{}
		}
		if cur.m == nil {
			cur.m = make(map[int]*NodeTreeMap)
		}
		if _, ok := cur.m[x]; !ok {
			cur.m[x] = &NodeTreeMap{}
		}

		//if  cur.m[x] == nil {
		//	cur.m[x] = &NodeTreeMap{}
		//}
		cur = cur.m[x]
		i %= (pow9[j-1] + 1)
	}
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	m := make(map[int]struct{})
	for i := 0; i < len(arr1); i++ {
		v := arr1[i]
		for j := 0; v > 0; j++ {
			if _, ok := m[v]; ok {
				break
			}
			m[v] = struct{}{}
			v /= 10
		}
	}
	maxPrefix := 0
	for i := 0; i < len(arr2); i++ {
		v := arr2[i]
		index, _ := slices.BinarySearch(pow9, v)
		for index > maxPrefix {
			if _, ok := m[v]; ok {
				maxPrefix = index
				break
			}
			index--
			v /= 10
		}
	}
	return maxPrefix
}

func longestCommonPrefix1(arr1 []int, arr2 []int) int {
	m := make(map[int]int)
	for i := 0; i < len(arr1); i++ {
		v := arr1[i]
		l, _ := slices.BinarySearch(pow9, v)
		for j := 0; v > 0; j++ {
			if _, ok := m[v]; ok {
				break
			}
			m[v] = l
			v /= 10
			l--
		}
	}
	maxPrefix := 0
	for i := 0; i < len(arr2); i++ {
		v := arr2[i]
		for v > 0 {
			if l, ok := m[v]; ok && l > maxPrefix {
				maxPrefix = l
				break
			}
			v /= 10
		}
	}
	return maxPrefix
}

func longestCommonPrefix0(arr1 []int, arr2 []int) int {
	tree := &NodeTree{}
	for _, i := range arr1 {
		parse(tree, i)
	}
	maxPrefix := 0
	for _, i := range arr2 {
		cur := tree
		candidateMaxPrefix := 0
		index, _ := slices.BinarySearch(pow9, i)
		for j := index; j > 0; j-- {
			x := i / (pow9[j-1] + 1)
			if cur[x] == nil {
				break
			} else {
				candidateMaxPrefix++
				cur = cur[x]
			}
			i %= (pow9[j-1] + 1)
		}
		maxPrefix = max(maxPrefix, candidateMaxPrefix)
	}
	return maxPrefix
}

func parse(cur *NodeTree, i int) {
	index, _ := slices.BinarySearch(pow9, i)
	for j := index; j > 0; j-- {
		x := i / (pow9[j-1] + 1)
		if cur[x] == nil {
			cur[x] = &NodeTree{}
		}
		cur = cur[x]
		i %= (pow9[j-1] + 1)
	}
}

func parse0(cur *NodeTree, i int) {
	for i > 0 {
		defer func(i int) {
			if cur[i%10] == nil {
				cur[i%10] = &NodeTree{}
			}
			cur = cur[i%10]
		}(i)
		i /= 10
	}
}
