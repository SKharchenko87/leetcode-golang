package main

import "fmt"

func maxOperations(nums []int, k int) int {
	m := map[int]int{}
	counter := 0
	for _, v := range nums {
		if _, ok := m[k-v]; ok {
			counter++
			m[k-v]--
			if m[k-v] == 0 {
				delete(m, k-v)
			}
		} else {
			m[v]++
		}
	}
	return counter
}

func main() {
	nums, k := []int{1, 2, 3, 4}, 5
	//nums, k := []int{3, 1, 3, 4, 3}, 6
	//nums, k := []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}, 3
	fmt.Println(maxOperations(nums, k))
}
