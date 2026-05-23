package p2615

func distance(nums []int) []int64 {
	n := len(nums)
	ans := make([]int64, n)

	// Группируем индексы по значениям
	// Ключ - число из nums, значение - слайс всех его позиций
	groups := make(map[int][]int)
	for i, v := range nums {
		groups[v] = append(groups[v], i)
	}

	// Обрабатываем каждую группу отдельно
	for _, group := range groups {
		m := len(group)
		if m <= 1 {
			continue // Если число одно, дистанция 0 (уже в ans)
		}

		// Считаем общую сумму индексов в группе для начальной точки
		var totalSum int64
		for _, idx := range group {
			totalSum += int64(idx)
		}

		var leftSum int64
		for i, idx := range group {
			idx64 := int64(idx)
			i64 := int64(i)
			m64 := int64(m)

			// Сумма для правой части: (сумма всех индексов справа) - (текущий индекс * кол-во справа)
			// Сумма всех индексов справа = totalSum - leftSum - idx64
			rightSumOfIndices := totalSum - leftSum - idx64
			rightCount := m64 - 1 - i64

			rightDistances := rightSumOfIndices - (idx64 * rightCount)

			// Сумма для левой части: (текущий индекс * кол-во слева) - (сумма всех индексов слева)
			leftDistances := (idx64 * i64) - leftSum

			ans[idx] = leftDistances + rightDistances

			// Обновляем префиксную сумму для следующего шага
			leftSum += idx64
		}
	}

	return ans
}
