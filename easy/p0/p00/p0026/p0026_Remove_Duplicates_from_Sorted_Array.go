package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	curIdx := 0
	for i := 1; i < len(nums); i++ {
		if nums[curIdx] < nums[i] {
			curIdx++
			nums[curIdx], nums[i] = nums[i], nums[curIdx]
		}
	}
	curIdx++
	return curIdx
}

func main() {
	m := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//x := []int{0, math.MaxInt, 1, math.MaxInt, math.MaxInt, 2, math.MaxInt, 3, math.MaxInt, 4}

	fmt.Println(removeDuplicates(m))
	fmt.Println(m)

}
