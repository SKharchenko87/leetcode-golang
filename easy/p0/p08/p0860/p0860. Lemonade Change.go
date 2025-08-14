package p0860

func lemonadeChange(bills []int) bool {
	coin5Count, coin10Count := 0, 0
	for i := 0; i < len(bills) && coin5Count >= 0; i++ {
		coin5Count--
		if bills[i] == 20 && coin10Count > 0 {
			coin10Count--
		} else if bills[i] == 20 {
			coin5Count -= 2
		} else if bills[i] == 10 {
			coin10Count++
		} else {
			coin5Count += 2
		}
	}
	return coin5Count >= 0
}

func lemonadeChange1(bills []int) bool {
	coin5Count, coin10Count := 0, 0
	for i := 0; i < len(bills) && coin5Count >= 0; i++ {
		coin5Count--
		switch {
		case bills[i] == 20 && coin10Count > 0:
			coin10Count--
		case bills[i] == 20:
			coin5Count -= 2
		case bills[i] == 10:
			coin10Count++
		case bills[i] == 5:
			coin5Count += 2
		}
	}
	return coin5Count >= 0
}

func lemonadeChange0(bills []int) bool {
	coin5Count := 0
	coin10Count := 0
	for i := 0; i < len(bills) && coin5Count >= 0; i++ {
		switch bills[i] {
		case 20:
			if coin10Count > 0 {
				coin10Count--
				coin5Count--
			} else {
				coin5Count -= 3
			}
		case 10:
			coin10Count++
			coin5Count--
		case 5:
			coin5Count++
		}
	}
	return coin5Count >= 0
}
