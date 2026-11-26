package p3360

func canAliceWin(n int) bool {
	diff := 10
	for n >= 0 {
		n -= diff
		diff--
	}
	return diff%2 == 0
}

func canAliceWin0(n int) bool {
	if n < 10 {
		return false
	}
	winAlice := true
	i := 10
	for n >= 0 {
		winAlice = !winAlice
		n -= i
		i--
	}
	return winAlice
}
