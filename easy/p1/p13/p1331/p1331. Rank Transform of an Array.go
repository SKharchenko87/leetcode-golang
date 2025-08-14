package p1331

import (
	"slices"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	l := len(arr)
	if l == 0 {
		return []int{}
	}
	orderSet := make([]int, 0, l)
	for i := 0; i < l; i++ {
		index, flg := slices.BinarySearch(orderSet, arr[i])
		if !flg {
			orderSet = slices.Insert(orderSet, index, arr[i])
		}
	}
	rank := make([]int, l)
	for i := 0; i < l; i++ {
		index, _ := slices.BinarySearch(orderSet, arr[i])
		rank[i] = index + 1
	}
	return rank
}

func arrayRankTransform1(arr []int) []int {
	l := len(arr)
	if l == 0 {
		return []int{}
	}
	arrCopy := make([]int, l)
	copy(arrCopy, arr)
	sort.Ints(arrCopy)
	rankMap := make(map[int]int, l)
	prev := arrCopy[0]
	currRank := 1
	rankMap[arrCopy[0]] = currRank
	for i := 1; i < l; i++ {
		if prev != arrCopy[i] {
			prev = arrCopy[i]
			currRank++
			rankMap[arrCopy[i]] = currRank
		}
	}
	rank := make([]int, l)
	for i := 0; i < l; i++ {
		rank[i] = rankMap[arr[i]]
	}
	return rank
}

type forSort struct {
	n1 []int
	n2 []int
}

func (fs *forSort) Len() int {
	return len(fs.n1)
}
func (fs *forSort) Swap(i, j int) {
	fs.n1[i], fs.n1[j] = fs.n1[j], fs.n1[i]
	fs.n2[i], fs.n2[j] = fs.n2[j], fs.n2[i]
}
func (fs *forSort) Less(i, j int) bool {
	return fs.n1[i] <= fs.n1[j]
}

func arrayRankTransform0(arr []int) []int {
	l := len(arr)
	if l == 0 {
		return []int{}
	}
	indexes := make([]int, l)
	for i := 0; i < l; i++ {
		indexes[i] = i
	}

	fs := forSort{arr, indexes}
	sort.Sort(&fs)

	prev := arr[0]
	arr[0] = 1
	for i := 1; i < l; i++ {
		if prev == arr[i] {
			arr[i] = arr[i-1]
		} else {
			prev = arr[i]
			arr[i] = arr[i-1] + 1
		}
	}

	fs = forSort{indexes, arr}
	sort.Sort(&fs)

	return arr
}
