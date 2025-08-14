package p1025

const MAX_LENGHT = 1000

var divisor = make([]bool, MAX_LENGHT+1)

func init() {
	divisor[2] = true
	divisor[3] = false
	for i := 4; i < MAX_LENGHT+1; i++ {
		for j := 1; j*j <= i; j++ {
			if i%j == 0 && !divisor[i-j] {
				divisor[i] = true
				break
			}
		}
	}
}

func divisorGame(n int) bool {
	return divisor[n]
}

func divisorGame0(n int) bool {
	switch n {
	case 0:
		return false
	case 1:
		return false
	case 2:
		return true
	case 3:
		return false
	}
	if n < 2 || n == 3 {
		return false
	}
	div := make([]bool, n+1)
	div[2] = true
	div[3] = false
	for i := 4; i < n+1; i++ {
		for j := 1; j*j <= i; j++ {
			if i%j == 0 && !div[i-j] {
				div[i] = true
				break
			}
		}
	}
	return div[n]
}

func divisorGame1(n int) bool {
	return n%2 == 0
}
