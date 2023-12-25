package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	//word1, word2 := "cabbba", "abbccc"
	//word1, word2 := "abc", "bca"
	word1, word2 := "uau", "ssx"
	fmt.Println(closeStrings(word1, word2))
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	arr1 := make([]int, 26)
	arr2 := make([]int, 26)
	m := make(map[uint8]int, 26)
	for i, v := range word1 {
		arr1[v-'a']++
		m[word1[i]]++
		arr2[word2[i]-'a']++
	}
	if !slices.EqualFunc(arr1, arr2, func(v1, v2 int) bool { return ((v1 == 0) == (v2 == 0)) }) {
		return false
	}
	//for i, v := range arr2 {
	//	if _, ok := m[uint8(i)+'a']; ok && v > 0 {
	//		delete(m, uint8(i)+'a')
	//	}
	//}
	//if len(m) != 0 {
	//	return false
	//}
	sort.Ints(arr1)
	sort.Ints(arr2)
	return slices.Equal(arr1, arr2)
}

func stringToMap(word string, c chan map[int32]int) {
	m := make(map[int32]int)
	for _, v := range word {
		m[v]++
	}
	c <- m
}
