package main

import "fmt"

func removeElement(nums []int, val int) int {
	//if len(nums) == 0 {
	//	return 0
	//}
	c := 0
	i := 0
	for i+c <= len(nums)-1 {
		if nums[i] == val {
			for j := i; j < len(nums)-1-c; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			c++
			i--
		}
		i++
	}
	return i
}

func main() {
	nums := []int{0, 3, 1, 3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
}

//{0, 3, 1, 3, 2, 2, 3}
//{0, 3, 1, 2, 3, 2, 3}
//{0, 3, 1, 2, 2, 3, 3}
//{0, 1, 3, 2, 2, 3, 3}
//{0, 1, 2, 3, 2, 3, 3}
//{0, 1, 2, 2, 3, 3, 3}

//{0, 3, 1, 3, 2, 2, 3}
//{0, 1, 3, 3, 2, 2, 3}
//{0, 1, 2, 3, 3, 2, 3}
//{0, 1, 2, 3, 2, 3, 3}
//{0, 1, 2, 2, 3, 3, 3}
