package p3432

//go:noinline
func countPartitions(nums []int) int {
	sum := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		sum = (sum + nums[i]) % 2
	}
	if sum == 0 {
		return l - 1
	}
	return 0
}

//go:noinline
func countPartitions2(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum = (sum + num) % 2
	}
	if sum == 0 {
		return len(nums) - 1
	}
	return 0
}

//go:noinline
func countPartitions1(nums []int) int {
	sum := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		sum ^= nums[i] & 1
	}
	if sum == 0 {
		return l - 1
	}
	return 0
}

//go:noinline
func countPartitions0(nums []int) int {
	evenSum := 0
	oddSum := 0
	sum := 0
	for _, num := range nums {
		sum = (sum + num) % 2
		if sum == 0 {
			evenSum++
		} else {
			oddSum++
		}
	}
	if sum == 0 {
		return evenSum + oddSum - 1
	}
	return 0
}
