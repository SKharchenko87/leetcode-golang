package p3201

func maximumLength(nums []int) int {

	cnt := [3]int{}

	b := nums[0] & 1
	cnt[b]++

	last := b
	cnt[2]++

	for i := 1; i < len(nums); i++ {
		b = nums[i] & 1
		cnt[b]++

		if b != last {
			last = b
			cnt[2]++
		}

	}
	return max(cnt[0], cnt[1], cnt[2])
}

func maximumLength2(nums []int) int {
	cnt0, cnt1 := 0, 0
	cnt := 0

	b := nums[0] & 1
	if b == 0 {
		cnt0++
	} else {
		cnt1++
	}

	cnt++
	last := b

	for i := 1; i < len(nums); i++ {
		b = nums[i] & 1
		if b == 0 {
			cnt0++
		} else {
			cnt1++
		}

		if b != last {
			last = b
			cnt++
		}

	}
	return max(cnt0, cnt1, cnt)
}

func maximumLength1(nums []int) int {
	cnt0, cnt1 := 0, 0
	cnt01, cnt10 := 0, 0
	last01, last10 := 1, 0
	b := nums[0] & 1
	if b == 0 {
		cnt0++
		cnt01++
		last01 = b
	} else {
		cnt1++
		cnt10++
		last10 = b
	}

	for i := 1; i < len(nums); i++ {
		b = nums[i] & 1
		if b == 0 {
			cnt0++
		} else {
			cnt1++
		}

		if b != last01 {
			last01 = b
			cnt01++
		}

		if b != last10 {
			last10 = b
			cnt10++
		}

	}
	return max(cnt0, cnt1, cnt10, cnt01)
}

func maximumLength0(nums []int) int {
	cnt0, cnt1 := 0, 0
	cnt01, cnt10 := 0, 0
	//last01, last10 := nums[0]&1, nums[0]&1
	last01, last10 := 2, 2

	for i := 0; i < len(nums); i++ {
		b := nums[i] & 1
		if b == 0 {
			cnt0++
		} else {
			cnt1++
		}

		if last01 < 2 {
			if b != last01 {
				last01 = b
				cnt01++
			}
		} else {
			last01 = b
			cnt01++
		}

		if last10 < 2 {
			if b != last10 {
				last10 = b
				cnt10++
			}
		} else {
			last10 = b
			cnt10++
		}
	}
	return max(cnt0, cnt1, cnt10, cnt01)
}
