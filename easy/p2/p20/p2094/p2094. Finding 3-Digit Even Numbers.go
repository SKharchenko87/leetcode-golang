package p2094

func findEvenNumbers(digits []int) []int {
	freq := [10]int8{}
	for i := 0; i < len(digits); i++ {
		freq[digits[i]]++
	}
	res := make([]int, 0)
	for i := 100; i < 1000; i += 2 {
		d0 := i % 10
		d1 := i / 10 % 10
		d2 := i / 100 % 10
		freq[d0]--
		freq[d1]--
		freq[d2]--
		if freq[d0] >= 0 && freq[d1] >= 0 && freq[d2] >= 0 {
			res = append(res, i)
		}
		freq[d0]++
		freq[d1]++
		freq[d2]++
	}
	return res
}

func findEvenNumbers0(digits []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(digits); i++ {
		m[digits[i]]++
	}
	res := make([]int, 0)
	for i := 100; i < 1000; i += 2 {
		d0 := i % 10
		d1 := i / 10 % 10
		d2 := i / 100 % 10
		m[d0]--
		m[d1]--
		m[d2]--
		if m[d0] >= 0 && m[d1] >= 0 && m[d2] >= 0 {
			res = append(res, i)
		}
		m[d0]++
		m[d1]++
		m[d2]++
	}
	return res
}
