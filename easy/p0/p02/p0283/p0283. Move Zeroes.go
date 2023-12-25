package main

import "fmt"

func moveZeroes(nums []int) {
	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[c], nums[i] = nums[i], nums[c]
			c++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

// 0,1,0,3,12
// |
//c,i

// 0,1,0,3,12
// | |
// c i

// 1,0,3,0,12
//   |
//  c,i

// 1,0,3,0,12
//   | |
//   c i

// 1,3,0,0,12
//     |
//    c,i

// 1,3,0,0,12
//     | |
//     c i

// 1,3,0,0,12
//     |   |
//     c   i

// 1,3,12,0,0
//        | |
//        c i
