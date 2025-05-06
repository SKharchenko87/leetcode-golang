package p0838

func pushDominoes(dominoes string) string {
	l := len(dominoes)
	left := make([]int, l)
	right := make([]int, l)
	move := false
	power := 0
	for i := 0; i < l; i++ {
		drop(&move, dominoes, i, &power, &right, 'L')
	}

	move = false
	power = 0
	for i := l - 1; i >= 0; i-- {
		drop(&move, dominoes, i, &power, &left, 'R')
	}
	res := []byte(dominoes)
	for i := 0; i < l; i++ {
		if left[i] > 0 && right[i] == 0 {
			res[i] = 'L'
		} else if left[i] == 0 && right[i] > 0 {
			res[i] = 'R'
		} else if left[i] > 0 && right[i] > 0 && left[i] < right[i] {
			res[i] = 'L'
		} else if left[i] > 0 && right[i] > 0 && left[i] > right[i] {
			res[i] = 'R'
		}
	}

	return string(res)
}

func drop(move *bool, dominoes string, i int, power *int, right *[]int, ch byte) {
	var antiCh byte
	if ch == 'L' {
		antiCh = 'R'
	} else {
		antiCh = 'L'
	}
	if *move {
		if dominoes[i] == ch {
			*move = false
			*power = 0
		} else if dominoes[i] == antiCh {
			*power = 0
		}
		(*right)[i] = *power
		*power++
	} else if dominoes[i] == antiCh {
		*move = true
		*power = 1
	}
}

func pushDominoes0(dominoes string) string {
	l := len(dominoes)
	left := make([]int, l)
	right := make([]int, l)
	move := false
	power := 0
	for i := 0; i < l; i++ {
		if move {
			if dominoes[i] == 'L' {
				move = false
				power = 0
			} else if dominoes[i] == 'R' {
				power = 0
			}
			right[i] = power
			power++
		} else if dominoes[i] == 'R' {
			move = true
			power = 1
		}
	}

	move = false
	power = 0
	for i := l - 1; i >= 0; i-- {
		if move {
			if dominoes[i] == 'R' {
				move = false
				power = 0
			} else if dominoes[i] == 'L' {
				power = 0
			}
			left[i] = power
			power++
		} else if dominoes[i] == 'L' {
			move = true
			power = 1
		}
	}
	res := []byte(dominoes)
	for i := 0; i < l; i++ {
		if left[i] > 0 && right[i] == 0 {
			res[i] = 'L'
		} else if left[i] == 0 && right[i] > 0 {
			res[i] = 'R'
		} else if left[i] > 0 && right[i] > 0 && left[i] < right[i] {
			res[i] = 'L'
		} else if left[i] > 0 && right[i] > 0 && left[i] > right[i] {
			res[i] = 'R'
		}
	}

	return string(res)
}
