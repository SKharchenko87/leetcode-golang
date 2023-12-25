package main

import "fmt"

//func searchInsert(nums []int, target int) int {
//	i := 0
//	for i < len(nums) {
//		if nums[i] >= target {
//			return i
//		}
//		i++
//	}
//	return i
//}

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	i := 0
	lenes := end - start
	for start <= end {
		i = start + lenes
		if nums[i] == target {
			return i
		} else if nums[i] < target {
			start = i + 1
		} else {
			end = i - 1
		}
		lenes = (end - start) / 2
	}
	if start > end {
		return start
	}
	return i
}

func main() {
	nums := []int{0, 1, 2, 3}
	target := -1
	//Output: 2
	fmt.Println(searchInsert(nums, target))

	//nums = []int{1, 3, 5, 6}
	//target = 2
	////Output: 1
	//fmt.Println(searchInsert(nums, target))
	//
	//nums = []int{1, 3, 5, 6}
	//target = 7
	////Output: 4
	//fmt.Println(searchInsert(nums, target))

}
