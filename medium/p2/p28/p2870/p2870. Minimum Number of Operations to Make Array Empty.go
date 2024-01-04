package p2870

func minOperations(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	count := 0
	for _, v := range m {
		if v == 1 {
			return -1
		}
		switch v % 3 {
		case 0:
			count += v / 3
		case 1, 2:
			count += v/3 + 1
		}
	}
	return count
}
