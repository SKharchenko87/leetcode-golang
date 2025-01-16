package p2429

func minimizeXor(num1 int, num2 int) int {
	countBits := 0
	for i := 0; num2>>i > 0; i++ {
		countBits += num2 >> i & 1
	}
	res := 0
	for i := 30; i >= 0 && countBits > 0; i-- {
		if (num1>>i)&1 == 1 {
			res |= 1 << i
			countBits--
		}
	}
	for i := 0; countBits > 0; i++ {
		if (res>>i)&1 == 0 {
			res |= 1 << i
			countBits--
		}
	}
	return res
}

func minimizeXor0(num1 int, num2 int) int {
	//fmt.Printf("%b\n", num1)
	//fmt.Printf("%b\n", num2)
	cnt := 0
	for i := 0; i < 31; i++ {
		cnt += num2 >> i & 1
	}
	res := 0
	for i := 30; i >= 0 && cnt > 0; i-- {
		if (num1>>i)&1 == 1 {
			res |= 1 << i
			cnt--
		}
	}
	for i := 0; i < 31 && cnt > 0; i++ {
		if (res>>i)&1 == 0 {
			res |= 1 << i
			cnt--
		}
	}
	return res
}
