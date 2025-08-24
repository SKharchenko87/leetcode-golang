package p1493

import "fmt"

func longestSubarray(nums []int) int {
	firstIndex0, secondIndex0 := 1, -1
	result := 0
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			result = max(result, i-firstIndex0)
			firstIndex0, secondIndex0 = secondIndex0+2, i
		}
	}
	result = max(result, i-firstIndex0)
	return result
}

func longestSubarray1(nums []int) int {
	firstIndex0, secondIndex0 := -1, -1
	result := 0
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			result = max(result, i-firstIndex0-2)
			firstIndex0, secondIndex0 = secondIndex0, i
		}
	}
	result = max(result, i-firstIndex0-2)
	return result
}

func longestSubarray0(nums []int) int {
	prevLength := 0
	cursorZero := -1
	longest := 0
	for i, v := range nums {
		if v == 0 {
			length := i - cursorZero - 1
			if longest < prevLength+length {
				longest = prevLength + length
			}
			prevLength = length
			cursorZero = i
		}
	}
	length := len(nums) - cursorZero - 1
	if longest < prevLength+length {
		longest = prevLength + length
	}
	if cursorZero == -1 {
		longest = longest - 1
	}
	return longest
}

func main() {
	//nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	nums := []int{1, 1, 1}
	//nums := []int{1, 1, 0, 1}
	//nums := []int{1, 1, 0, 0, 1, 1, 1, 0, 1}
	fmt.Println(longestSubarray0(nums))
}
