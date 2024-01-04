package p2610

func findMatrix(nums []int) [][]int {
	m := map[int]int{}
	res := [][]int{}
	for _, v := range nums {
		m[v]++
		if len(res) < m[v] {
			res = append(res, []int{})
		}
		res[m[v]-1] = append(res[m[v]-1], v)
	}
	return res
}
