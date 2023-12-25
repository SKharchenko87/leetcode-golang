package main

import "fmt"

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	start := 0
	end := len(nums) - 1
	current := start + (end-start)/2
	for start <= end {
		if nums[current] == target {
			b := current
			bflg := false
			e := current
			eflg := false
			for i := 1; i < len(nums); i++ {
				if current-i == 0 {
					bflg = true
				}
				if current-i >= 0 && nums[current-i] == target {
					b = current - i
				} else {
					bflg = true
				}

				if current+i == len(nums)-1 {
					eflg = true
				}
				if current+i <= len(nums)-1 && nums[current+i] == target {
					e = current + i
				} else {
					eflg = true
				}

				if bflg && eflg {
					return []int{b, e}
				}
			}
			return []int{current, current}
		} else if nums[current] < target {
			start = current + 1
		} else {
			end = current - 1
		}
		current = start + (end-start)/2
	}
	return res
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	//nums := []int{2, 2}
	fmt.Println(searchRange(nums, 8))
}
