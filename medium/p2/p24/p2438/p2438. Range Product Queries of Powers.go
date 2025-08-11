package p2438

const Mod = 1e9 + 7

func productQueries(n int, queries [][]int) []int {
	power2 := make([]int, 0, 31)
	for i := 0; i < 31; i++ {
		p2 := 1 << i
		if n&p2 > 0 {
			power2 = append(power2, p2)
		}
	}
	l := len(power2)
	result := make([][]int, l)
	for i := 0; i < l; i++ {
		result[i] = make([]int, l)
		x := 1
		for j := i; j < l; j++ {
			x = x * power2[j] % Mod
			result[i][j] = x
		}
	}
	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		ans[i] = result[queries[i][0]][queries[i][1]]
	}
	return ans
}

func productQueries0(n int, queries [][]int) []int {
	power2 := make([]int, 0, 31)
	for i := 0; i < 31; i++ {
		p2 := 1 << i
		if n&p2 > 0 {
			power2 = append(power2, p2)
		}
	}
	l := len(queries)
	ans := make([]int, l)
	for i := 0; i < l; i++ {
		res := 1
		for j := queries[i][0]; j <= queries[i][1]; j++ {
			res *= power2[j]
			res %= 1e9 + 7
		}
		ans[i] = res
	}
	return ans
}
