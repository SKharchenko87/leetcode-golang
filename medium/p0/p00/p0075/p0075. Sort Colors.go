package p0075

func sortColors(nums []int) {
	i, lastIndex0, firstIndex2 := 0, -1, len(nums)
	for i < firstIndex2 {
		if nums[i] == 0 {
			lastIndex0++
			nums[i], nums[lastIndex0] = nums[lastIndex0], nums[i]
		} else if nums[i] == 2 {
			firstIndex2--
			nums[i], nums[firstIndex2] = nums[firstIndex2], nums[i]
			continue
		}
		i++
	}
}

func sortColors1(nums []int) {
	lastIndex0 := -1
	firstIndex2 := len(nums)
	i := 0
	for i < firstIndex2 {
		switch nums[i] {
		case 0:
			lastIndex0++
			nums[i], nums[lastIndex0] = nums[lastIndex0], nums[i]
		case 2:
			firstIndex2--
			nums[i], nums[firstIndex2] = nums[firstIndex2], nums[i]
			continue
		}
		i++
	}
}

func sortColors0(nums []int) {
	cntArr := [3]int{}
	for _, num := range nums {
		cntArr[num]++
	}
	index := 0
	for i, cnt := range cntArr {
		for j := 0; j < cnt; j++ {
			nums[index] = i
			index++
		}
	}
}

func run(nums []int) []int {
	sortColors(nums)
	return nums
}
