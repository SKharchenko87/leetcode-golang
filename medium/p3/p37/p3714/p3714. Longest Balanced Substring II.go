package p3714

func longestBalanced(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	res := 1

	// Для случая a=b=c используем мапу, так как 2D-массив слишком велик
	// Упаковываем две разности в один int64 для скорости
	abc := make(map[int64]int, n)
	abc[0] = -1 // (totalA-totalB)=0, (totalB-totalC)=0 на старте

	// Для пар {a,b}, {a,c}, {b,c} используем слайсы и offset
	// Чтобы не делать clear(), храним [индекс, id_стены]
	type entry struct{ pos, wallID int }
	pairMaps := [3][]entry{
		make([]entry, 2*n+1), // для ab (стена c)
		make([]entry, 2*n+1), // для ac (стена b)
		make([]entry, 2*n+1), // для bc (стена a)
	}

	// Инициализация: ставим -2 в pos, чтобы пометить "пусто"
	for j := 0; j < 3; j++ {
		for k := range pairMaps[j] {
			pairMaps[j][k].pos = -2
		}
		pairMaps[j][n] = entry{-1, 0} // Баланс 0 на старте
	}

	tA, tB, tC := 0, 0, 0 // Префиксные суммы
	wA, wB, wC := 0, 0, 0 // Счетчики "стен" (сколько раз встретили символ)

	// Для серий из одного символа (aaaa)
	curLen := 0
	var lastChar byte

	for i := 0; i < n; i++ {
		char := s[i]

		// 1. Одиночные серии
		if char == lastChar {
			curLen++
		} else {
			curLen = 1
			lastChar = char
		}
		if curLen > res {
			res = curLen
		}

		// 2. Обновляем счетчики
		if char == 'a' {
			tA++
		} else if char == 'b' {
			tB++
		} else if char == 'c' {
			tC++
		}
		// Счетчики стен увеличиваются только когда символ РЕАЛЬНО мешает паре
		if char == 'a' {
			wA++
		}
		if char == 'b' {
			wB++
		}
		if char == 'c' {
			wC++
		}

		// 3. Логика для ПАР (ab, ac, bc)
		// Пара AB (стена C)
		diffAB := tA - tB + n
		if pairMaps[0][diffAB].pos != -2 && pairMaps[0][diffAB].wallID == wC {
			if dist := i - pairMaps[0][diffAB].pos; dist > res {
				res = dist
			}
		} else {
			pairMaps[0][diffAB] = entry{i, wC}
		}

		// Пара AC (стена B)
		diffAC := tA - tC + n
		if pairMaps[1][diffAC].pos != -2 && pairMaps[1][diffAC].wallID == wB {
			if dist := i - pairMaps[1][diffAC].pos; dist > res {
				res = dist
			}
		} else {
			pairMaps[1][diffAC] = entry{i, wB}
		}

		// Пара BC (стена A)
		diffBC := tB - tC + n
		if pairMaps[2][diffBC].pos != -2 && pairMaps[2][diffBC].wallID == wA {
			if dist := i - pairMaps[2][diffBC].pos; dist > res {
				res = dist
			}
		} else {
			pairMaps[2][diffBC] = entry{i, wA}
		}

		// 4. Логика для ТРОЙКИ (a=b=c)
		// Упаковываем разности (tA-tB) и (tB-tC) в 64 бита
		key := int64(tA-tB)<<32 | int64(tB-tC)&0xFFFFFFFF
		if v, ok := abc[key]; ok {
			if dist := i - v; dist > res {
				res = dist
			}
		} else {
			abc[key] = i
		}
	}

	return res
}

func longestBalanced0(s string) int {
	n := len(s)
	var a, b, c, keyAB, keyAC, keyBC int
	ab := make(map[int]int, n) // a = +1 b = -1
	ac := make(map[int]int, n) // a = +1 c = -1
	bc := make(map[int]int, n) // b = +1 c = -1
	res := 1
	ab[0] = -1
	ac[0] = -1
	bc[0] = -1
	abc := make(map[[2]int]int, n)
	abc[[2]int{0, 0}] = -1
	totalA, totalB, totalC := 0, 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			clear(bc)
			bc[0] = i
			keyBC = 0
			keyAB++
			keyAC++
			a++
			b = 0
			c = 0

			totalA++
		case 'b':
			clear(ac)
			ac[0] = i
			keyAC = 0
			keyAB--
			keyBC++
			a = 0
			b++
			c = 0

			totalB++

		case 'c':
			clear(ab)
			ab[0] = i
			keyAB = 0
			keyAC--
			keyBC--
			a = 0
			b = 0
			c++

			totalC++

		}
		res = max(res, a, b, c)
		key := [2]int{totalA - totalB, totalB - totalC}
		if v, ok := ab[keyAB]; ok {
			res = max(res, i-v)
		} else {
			ab[keyAB] = i
		}
		if v, ok := ac[keyAC]; ok {
			res = max(res, i-v)
		} else {
			ac[keyAC] = i
		}
		if v, ok := bc[keyBC]; ok {
			res = max(res, i-v)
		} else {
			bc[keyBC] = i
		}
		if v, ok := abc[key]; ok {
			res = max(res, i-v)
		} else {
			abc[key] = i
		}
	}

	return res
}
