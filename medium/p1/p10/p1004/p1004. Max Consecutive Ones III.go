package main

import "fmt"

func longestOnes(nums []int, k int) int {
	startWindow := 0
	//endWindow:=-1
	longest := 0
	counterZero := 0
	for i, v := range nums {
		if v == 0 {
			if counterZero == k {
				if longest < i-startWindow {
					longest = i - startWindow
				}
				if k == 0 {
					startWindow = startWindow
				} else {
					j := startWindow + 1
					for ; j < i; j++ {
						if nums[j] == 0 {
							break
						}
					}
					if j+1 < len(nums) {
						startWindow = j + 1
						if nums[j+1] == 1 {
							counterZero--
						}
					} else {
						startWindow = i
					}
				}

			} else {
				counterZero++
			}
		}
	}

	if longest < len(nums)-startWindow {
		longest = len(nums) - startWindow
	}
	return longest
}

func main() {
	//nums, k := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3
	//nums, k := []int{0, 0, 1, 1, 1, 0, 0}, 0
	//nums, k := []int{0, 0, 1, 1, 1, 0, 0}, 2
	nums, k := []int{0, 0, 0, 0}, 0
	//nums, k := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2
	//nums, k := []int{0, 0, 1, 1, 1}, 2
	fmt.Println(longestOnes(nums, k))
}

//      i
//0, 1, 0, 0, 0, 1, 1, 1, 0
//|     |
//t0    t1
//sw
//
//
//         i
//0, 1, 0, 0, 0, 1, 1, 1, 0
//   |  |  |
//      t0 t1
//   sw
//
//
//
//            i
//0, 1, 0, 0, 0, 1, 1, 1, 0
//         |  |
//         t0 t1
//         sw
//
//
//
//
//                        i
//0, 1, 0, 0, 0, 1, 1, 1, 0
//            |           |
//            t0          t1
//            sw
