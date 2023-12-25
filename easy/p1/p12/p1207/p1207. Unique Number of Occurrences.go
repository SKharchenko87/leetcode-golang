package main

import "fmt"

func main() {
	//arr := []int{1, 2, 2, 1, 1, 3}
	arr := []int{1, 2}
	fmt.Println(uniqueOccurrences(arr))
}

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	q := make(map[int]int)
	for _, v := range m {
		q[v]++
		if q[v] == 2 {
			return false
		}
	}
	return true
}
