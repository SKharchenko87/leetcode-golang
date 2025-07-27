package p2210

func countHillValley(nums []int) int {
	prevIndex := 0
	result := 0
	for index := 1; index < len(nums)-1; index++ {
		if nums[index] == nums[index+1] {
			continue
		}
		if nums[prevIndex] < nums[index] && nums[index] > nums[index+1] ||
			nums[prevIndex] > nums[index] && nums[index] < nums[index+1] {
			result++
		}
		prevIndex = index
	}
	return result
}

func countHillValley0(nums []int) int {
	index := 1
	for ; index < len(nums) && nums[index-1] == nums[index]; index++ {
	}
	prevIndex := index - 1

	result := 0
	for ; index < len(nums)-1; index++ {
		if nums[index] == nums[index+1] {
			continue
		}
		if nums[prevIndex] < nums[index] && nums[index] > nums[index+1] ||
			nums[prevIndex] > nums[index] && nums[index] < nums[index+1] {
			result++
		}
		prevIndex = index
	}
	return result
}
