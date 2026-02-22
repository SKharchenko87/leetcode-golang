package p0868

import "math/bits"

func binaryGap(n int) int {
	x := uint(n)
	x <<= bits.LeadingZeros(x) + 1
	res := 0
	for x > 0 {
		res = max(res, bits.LeadingZeros(x)+1)
		x <<= bits.LeadingZeros(x) + 1
	}
	return res
}

func binaryGap3(n int) int {
	res := 0
	lastIndexOne := -1
	for index := 0; n > 0; index++ {
		if n&1 == 1 {
			if lastIndexOne != -1 {
				res = max(res, index-lastIndexOne)
			}
			lastIndexOne = index
		}
		n >>= 1
	}
	return res
}

func binaryGap2(n int) int {
	isViewOne := false
	res := 0
	l := 0
	for n > 0 {
		for n > 0 && n&1 == 0 {
			n >>= 1
			l++
		}
		n >>= 1
		if !isViewOne {
			isViewOne = true
		} else {
			res = max(res, l)
		}
		l = 1
	}
	return res
}

func binaryGap1(n int) int {
	l := 0
	res := 0

	execute := func(f func() int) {
		for n > 0 && n&1 == 0 {
			n >>= 1
			l++
		}
		n >>= 1
		res = f()
		l = 1
	}

	execute(func() int {
		return 0
	})

	for n > 0 {
		execute(func() int {
			return max(res, l)
		})
	}
	return res
}

func binaryGap0(n int) int {
	l := 0

	for n > 0 && n&1 == 0 {
		n >>= 1
		l++
	}
	n >>= 1
	res := 0
	l = 1

	for n > 0 {
		for n > 0 && n&1 == 0 {
			n >>= 1
			l++
		}
		n >>= 1
		res = max(res, l)
		l = 1
	}
	return res
}
