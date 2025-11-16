package p3234

func numberOfSubstrings(s string) int {
	l := len(s)
	ones := make([]int, l+1)
	zeros := make([]int, l+1)
	orderOne := make([]int, l+1)
	orderZero := make([]int, l+1)
	for i := l - 1; i >= 0; i-- {
		zeros[i] = zeros[i+1]
		ones[i] = ones[i+1]
		if s[i] == '0' {
			zeros[i]++
			orderOne[i] = 0
			orderZero[i] = orderZero[i+1] + 1
		} else {
			ones[i]++
			orderOne[i] = orderOne[i+1] + 1
			orderZero[i] = 0
		}
	}
	var count, one, zero int
	for i := 0; i < l; i++ {
		for j := i; j < l && l-j+one >= zero*zero; j++ {
			one = ones[i] - ones[j+1]
			zero = zeros[i] - zeros[j+1]
			if zeros[i] == 0 {
				count += l - j
				break
			}
			if one >= zero*zero {
				if orderOne[j] > 1 {
					count += orderOne[j]
					j += orderOne[j] - 1
				} else {
					count++
				}
			} else if orderZero[j] > 1 {
				j += orderZero[j] - 1
			}
		}
		one, zero = 0, 0
	}
	return count
}

// TLE
func numberOfSubstrings1(s string) int {
	l := len(s)
	ones := make([]int, l+1)
	zeros := make([]int, l+1)

	for i := l - 1; i >= 0; i-- {
		zeros[i] = zeros[i+1]
		ones[i] = ones[i+1]
		if s[i] == '0' {
			zeros[i]++
		} else {
			ones[i]++
		}
	}
	var count, one, zero int
	for i := 0; i < l; i++ {
		for j := i; j < l && l-j+one >= zero*zero; j++ {
			one = ones[i] - ones[j+1]
			zero = zeros[i] - zeros[j+1]
			if one >= zero*zero {
				count++
			}
		}
		one, zero = 0, 0
	}
	return count
}

// TLE
func numberOfSubstrings0(s string) int {
	l := len(s)
	ones := make([]int, l+1)
	zeros := make([]int, l+1)

	for i := l - 1; i >= 0; i-- {
		zeros[i] = zeros[i+1]
		ones[i] = ones[i+1]
		if s[i] == '0' {
			zeros[i]++
		} else {
			ones[i]++
		}
	}
	var count int
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			one := ones[i] - ones[j+1]
			zero := zeros[i] - zeros[j+1]
			if one >= zero*zero {
				count++
			}
		}
	}
	return count
}
