package p1262

import "sort"

const LIMIT = 40001

func maxSumDivThree(nums []int) int {
	min11, min12 := LIMIT, LIMIT
	min21, min22 := LIMIT, LIMIT
	count1, count2 := 0, 0
	sum := 0
	var exe = func(count *int, min1, min2 *int, x *int) {
		*count++
		if *min2 >= *x {
			*min1, *min2 = *min2, *x
		} else if *min1 > *x {
			*min1 = *x
		}
	}
	for _, num := range nums {
		sum += num
		switch d := num % 3; d {
		case 1:
			exe(&count1, &min11, &min12, &num)
		case 2:
			exe(&count2, &min21, &min22, &num)
		}
	}

	d1 := count1 % 3
	d2 := count2 % 3
	if d1 == d2 {
		return sum
	}

	if (d1 == 0 && d2 == 1) || (d1 == 1 && d2 == 2) || (d1 == 2 && d2 == 0) {
		return max(sum-min11-min12, sum-min22)
	}

	return max(sum-min21-min22, sum-min12)
}

func maxSumDivThree4(nums []int) int {
	min11, min12 := LIMIT, LIMIT
	min21, min22 := LIMIT, LIMIT
	count1, count2 := 0, 0
	sum := 0
	var exe = func(count *int, min1, min2 *int, x *int) {
		*count++
		if *min2 >= *x {
			*min1, *min2 = *min2, *x
		} else if *min1 >= *x {
			*min1 = *x
		}
	}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		switch nums[i] % 3 {
		case 1:
			exe(&count1, &min11, &min12, &nums[i])
		case 2:
			exe(&count2, &min21, &min22, &nums[i])
		}
	}

	d1 := count1 % 3
	d2 := count2 % 3
	if d1 == d2 {
		return sum
	}

	if (d1 == 0 && d2 == 1) || (d1 == 1 && d2 == 2) || (d1 == 2 && d2 == 0) {
		return max(sum-min11-min12, sum-min22)
	}

	return max(sum-min21-min22, sum-min12)
}

func maxSumDivThree3(nums []int) int {
	min11, min12 := LIMIT, LIMIT
	min21, min22 := LIMIT, LIMIT
	count1, count2 := 0, 0
	sum := 0
	var exe = func(count *int, min1, min2 *int, x *int) {
		*count++
		if *min2 >= *x {
			*min1, *min2 = *min2, *x
		} else if *min1 >= *x {
			*min1 = *x
		}
	}
	for _, num := range nums {
		sum += num
		d := num % 3
		if d == 1 {
			exe(&count1, &min11, &min12, &num)
		} else if d == 2 {
			exe(&count2, &min21, &min22, &num)
		}
	}

	d1 := count1 % 3
	d2 := count2 % 3
	if d1 == d2 {
		return sum
	}

	if (d1 == 0 && d2 == 1) || (d1 == 1 && d2 == 2) || (d1 == 2 && d2 == 0) {
		return max(sum-min11-min12, sum-min22)
	}

	return max(sum-min21-min22, sum-min12)
}

func maxSumDivThree2(nums []int) int {
	min11, min12 := LIMIT, LIMIT
	min21, min22 := LIMIT, LIMIT
	count1, count2 := 0, 0
	sum := 0
	var exe = func(count *int, min1, min2 *int, x *int) {
		*count++
		if *min2 >= *x {
			*min1, *min2 = *min2, *x
		} else if *min1 >= *x {
			*min1 = *x
		}
	}
	for _, num := range nums {
		sum += num
		switch d := num % 3; d {
		case 1:
			exe(&count1, &min11, &min12, &num)
		case 2:
			exe(&count2, &min21, &min22, &num)
		}
	}

	d1 := count1 % 3
	d2 := count2 % 3
	if d1 == d2 {
		return sum
	}

	if (d1 == 0 && d2 == 1) || (d1 == 1 && d2 == 2) || (d1 == 2 && d2 == 0) {
		return max(sum-min11-min12, sum-min22)
	}

	return max(sum-min21-min22, sum-min12)
}

func maxSumDivThree1(nums []int) int {
	min10, min11, min12 := LIMIT, LIMIT, LIMIT
	min20, min21, min22 := LIMIT, LIMIT, LIMIT
	count1, count2 := 0, 0
	sum := 0
	var exe = func(count *int, min0, min1, min2 *int, x *int) {
		(*count)++
		if *min2 >= *x {
			*min0, *min1, *min2 = *min1, *min2, *x
		} else if *min1 >= *x {
			*min0, *min1 = *min1, *x
		} else if *min0 >= *x {
			*min0 = *x
		}
	}
	for _, num := range nums {
		sum += num
		switch d := num % 3; d {
		case 1:
			exe(&count1, &min10, &min11, &min12, &num)
		case 2:
			exe(&count2, &min20, &min21, &min22, &num)
		}
	}

	d1 := count1 % 3
	d2 := count2 % 3
	if d1 == d2 {
		return sum
	}

	if (d1 == 0 && d2 == 1) || (d1 == 1 && d2 == 2) || (d1 == 2 && d2 == 0) {
		return max(sum-min11-min12, sum-min22)
	}

	return max(sum-min21-min22, sum-min12)
}

func maxSumDivThree0(nums []int) int {
	min10, min11, min12 := LIMIT, LIMIT, LIMIT
	min20, min21, min22 := LIMIT, LIMIT, LIMIT
	count1, count2 := 0, 0
	sum := 0
	for _, num := range nums {
		sum += num
		switch d := num % 3; d {
		case 1:
			count1++
			if min12 >= num {
				min10, min11, min12 = min11, min12, num
			} else if min11 >= num {
				min10, min11 = min11, num
			} else if min10 >= num {
				min10 = num
			}
		case 2:
			count2++
			if min22 >= num {
				min20, min21, min22 = min21, min22, num
			} else if min21 >= num {
				min20, min21 = min21, num
			} else if min20 >= num {
				min20 = num
			}
		}
	}

	d1 := count1 % 3
	d2 := count2 % 3
	if d1 == d2 {
		return sum
	}

	if (d1 == 0 && d2 == 1) || (d1 == 1 && d2 == 2) || (d1 == 2 && d2 == 0) {
		return max(sum-min11-min12, sum-min22)
	}

	return max(sum-min21-min22, sum-min12)
}

func accumulate(v []int) int {
	ans := 0
	for _, x := range v {
		ans += x
	}
	return ans
}

func maxSumDivThreeEditorial1(nums []int) int {
	// Use v[0], v[1], v[2] to represent a, b, c respectively.
	v := make([][]int, 3)
	for _, num := range nums {
		v[num%3] = append(v[num%3], num)
	}
	sort.Slice(v[1], func(i, j int) bool {
		return v[1][i] > v[1][j]
	})
	sort.Slice(v[2], func(i, j int) bool {
		return v[2][i] > v[2][j]
	})
	ans, lb, lc := 0, len(v[1]), len(v[2])
	for cntb := max(lb-2, 0); cntb <= lb; cntb++ {
		for cntc := max(lc-2, 0); cntc <= lc; cntc++ {
			if (cntb-cntc)%3 == 0 {
				ans = max(ans, accumulate(v[1][:cntb])+accumulate(v[2][:cntc]))
			}
		}
	}
	return ans + accumulate(v[0])
}

func maxSumDivThreeEditorial2(nums []int) int {
	// Use v[0], v[1], v[2] to represent a, b, c respectively.
	v := make([][]int, 3)
	for _, num := range nums {
		v[num%3] = append(v[num%3], num)
	}
	sort.Slice(v[1], func(i, j int) bool {
		return v[1][i] > v[1][j]
	})
	sort.Slice(v[2], func(i, j int) bool {
		return v[2][i] > v[2][j]
	})
	tot, remove := accumulate(nums), 0x3f3f3f3f
	if tot%3 == 0 {
		remove = 0
	} else if tot%3 == 1 {
		if len(v[1]) >= 1 {
			remove = min(remove, v[1][len(v[1])-1])
		}
		if len(v[2]) >= 2 {
			remove = min(remove, v[2][len(v[2])-2]+v[2][len(v[2])-1])
		}
	} else {
		if len(v[1]) >= 2 {
			remove = min(remove, v[1][len(v[1])-2]+v[1][len(v[1])-1])
		}
		if len(v[2]) >= 1 {
			remove = min(remove, v[2][len(v[2])-1])
		}
	}
	return tot - remove
}

func maxSumDivThreeEditorial3(nums []int) int {
	f := []int{0, -0x3f3f3f3f, -0x3f3f3f3f}
	for _, num := range nums {
		g := make([]int, 3)
		for i := 0; i < 3; i++ {
			g[(i+num)%3] = max(f[(i+num)%3], f[i]+num)
		}
		f = g
	}
	return f[0]
}
