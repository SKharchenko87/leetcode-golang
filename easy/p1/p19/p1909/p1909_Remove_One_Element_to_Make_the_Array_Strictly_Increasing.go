package main

import "fmt"

func canBeIncreasing(nums []int) bool {
	if len(nums) == 2 {
		return true
	}
	flg := nums[len(nums)-2] >= nums[len(nums)-1]
	for i := len(nums) - 3; i >= 1; i-- {
		if nums[i] >= nums[i+1] {
			if flg {
				return false
			} else {
				flg = true
			}
			if !(nums[i-1] < nums[i+1] || (i != len(nums)-2 && nums[i] < nums[i+2])) {
				return false
			}
		}
	}
	if flg && !(nums[0] < nums[1]) {
		return false
	} else {
		return true
	}
}

func main() {
	nums := []int{105, 924, 32, 968}
	//nums := []int{1, 2, 10, 5, 7}
	//nums := []int{2, 3, 1, 2}
	//nums := []int{1, 1, 1}
	//nums := []int{2, 3, 4, 5, 1, 5}
	//nums := []int{100, 21, 100}
	//nums := []int{512, 867, 904, 997, 403}
	fmt.Println(canBeIncreasing(nums))
}
