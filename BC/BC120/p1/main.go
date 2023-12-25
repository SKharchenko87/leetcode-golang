package main

import "fmt"

func main() {
	//nums := []int{8, 7, 6, 6}
	nums := []int{3, 7, 10, 6}
	//nums := []int{1, 2, 3, 4}
	//nums := []int{6, 5, 7, 8}
	//nums := []int{3, 5, 3, 5}
	fmt.Println(incremovableSubarrayCount(nums))
}
func incremovableSubarrayCount(nums []int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		j := i
		for ; j < len(nums); j++ {
			k := 1
			p := 0
			for ; k < len(nums); k++ {
				if i <= k && k <= j {
					k = j + 1
				}
				if i <= p && p <= j {
					p = j + 1
					k = p + 1
				}

				if p < 0 || p >= len(nums) || k >= len(nums) || nums[p] >= nums[k] {
					break
				}
				p++
			}
			if k >= len(nums) {
				fmt.Println(i, j)
				cnt++
			}
		}
	}
	return cnt
}
