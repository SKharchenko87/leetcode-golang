package p1769

func minOperations(boxes string) []int {
	l := len(boxes)
	res := make([]int, l)
	i, cnt, prev := 0, 0, 0
	var f = func() {
		res[i] += cnt + prev
		prev = cnt + prev
		if boxes[i] == '1' {
			cnt++
		}
	}
	for i = 0; i < l; i++ {
		f()
	}
	cnt, prev = 0, 0
	for i = l - 1; i >= 0; i-- {
		f()
	}
	return res
}

func minOperations1(boxes string) []int {
	l := len(boxes)
	res := make([]int, l)
	cnt := 0
	prev := 0
	for i := 0; i < l; i++ {
		res[i] += cnt + prev
		prev = cnt + prev
		if boxes[i] == '1' {
			cnt++
		}
	}
	cnt = 0
	prev = 0
	for i := l - 1; i >= 0; i-- {
		res[i] += cnt + prev
		prev = cnt + prev
		if boxes[i] == '1' {
			cnt++
		}
	}
	return res
}

func minOperations0(boxes string) []int {
	l := len(boxes)
	cnt := 0
	right := make([]int, l)
	i := 0
	right[i] = 0
	if boxes[i] == '1' {
		cnt++
	}
	for i = 1; i < l; i++ {
		right[i] += cnt + right[i-1]
		if boxes[i] == '1' {
			cnt++
		}
	}
	cnt = 0
	left := make([]int, l)
	i = l - 1
	left[i] = 0
	if boxes[i] == '1' {
		cnt++
	}
	for i = l - 2; i >= 0; i-- {
		left[i] += cnt + left[i+1]
		if boxes[i] == '1' {
			cnt++
		}
	}
	res := make([]int, l)
	for j := range res {
		res[j] = left[j] + right[j]
	}
	return res
}

// 1100
// 1222 >
// 0135 >
// 2100 <
// 1000 <

// 1135

// 110
// 122
// 013
