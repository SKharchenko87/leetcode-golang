package p0260

func singleNumber(nums []int) []int {
	// xor всех элементов массива будет равен xor двух элементов ответа xorNums = res1^res
	var xorNums int
	for _, num := range nums {
		xorNums ^= num
	}

	// В xorNums нужно найти specificBitPosition - позиция какого-нибудь бита равного 1.
	// Бит в этой позиции будет различен у чисел res1 и res2
	specificBitPosition := 0
	for ; specificBitPosition < 32; specificBitPosition++ {
		if xorNums>>specificBitPosition&1 == 1 {
			break
		}
	}

	// Еще раз пройдемся xor'ом по элементам массива, с условием, что на позиции specificBitPosition бит равен 1.
	//Результат этой операции сохраним в xorNumsWithSpecificBit
	var xorNumsWithSpecificBit int
	for _, num := range nums {
		if (num >> specificBitPosition & 1) == 1 {
			xorNumsWithSpecificBit ^= num
		}
	}

	// Т.о. в xorNumsWithSpecificBit будет лежать res1, т.к. есть условие,
	// что бит равным 1 на позиции specificBitPosition и он различен у чисел res1 и res2
	res := []int{xorNumsWithSpecificBit, xorNums ^ xorNumsWithSpecificBit}
	return res
}

func singleNumber0(nums []int) []int {
	m := make(map[int]bool, 15000)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = true
		}
	}
	res := []int{}
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}
