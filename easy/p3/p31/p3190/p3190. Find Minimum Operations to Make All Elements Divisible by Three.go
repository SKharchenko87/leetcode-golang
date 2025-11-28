package p3190

func minimumOperations(nums []int) int {
	res := 0
	var remainder int
	for _, num := range nums {
		remainder = num % 3
		res += remainder&1 + remainder>>1&1
	}
	return res
}

func minimumOperations0(nums []int) int {
	res := 0
	for _, num := range nums {
		if num%3 != 0 {
			res++
		}
	}
	return res
}

func minimumOperationsX(nums []int) int {
	sum := 0
	for _, x := range nums {
		remainder := x % 3
		sum += min(remainder, 3-remainder)
	}
	return sum
}
