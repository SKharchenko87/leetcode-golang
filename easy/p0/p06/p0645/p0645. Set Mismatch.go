package p0645

// ToDo banchmark
func findErrorNums(nums []int) []int {
	arr := make([]int, len(nums))
	repetitionOfOneNumber := 0
	for _, v := range nums {
		arr[v-1]++
		if arr[v-1] == 2 {
			repetitionOfOneNumber = v
		}
	}
	for i, v := range arr {
		if v == 0 {
			return []int{repetitionOfOneNumber, i + 1}
		}

	}
	return []int{0, 0}
}

/*With sort
func findErrorNums(nums []int) []int {
	l := len(nums)
	sort.Ints(nums)
	fmt.Println(nums)
	repetitionOfOneNumber := 0
	lossOfAnother := 0
	koef := 1
	for i := 0; i < l && (repetitionOfOneNumber == 0 || lossOfAnother == 0); i++ {
		if lossOfAnother == 0 && nums[i] != i+koef {
			lossOfAnother = i + koef
		}
		if i < l-1 && nums[i] == nums[i+1] {
			repetitionOfOneNumber = nums[i]
			koef = 0
		}
	}
	if lossOfAnother == 0 {
		lossOfAnother = l
	}
	return []int{repetitionOfOneNumber, lossOfAnother}
}
*/
