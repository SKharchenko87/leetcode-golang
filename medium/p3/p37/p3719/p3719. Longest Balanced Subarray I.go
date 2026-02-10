package p3719

func longestBalanced(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// Предрасчет для мгновенной проверки уникальности
	prevOccurrence := make([]int, n)
	nextOccurrence := make([]int, n)
	lastSeen := make(map[int]int)

	// Заполняем prev
	for i, v := range nums {
		if idx, ok := lastSeen[v]; ok {
			prevOccurrence[i] = idx
		} else {
			prevOccurrence[i] = -1
		}
		lastSeen[v] = i
	}

	// Заполняем next
	for i := range lastSeen {
		delete(lastSeen, i)
	}
	for i := n - 1; i >= 0; i-- {
		if idx, ok := lastSeen[nums[i]]; ok {
			nextOccurrence[i] = idx
		} else {
			nextOccurrence[i] = -1
		}
		lastSeen[nums[i]] = i
	}

	// Считаем начальное кол-во уникальных для всего массива [0, n-1]
	totalEvens, totalOdds := 0, 0
	for i, v := range nums {
		if prevOccurrence[i] == -1 {
			if v%2 == 0 {
				totalEvens++
			} else {
				totalOdds++
			}
		}
	}

	res := 0
	for i := 0; i < n && n-i > res; i++ {
		if n-i <= res {
			break
		}
		// Временные счетчики для текущего i
		curEvens, curOdds := totalEvens, totalOdds

		for j := n - 1; j >= i && j-i+1 > res; j-- {

			if curEvens == curOdds && curEvens > 0 {
				res = j - i + 1
				break
			}

			// "Удаляем" элемент nums[j] для следующей итерации j-цикла
			// Элемент перестает быть уникальным в [i, j], если его ПРЕДЫДУЩЕЕ вхождение >= i
			if prevOccurrence[j] < i {
				if nums[j]%2 == 0 {
					curEvens--
				} else {
					curOdds--
				}
			}
		}

		// Обновляем totalEvens/totalOdds для следующего i (удаляем nums[i] из рассмотрения)
		// Элемент nums[i] исчезает из уникальных для суффикса, если он больше НЕ встречается
		if nextOccurrence[i] == -1 {
			if nums[i]%2 == 0 {
				totalEvens--
			} else {
				totalOdds--
			}
		}
	}

	return res
}
