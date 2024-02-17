package p3038

// toDo banchmark
func maxOperations(nums []int) int {
	cnt := 1 // по условию длина nums массива всегда >= 2
	for i := 2; i < len(nums)-1; i += 2 {
		if nums[i+1]+nums[i] != nums[i-1]+nums[i-2] {
			break
		}
		cnt++
	}
	return cnt
}

func maxOperations1(nums []int) int {
	cnt := 1
	nums[1] += nums[0]
	for i := 2; i < len(nums)-1; i += 2 {
		nums[i+1] += nums[i]
		if nums[i+1] != nums[i-1] {
			break
		}
		cnt++
	}
	return cnt
}
