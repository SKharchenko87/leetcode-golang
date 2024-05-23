package p2597

import (
	"fmt"
	"sort"
)

// ToDo разобраться в условии задачи
// Рабочий вариант
func beautifulSubsets(nums []int, k int) int {
	visited := make(map[int]int)
	return dfs(0, &nums, visited, k) - 1
}

func dfs(idx int, nums *[]int, visited map[int]int, k int) int {
	if idx == len(*nums) {
		return 1
	}
	num := (*nums)[idx]
	answer := 0
	if visited[num+k] == 0 && visited[num-k] == 0 {
		visited[num] += 1
		answer += dfs(idx+1, nums, visited, k)
		visited[num] -= 1
	}
	answer += dfs(idx+1, nums, visited, k)
	return answer
}

// Не рабочий вариант
func beautifulSubsets0(nums []int, k int) (res int) {
	sort.Ints(nums)
	fmt.Println(nums)
	l := len(nums)
	var recursive func(prev, prevIndex int)
	recursive = func(prev, prevIndex int) {
		index := prevIndex + 1
		if index == l {
			return
		}
		if prev-nums[index] != k {
			res++
			recursive(nums[index], index)
			println("Yes", prev, nums[index])
		}
		//else {
		//	println("No", prev, nums[index])
		//}
		recursive(prev, index)
	}
	for i, n := range nums {
		println("Yes", n)
		res++
		recursive(n, i)
	}
	return res
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}
