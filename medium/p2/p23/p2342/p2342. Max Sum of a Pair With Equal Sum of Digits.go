package p2342

func maximumSum(nums []int) int {
	prevMax := make([]int, 82)
	res := -1
	for _, n := range nums {
		x := sumOfDigit(n)
		if prevMax[x] > 0 {
			res = max(prevMax[x]+n, res)
		}
		prevMax[x] = max(prevMax[x], n)
	}
	return res
}

func maximumSum2(nums []int) int {
	m := make([][2]int, 82)
	res := -1
	for _, n := range nums {
		x := sumOfDigit(n)
		y := m[x]
		if y[0] > 0 {
			if y[0] <= n {
				m[x][0] = n
				m[x][1] = y[0]
				if res < n+y[0] {
					res = n + y[0]
				}
			} else if y[0] > n && y[1] <= n {
				m[x][0] = y[0]
				m[x][1] = n
				if res < n+y[0] {
					res = n + y[0]
				}
			}
		} else {
			m[x] = [2]int{n, 0}
		}
	}
	return res
}

func maximumSum1(nums []int) int {
	m := make(map[int][2]int, 81)
	res := -1
	for _, n := range nums {
		x := sumOfDigit(n)
		if y, ok := m[x]; ok {
			if y[0] <= n {
				m[x] = [2]int{n, y[0]}
				if res < n+y[0] {
					res = n + y[0]
				}
			} else if y[0] > n && y[1] <= n {
				m[x] = [2]int{y[0], n}
				if res < n+y[0] {
					res = n + y[0]
				}
			}
		} else {
			m[x] = [2]int{n, 0}
		}
	}
	return res
}

func maximumSum0(nums []int) int {
	m := make(map[int][2]int, 81)
	for _, n := range nums {
		x := sumOfDigit(n)
		if y, ok := m[x]; ok {
			if y[0] <= n {
				m[x] = [2]int{n, y[0]}
			} else if y[0] > n && y[1] <= n {
				m[x] = [2]int{y[0], n}
			}
		} else {
			m[x] = [2]int{n, 0}
		}
	}
	res := -1
	for _, v := range m {
		if v[1] != 0 && res < v[0]+v[1] {
			res = v[0] + v[1]
		}
	}
	return res
}

func sumOfDigit(n int) int {
	res := 0
	for n > 0 {
		res += n % 10
		n /= 10
	}
	return res
}
