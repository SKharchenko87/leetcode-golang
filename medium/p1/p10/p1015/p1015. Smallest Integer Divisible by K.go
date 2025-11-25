package p1015

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	start := 1
	count := 1
	for start < k {
		start = start*10 + 1
		count++
	}
	for start%k > 0 {
		start = (start%k)*10 + 1
		count++
	}
	return count
}

func smallestRepunitDivByK2(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	count := 1
	for start := 1; start%k > 0; count++ {
		start = (start%k)*10 + 1
	}
	return count
}

func smallestRepunitDivByK1(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	start := 1
	count := 1
	for start%k > 0 {
		start = (start%k)*10 + 1
		count++
	}
	return count
}

func smallestRepunitDivByK0(k int) (count int) {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	d := 0
	for {
		d = (d*10 + 1) % k
		count++
		if d == 0 {
			return count
		}
	}
}
