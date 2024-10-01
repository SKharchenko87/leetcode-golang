package p1497

func canArrange(arr []int, k int) bool {
	cnt := len(arr) / 2
	m := make(map[int]int, k)
	for _, c := range arr {
		remainder := c % k
		if remainder <= 0 {
			remainder += k
		}
		if m[k-remainder] > 0 {
			m[k-remainder]--
			cnt--
		} else {
			m[remainder%k]++
		}
	}
	return cnt == 0
}

func canArrange1(arr []int, k int) bool {
	cnt := len(arr) / 2
	m := make(map[int]int, cnt)
	for _, c := range arr {
		switch {
		case m[k-c%k] > 0:
			m[k-c%k]--
		case m[-k-c%k] > 0:
			m[-k-c%k]--
		case m[-c%k] > 0:
			m[-c%k]--
		default:
			m[c%k]++
			continue
		}
		cnt--
	}
	return cnt == 0
}

func canArrange0(arr []int, k int) bool {
	l := len(arr)
	cnt := 0
	m := make(map[int]int, l)
	for _, c := range arr {
		if m[k-c%k] > 0 {
			m[k-c%k]--
			cnt++
		} else if m[-k-c%k] > 0 {
			m[-k-c%k]--
			cnt++
		} else if m[-c%k] > 0 {
			m[-c%k]--
			cnt++
		} else {
			m[c%k]++
		}
	}
	return cnt == l/2
}

//-4,-7,5,2,9,1,10,4,-8,-3
//-1,-1,2,2,0,1, 1,1,-2, 0
