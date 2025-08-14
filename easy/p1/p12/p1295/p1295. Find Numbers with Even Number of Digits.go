package p1295

func findNumbers(nums []int) int {
	count := 0
	for _, num := range nums {
		if getCountDigit(num)%2 == 0 {
			count++
		}
	}
	return count
}

func getCountDigit(n int) int {
	if n == 0 {
		return 1
	}
	res := 0
	for n > 0 {
		res++
		n /= 10
	}
	return res
}
