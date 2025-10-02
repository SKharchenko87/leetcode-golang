package p1726

func tupleSameProduct(nums []int) int {
	m := make(map[int]int)
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			m[nums[i]*nums[j]]++
			res += (m[nums[i]*nums[j]] - 1) * 8
		}
	}
	return res
}
