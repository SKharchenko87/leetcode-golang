package p0877

/*
Вообще всегда побеждает при четном кол-во
func stoneGame(piles []int) bool {
    return true
}
*/

func stoneGame(piles []int) bool {
	n := len(piles)
	if n%2 == 0 {
		return true
	}
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, n)
	}

	var f func(l, r int) int
	f = func(l, r int) int {
		// Базовый случай: осталась 1 кучка, текущий игрок забрает её целиком
		if l == r {
			return piles[l]
		}
		if mem[l][r] != 0 {
			return mem[l][r]
		}

		// Текущий игрок выбирает кучку и ВЫЧИТАЕТ из неё лучший ход соперника
		takeLeft := piles[l] - f(l+1, r)
		takeRight := piles[r] - f(l, r-1)

		mem[l][r] = max(takeLeft, takeRight)
		return mem[l][r]
	}

	// Если итоговый перевес первого игрока > 0, он выиграл
	return f(0, n-1) > 0
}
