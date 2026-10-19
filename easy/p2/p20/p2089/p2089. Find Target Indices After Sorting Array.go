package p2089

func targetIndices(nums []int, target int) []int {
	countTarget := 0
	countLessTarget := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < target {
			countLessTarget++
		} else if nums[i] == target {
			countTarget++
		}
	}
	res := make([]int, countTarget)
	for i := 0; i < countTarget; i++ {
		res[i] = countLessTarget + i
	}
	return res
}
