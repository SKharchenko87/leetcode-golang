package p1096

import "slices"

/*ToDo*/
func braceExpansionII(expression string) []string {
	res, _ := parse(expression, 0)

	// 1. Сортируем срез строк
	slices.Sort(res)

	// 2. Оставляем только уникальные элементы
	unique := make([]string, 0, len(res))
	for i, s := range res {
		if i == 0 || s != res[i-1] {
			unique = append(unique, s)
		}
	}

	return unique
}

func parse(s string, i int) ([]string, int) {
	total := []string{}
	current := []string{""}

	for i < len(s) {
		ch := s[i]

		switch ch {
		case '{':
			// Ныряем в рекурсию разбирать внутренности скобок
			subResult, nextI := parse(s, i+1)
			// Умножаем (декартово произведение) текущую цепочку на результат скобок
			current = cartesianProduct(current, subResult)
			i = nextI

		case '}':
			// Завершаем текущий уровень скобок
			total = append(total, current...)
			return total, i + 1

		case ',':
			// Запятая закрывает текущую ветку конкатенации
			total = append(total, current...)
			current = []string{""}
			i++

		default:
			// Считываем последовательность обычных букв
			start := i
			for i < len(s) && s[i] >= 'a' && s[i] <= 'z' {
				i++
			}
			word := s[start:i]
			current = cartesianProduct(current, []string{word})
		}
	}

	// Конец всей строки (верхний уровень)
	total = append(total, current...)
	return total, i
}

func cartesianProduct(a, b []string) []string {
	res := make([]string, 0, len(a)*len(b))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			res = append(res, a[i]+b[j])
		}
	}
	return res
}
