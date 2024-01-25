package p4

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type it struct {
	l, r int
}

func countOfPairs(n int, x int, y int) []int64 {
	if x > y {
		x, y = y, x
	}
	ret := []int64{}
	intervals := []it{}
	for i := 1; i < n; i++ {
		short := 1 + abs(i-x)
		if y-i > short {
			if y < n {
				intervals = append(intervals, it{short + 1, short + n - y})
			}
			t := y - i - short
			step := t / 2
			if t&1 == 1 {
				step++
			}
			intervals = append(intervals, it{short, short + step - 1})
			intervals = append(intervals, it{1, y - step - i})
		} else {
			intervals = append(intervals, it{1, n - i})
		}
	}
	arr := make([]int64, n+1)
	for _, i := range intervals {
		arr[i.l]++
		arr[i.r+1]--
	}
	cur := int64(0)
	for k := 1; k <= n; k++ {
		cur += arr[k]
		ret = append(ret, 2*cur)
	}

	return ret
}
