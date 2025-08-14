package p0977

func sortedSquares(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	indexLeft, indexRight := 0, l-1
	l--
	for indexLeft <= indexRight {
		tmpLeft, tmpRight := nums[indexLeft]*nums[indexLeft], nums[indexRight]*nums[indexRight]
		if tmpLeft > tmpRight {
			res[l] = tmpLeft
			indexLeft++
		} else {
			res[l] = tmpRight
			indexRight--
		}
		l--
	}
	return res
}
