package p2169

func countOperations(num1 int, num2 int) int {
	cnt := 0
	for num1 != 0 && num2 != 0 {
		cnt += num1 / num2
		num1, num2 = num2, num1%num2
	}
	return cnt
}

func countOperations2(num1 int, num2 int) int {
	cnt := 0
	for ; num1 != 0 && num2 != 0; cnt++ {
		if num1 < num2 {
			num2 -= num1
		} else {
			num1 -= num2
		}
	}
	return cnt
}

func countOperations1(num1 int, num2 int) int {
	cnt := 0
	for num1 != 0 && num2 != 0 {
		if num1 < num2 {
			num1, num2 = num2, num1
		}
		num1 -= num2
		cnt++
	}
	return cnt
}

func countOperations0(num1 int, num2 int) int {
	if num1 < num2 {
		num1, num2 = num2, num1
	}
	if num2 == 0 {
		return 0
	}
	return countOperations0(num1-num2, num2) + 1
}
