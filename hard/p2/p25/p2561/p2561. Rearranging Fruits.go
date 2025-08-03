package p2561

import (
	"sort"
)

func minCost1(basket1 []int, basket2 []int) int64 {

	n := len(basket1)
	sort.Ints(basket1)
	sort.Ints(basket2)

	minFruitCost := min(basket1[0], basket2[0])

	// Списки различающихся элементов
	var diff1, diff2 []int
	i, j := 0, 0

	for i < n && j < n {
		if basket1[i] == basket2[j] {
			i++
			j++
		} else if basket1[i] < basket2[j] {
			// Добавляем пару одинаковых значений basket1[i]
			if i+1 >= n || basket1[i] != basket1[i+1] {
				return -1 // нечётная разность — невозможно уравнять
			}
			diff1 = append(diff1, basket1[i])
			i += 2
		} else {
			// Добавляем пару одинаковых значений basket2[j]
			if j+1 >= n || basket2[j] != basket2[j+1] {
				return -1
			}
			diff2 = append(diff2, basket2[j])
			j += 2
		}
	}

	// Обработка оставшихся элементов
	for i < n {
		if i+1 >= n || basket1[i] != basket1[i+1] {
			return -1
		}
		diff1 = append(diff1, basket1[i])
		i += 2
	}
	for j < n {
		if j+1 >= n || basket2[j] != basket2[j+1] {
			return -1
		}
		diff2 = append(diff2, basket2[j])
		j += 2
	}

	if len(diff1) != len(diff2) {
		return -1
	}

	// Минимизация стоимости обмена
	sort.Ints(diff1)
	sort.Sort(sort.Reverse(sort.IntSlice(diff2)))

	var totalCost int64
	for i := 0; i < len(diff1); i++ {
		totalCost += int64(min(2*minFruitCost, min(diff1[i], diff2[i])))
	}

	return totalCost
}

func minCost(basket1 []int, basket2 []int) int64 {
	l := len(basket1)
	sort.Ints(basket1)
	sort.Ints(basket2)
	minFruitCost := min(basket1[0], basket2[0])

	arr1 := []int{}
	arr2 := []int{}

	index1, index2 := 0, 0
	for index1 < l && index2 < l {
		if basket1[index1] == basket2[index2] {
			index1++
			index2++
		} else if basket1[index1] < basket2[index2] {
			arr1 = append(arr1, basket1[index1])
			index1++
			if index1 == l || basket1[index1-1] != basket1[index1] {
				return -1
			}
			index1++
		} else {
			arr2 = append(arr2, basket2[index2])
			index2++

			if index2 == l || basket2[index2-1] != basket2[index2] {
				return -1
			}
			index2++
		}
	}

	for index1 < l {
		arr1 = append(arr1, basket1[index1])
		index1 += 2
	}

	for index2 < l {
		arr2 = append(arr2, basket2[index2])
		index2 += 2
	}

	if len(arr1) != len(arr2) {
		return -1
	}

	var res int64
	for i := 0; i < len(arr1); i++ {
		res += int64(min(2*minFruitCost, arr1[i], arr2[len(arr1)-1-i]))
	}

	return res

}

func minCost0(basket1 []int, basket2 []int) int64 {
	freq1, minFruitCost1 := getFreq(basket1)
	freq2, minFruitCost2 := getFreq(basket2)

	minFruitCost := min(minFruitCost1, minFruitCost2)

	surplusBasket1, surplusBasket2 := getSurplusBasket(freq1, freq2)
	if surplusBasket1 == nil {
		return -1
	}

	arr1 := make([]int, 0, len(surplusBasket1))
	arr2 := make([]int, 0, len(surplusBasket2))
	for i, c := range surplusBasket1 {
		for j := 0; j < c; j++ {
			arr1 = append(arr1, i)
		}
	}
	for i, c := range surplusBasket2 {
		for j := 0; j < c; j++ {
			arr2 = append(arr2, i)
		}
	}

	sort.Ints(arr1)
	sort.Sort(sort.Reverse(sort.IntSlice(arr2)))

	var res int64
	for i := 0; i < len(arr1); i++ {
		res += int64(min(2*minFruitCost, arr1[i], arr2[i]))
	}

	return res
}

func getFreq(basket []int) (map[int]int, int) {
	m := make(map[int]int)
	minFruitCost := basket[0]
	for _, num := range basket {
		m[num]++
		minFruitCost = min(minFruitCost, num)
	}
	return m, minFruitCost
}

func getSurplusBasket(freq1, freq2 map[int]int) (map[int]int, map[int]int) {
	keys := make(map[int]struct{})
	for k := range freq1 {
		keys[k] = struct{}{}
	}
	for k := range freq2 {
		keys[k] = struct{}{}
	}

	surplusBasket1 := make(map[int]int)
	surplusBasket2 := make(map[int]int)
	for key := range keys {
		total := freq1[key] + freq2[key]

		if total%2 == 1 {
			return nil, nil
		}
		ideal := total / 2
		if v := freq1[key]; v != ideal {
			diff := v - ideal
			if diff > 0 {
				surplusBasket1[key] = diff
			} else {
				surplusBasket2[key] = -diff
			}

		}
	}

	return surplusBasket1, surplusBasket2
}
