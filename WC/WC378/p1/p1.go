package p1

func hasTrailingZeros(nums []int) bool {
	cnt := 0
	for _, v := range nums {
		if v%2 == 0 {
			cnt++
			if cnt > 1 {
				return true
			}
		}
	}
	return false
}
