package p0679

import (
	"math"
)

func judgePoint24(cards []int) bool {
	var calc func(digits []float64) bool
	calc = func(digits []float64) bool {
		// Базовый случай: если в срезе осталось одно число, проверяем, равно ли оно 24
		if len(digits) == 1 {
			return math.Round(digits[0]*10) == 240.0
		}

		// Рекурсивный шаг: перебираем все возможные пары чисел
		for i := 0; i < len(digits); i++ {
			for j := i + 1; j < len(digits); j++ {
				// Создаем новый срез, который будет на 1 элемент меньше
				// (мы убираем два числа и добавляем одно)
				newDigits := make([]float64, 0, len(digits)-1)

				// Добавляем в новый срез все числа, кроме тех, что мы только что выбрали
				for k := 0; k < len(digits); k++ {
					if k != i && k != j {
						newDigits = append(newDigits, digits[k])
					}
				}

				// Теперь перебираем все операторы для выбранной пары чисел
				for _, op := range operators {
					// Выполняем операцию и добавляем результат в конец нового среза
					newDigitsWithResult := append(newDigits, e(digits[i], op, digits[j]))

					// Рекурсивно вызываем calculate с новым срезом
					if calc(newDigitsWithResult) {
						return true
					}

					// Для операции вычитания и деления нужно проверить
					// обратный порядок, так как они не коммутативны (a-b != b-a)
					if op == '-' || op == '/' {
						newDigitsWithResult = append(newDigits, e(digits[j], op, digits[i]))
						if calc(newDigitsWithResult) {
							return true
						}
					}
				}
			}
		}

		return false

	}
	digits := make([]float64, len(cards))
	for i, card := range cards {
		digits[i] = float64(card)
	}

	return calc(digits)
}

func judgePoint241(cards []int) bool {
	l := len(cards)
	currentPermutation := make([]float64, l)
	var permutations func(used []bool, level int) bool
	permutations = func(used []bool, level int) bool {
		if level == l {
			if calculate(currentPermutation) {
				//fmt.Println(currentPermutation)
				return true
			}
			return false
		}
		for i := 0; i < l; i++ {
			if !used[i] {
				used[i] = true
				currentPermutation[level] = float64(cards[i])
				if permutations(used, level+1) {
					return true
				}
				used[i] = false
			}
		}
		return false
	}
	return permutations(make([]bool, l), 0)
}

var operators = [4]byte{'+', '-', '/', '*'}

func calculate(digits []float64) bool {
	if len(digits) == 1 && math.Round(digits[0]*10) == 240.0 {
		return true
	}
	for i := 0; i < len(digits)-1; i++ {
		newDigit := make([]float64, len(digits)-1)
		for j := 0; j < i; j++ {
			newDigit[j] = digits[j]
		}
		for j := i + 2; j < len(digits); j++ {
			newDigit[j-1] = digits[j]
		}
		for _, op := range operators {
			newDigit[i] = e(digits[i], op, digits[i+1])
			if calculate(newDigit) {
				return true
			}
		}
	}
	return false
}

func judgePoint240(cards []int) bool {
	var a, b, c, d float64

	a = float64(cards[0])
	b = float64(cards[1])
	c = float64(cards[2])
	d = float64(cards[3])
	if run2(a, b, c, d) {
		return true
	}
	if run2(a, b, d, c) {
		return true
	}
	if run2(a, c, b, d) {
		return true
	}
	if run2(a, c, d, b) {
		return true
	}
	if run2(a, d, b, c) {
		return true
	}
	if run2(a, d, c, b) {
		return true
	}

	if run2(b, a, c, d) {
		return true
	}
	if run2(b, a, d, c) {
		return true
	}
	if run2(b, c, a, d) {
		return true
	}
	if run2(b, c, d, a) {
		return true
	}
	if run2(b, d, a, c) {
		return true
	}
	if run2(b, d, c, a) {
		return true
	}

	if run2(c, a, b, d) {
		return true
	}
	if run2(c, a, d, b) {
		return true
	}
	if run2(c, b, a, d) {
		return true
	}
	if run2(c, b, d, a) {
		return true
	}
	if run2(c, d, a, b) {
		return true
	}
	if run2(c, d, b, a) {
		return true
	}

	if run2(d, a, b, c) {
		return true
	}
	if run2(d, a, c, b) {
		return true
	}
	if run2(d, b, a, c) {
		return true
	}
	if run2(d, b, c, a) {
		return true
	}
	if run2(d, c, a, b) {
		return true
	}
	if run2(d, c, b, a) {
		return true
	}

	return false
}

func run2(a, b, c, d float64) bool {
	var operators = [4]byte{'+', '-', '/', '*'}
	for _, oAB := range operators {
		for _, oBC := range operators {
			for _, oCD := range operators {
				if run(a, b, c, d, oAB, oBC, oCD) {
					return true
				}
			}
		}
	}
	return false
}

func e(digitLeft float64, operation byte, digitRight float64) float64 {
	switch operation {
	case '-':
		return digitLeft - digitRight
	case '*':
		return digitLeft * digitRight
	case '/':
		return digitLeft / digitRight
	default:
		return digitLeft + digitRight
	}
}

func run(a, b, c, d float64, oAB, oBC, oCD byte) bool {
	if math.Round(e(a, oAB, e(b, oBC, e(c, oCD, d)))*10) == 240.0 {
		return true
	}
	if math.Round(e(a, oAB, e(e(b, oBC, c), oCD, d))*10) == 240.0 {
		return true
	}
	if math.Round(e(e(a, oAB, b), oBC, e(c, oCD, d))*10) == 240.0 {
		return true
	}
	if math.Round(e(e(a, oAB, e(b, oBC, c)), oCD, d)*10) == 240.0 {
		return true
	}
	if math.Round(e(e(e(a, oAB, b), oBC, c), oCD, d)*10) == 240.0 {
		return true
	}
	return false
}

//func run(a, b, c, d float64, oAB, oBC, oCD byte) bool {
//	//a b c d
//	if math.Round(e(e(e(a, oAB, b), oBC, c), oCD, d)*10) == 240.0 {
//		return true
//	}
//
//	//(a b) c d
//	if math.Round(e(e(e(a, oAB, b), oBC, c), oCD, d)*10) == 240.0 {
//		return true
//	}
//
//	//a (b c) d
//	if math.Round(e(e(a, oAB, e(b, oBC, c)), oCD, d)*10) == 240.0 {
//		return true
//	}
//
//	//a b (c d)
//	if math.Round(e(e(a, oAB, b), oBC, e(c, oCD, d))*10) == 240.0 {
//		return true
//	}
//
//	//(a b) (c d)
//	if math.Round(e(e(a, oAB, b), oBC, e(c, oCD, d))*10) == 240.0 {
//		return true
//	}
//
//	//(a b c) d
//	if math.Round(e(e(e(a, oAB, b), oBC, c), oCD, d)*10) == 240.0 {
//		return true
//	}
//
//	//((a b) c) d
//	if math.Round(e(e(e(a, oAB, b), oBC, c), oCD, d)*10) == 240.0 {
//		return true
//	}
//
//	//(a (b c)) d
//	if math.Round(e(e(a, oAB, e(b, oBC, c)), oCD, d)*10) == 240.0 {
//		return true
//	}
//
//	//a (b c d)
//	if math.Round(e(a, oAB, e(e(b, oBC, c), oCD, d))*10) == 240.0 {
//		return true
//	}
//
//	//a ((b c) d)
//	if math.Round(e(a, oAB, e(e(b, oBC, c), oCD, d))*10) == 240.0 {
//		return true
//	}
//
//	//a (b (c d))
//	if math.Round(e(a, oAB, e(b, oBC, e(c, oCD, d)))*10) == 240.0 {
//		return true
//	}
//
//	return false
//}

/*


   a b c d
   4 знака +-/*
   3 позиции для знакаов, т.е. 4^3=64
   скобки
   a b c d
   (a b) c d
   a (b c) d
   a b (c d)
   (a b) (c d)
   (a b c) d
   ((a b) c) d
   (a (b c)) d
   a (b c d)
   a ((b c) d)
   a (b (c d))
   64*11 = 704


*/

// 5 1 9 1

// 3, 7, 4, 6
