package p2616

import "sort"

func minimizeMax(nums []int, p int) int {
	// Если не нужно создавать ни одной пары, максимальная разница 0
	if p == 0 {
		return 0
	}

	// Сортируем массив для упрощения поиска пар
	sort.Ints(nums)
	n := len(nums)

	// Функция для проверки, можно ли выбрать p пар с максимальной разницей <= mid
	canFormPairs := func(mid int) bool {
		count := 0
		i := 0
		for i < n-1 && count < p {
			// Если текущая пара удовлетворяет условию, берем ее и пропускаем следующий элемент
			if nums[i+1]-nums[i] <= mid {
				count++
				i += 2
			} else {
				i++
			}
		}
		return count >= p
	}

	// Бинарный поиск минимально возможной максимальной разницы
	left, right := 0, nums[n-1]-nums[0]
	for left < right {
		mid := left + (right-left)/2
		if canFormPairs(mid) {
			// Пробуем найти меньшее значение
			right = mid
		} else {
			// Нужно увеличить допустимую разницу
			left = mid + 1
		}
	}

	return left
}

// 1 2 2 3
// {2;2} {1;3} => max(2-2; 3-1) = max(0,2) = 2
// {1;2} {2;3} => max(2-1; 3-2) = max(1,1) = 1
