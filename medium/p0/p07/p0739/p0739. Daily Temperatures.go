package p0739

func dailyTemperatures(temperatures []int) []int {
	l := len(temperatures)
	res := make([]int, l)
	for i := l - 2; i >= 0; i-- {
		j := i + 1
		if temperatures[i] == temperatures[j] {
			if res[j] == 0 {
				res[i] = 0
			} else {
				res[i] = res[j] + 1
			}
		} else {
			for ; j < l; j++ {
				if temperatures[i] < temperatures[j] {
					break
				}
			}
			if j == l {
				res[i] = 0
			} else {
				res[i] = j - i
			}
		}
	}
	return res
}
