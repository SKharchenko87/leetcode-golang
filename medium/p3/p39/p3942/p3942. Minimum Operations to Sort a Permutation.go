package p3942

func minOperations(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	// Ищем положение нуля и проверяем что отсартировано
	first, last := -1, -1
	for i := 0; i < n; i++ {
		switch nums[i] {
		case 0:
			first = i
		case n - 1:
			last = i
		default:
			if abs(nums[(n+i-1)%n]-nums[i]) != 1 || abs(nums[i]-nums[(i+1)%n]) != 1 {
				return -1
			}
		}
	}
	if first == 0 && last == n-1 {
		return 0
	}
	res := 0
	if nums[(first+1)%n]-nums[first] == 1 {
		res = min(first, n-first+2)
	} else {
		res = 1 + min(n-first-1, first+1)
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
