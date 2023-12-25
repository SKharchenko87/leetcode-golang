package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

`func search(nums []int, target int) int {
	leftId := 0
	rightId := len(nums) - 1
	if target == nums[rightId] {
		return rightId
	}
	mid := leftId + (rightId-leftId)/2
	for leftId <= rightId {
		cur := nums[mid]
		if cur == target {
			return mid
		}
		if mid+1 < len(nums) && nums[mid+1] == target {
			return mid + 1
		}
		left := nums[leftId]
		right := nums[rightId]
		if target < cur {
			if left < cur {
				if left <= target {
					leftId = leftId
					rightId = mid - 1
				} else {
					leftId = mid + 1
					rightId = rightId
				}
			} else {
				leftId = leftId
				rightId = mid - 1
			}
		} else {
			if cur < right {
				if target <= right {
					leftId = mid + 1
					rightId = rightId
				} else {
					leftId = leftId
					rightId = mid - 1
				}
			} else {
				leftId = mid + 1
				rightId = rightId
			}
		}

		mid = leftId + (rightId-leftId)/2
	}
	return -1
}`

func main() {
	//nums := []int{1, 3}
	//nums := []int{4, 5, 6, 7, 8, 1, 2, 3}
	//nums := []int{3, 1}
	nums := []int{5, 1, 2, 3, 4}
	//nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(search(nums, 1))
}
