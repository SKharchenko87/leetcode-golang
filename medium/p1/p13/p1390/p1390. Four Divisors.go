package p1390

import (
	"fmt"
	"time"
)

/*
1 	2 	3 	6 		min
1 	2 	b 	<1e5 	max
b простое число
2*b ~ 1e5
b<50000

5133 — количество простых чисел в диапазоне от 1 до 50 000.
*/

// bitToRem: индекс бита (0-7) -> остаток от 30 (спицы колеса)
var bitToRem = [8]uint64{1, 7, 11, 13, 17, 19, 23, 29}

// remToBit: остаток % 30 -> индекс бита (0-7). -1 для чисел, кратных 2, 3, 5
var remToBit = [30]int{
	-1, 0, -1, -1, -1, -1, -1, 1, -1, -1, -1, 2, -1, 3, -1, -1,
	-1, 4, -1, 5, -1, -1, -1, 6, -1, -1, -1, -1, -1, 7,
}

// wheelSteps: расстояния между спицами колеса
var wheelSteps = [8]uint64{6, 4, 2, 4, 2, 4, 6, 2}

type Jump struct {
	byteMove int   // На сколько байт сдвинуться вперед
	nextBit  uint8 // Индекс следующего бита (0-7)
}

func Sieve(limit uint64) []uint64 {
	if limit < 2 {
		return []uint64{}
	}

	size := (limit / 30) + 1
	data := make([]byte, size)
	for i := range data {
		data[i] = 0xFF
	}
	data[0] &= ^byte(1 << 0) // 1 - не простое

	for n := uint64(7); n*n <= limit; n++ {
		rem := int(n % 30)
		bitIdx := remToBit[rem]

		if bitIdx == -1 || data[n/30]&(1<<uint(bitIdx)) == 0 {
			continue
		}

		jumps := genJumpsForP(n)
		strikeOut(data, n, limit, jumps)
	}

	return collectPrimes(data, limit)
}

func genJumpsForP(p uint64) [8]Jump {
	var jumps [8]Jump

	// Для каждой возможной спицы i (0-7), на которой может оказаться кратное p:
	for i := 0; i < 8; i++ {
		currentRem := bitToRem[i]

		// Нам нужно найти такое k (из колеса), что (p * k) % 30 == currentRem.
		// Это сложно. Проще: если мы сейчас на остатке currentRem,
		// то следующий остаток в последовательности кратных p будет:
		// (currentRem + p * (шаг колеса в правильной позиции)) % 30.

		// Чтобы понять, какой сейчас "шаг колеса", нужно знать,
		// какому числу k соответствует currentRem.
		// k = (currentRem * p^-1) mod 30.

		// Но есть способ проще: мы знаем, что последовательность остатков
		// p * 1, p * 7, p * 11... - это просто перестановка 1, 7, 11...
		// Поэтому мы можем просто найти следующий остаток в этой последовательности.

		// Вычисляем шаг: p * (дистанция до следующего числа в колесе)
		// Нам нужно "повернуть" wheelSteps в зависимости от p.
		// stepIdx := (i * int(p%30)) % 8 // Это грубое приближение, заменим на честный поиск

		// Честный поиск следующей спицы для p:
		// Мы находимся в остатке currentRem. Следующее кратное: currentRem + p*step?
		// Нет. p*1, p*7, p*11...
		// Дистанция между ними: p*(7-1), p*(11-7)...
		// Эти дистанции p*6, p*4, p*2... всегда возвращают нас в колесо!

		// Нам нужно сопоставить i (индекс текущего остатка) с индексом шага wheelSteps.
		// Индекс шага j такой, что bitToRem[j] * p % 30 == currentRem.
		var wheelPos int
		for j := 0; j < 8; j++ {
			if (bitToRem[j]*p)%30 == currentRem {
				wheelPos = j
				break
			}
		}

		move := p * wheelSteps[wheelPos]
		nextFull := currentRem + move
		nextRem := nextFull % 30
		nextBit := remToBit[nextRem]

		if nextBit == -1 {
			panic(fmt.Sprintf("Logic error: p=%d, currentRem=%d, move=%d, nextRem=%d", p, currentRem, move, nextRem))
		}

		jumps[i] = Jump{
			byteMove: int(nextFull/30 - currentRem/30),
			nextBit:  uint8(nextBit),
		}
	}
	return jumps
}

func strikeOut(data []byte, p uint64, limit uint64, jumps [8]Jump) {
	n := p * p
	if n > limit {
		return
	}

	limitBytes := len(data)
	byteIdx := int(n / 30)
	bitIdx := uint8(remToBit[n%30])

	for byteIdx < limitBytes {
		data[byteIdx] &= ^(1 << bitIdx)
		j := jumps[bitIdx]
		byteIdx += j.byteMove
		bitIdx = j.nextBit
	}
}

func collectPrimes(data []byte, limit uint64) []uint64 {
	primes := make([]uint64, 0, limit/10)
	if limit >= 2 {
		primes = append(primes, 2)
	}
	if limit >= 3 {
		primes = append(primes, 3)
	}
	if limit >= 5 {
		primes = append(primes, 5)
	}

	for i := 0; i < len(data); i++ {
		b := data[i]
		if b == 0 {
			continue
		}
		for bit := 0; bit < 8; bit++ {
			if b&(1<<uint(bit)) != 0 {
				p := uint64(i)*30 + bitToRem[bit]
				if p <= limit && p > 1 {
					primes = append(primes, p)
				}
			}
		}
	}
	return primes
}

func run() int {
	var limit uint64 = 50000
	start := time.Now()

	primes := Sieve(limit)

	elapsed := time.Since(start)

	// Контрольная сумма: до 50,000 должно быть ровно 5133 простых числа
	fmt.Printf("Диапазон: 0 - %d\n", limit)
	fmt.Printf("Найдено простых чисел: %d\n", len(primes))
	fmt.Printf("Время выполнения: %s\n", elapsed)

	if len(primes) > 10 {
		fmt.Printf("Первые 10: %v\n", primes[:10])
		fmt.Printf("Последние 5: %v\n", primes[len(primes)-5:])
	}

	m := map[int]int{}

	for i := 0; i < len(primes)-1; i++ {
		if primes[i]*primes[i+1] > 1e5 {
			break
		}
		for j := i + 1; j < len(primes); j++ {
			d := int(primes[i] * primes[j])
			if d > 1e5 {
				break
			}
			m[d] = d + int(primes[i]+primes[j]) + 1
		}
	}

	for i := 0; i < len(primes)-1; i++ {
		d := primes[i] * primes[i] * primes[i]
		if d > 1e5 {
			break
		}

		m[int(d)] = int(1 + primes[i] + primes[i]*primes[i] + d)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Printf("Найдено простых чисел: %d\n", len(m))

	return 0
}

var m map[int]int

func init() {
	m = make(map[int]int, 30000)
	var limit uint64 = 50000

	primes := Sieve(limit)

	if len(primes) > 10 {
		fmt.Printf("Первые 10: %v\n", primes[:10])
		fmt.Printf("Последние 5: %v\n", primes[len(primes)-5:])
	}

	for i := 0; i < len(primes)-1; i++ {
		d := primes[i] * primes[i] * primes[i]
		if d > 1e5 {
			break
		}

		m[int(d)] = int(1 + primes[i] + primes[i]*primes[i] + d)
	}

	for i := 0; i < len(primes)-1; i++ {
		if primes[i]*primes[i+1] > 1e5 {
			break
		}
		for j := i + 1; j < len(primes); j++ {
			d := int(primes[i] * primes[j])
			if d > 1e5 {
				break
			}
			m[d] = d + int(primes[i]+primes[j]) + 1
		}
	}
}

func sumFourDivisors(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		result += m[nums[i]]
	}
	return result
}

func debugMap() {
	invalidCount := 0
	for num := range m {
		divs := []int{}
		// Честно ищем все делители перебором
		for i := 1; i <= num; i++ {
			if num%i == 0 {
				divs = append(divs, i)
			}
		}

		// Если делителей НЕ 4, значит это число лишнее в нашей карте
		if len(divs) != 4 {
			invalidCount++
			if invalidCount <= 10 { // Выведем первые 10 для примера
				fmt.Printf("Лишнее число: %d, Делители: %v (всего %d)\n", num, divs, len(divs))
			}
		}
	}
	fmt.Printf("Всего в карте: %d\n", len(m))
	fmt.Printf("Из них некорректных: %d\n", invalidCount)
	fmt.Printf("Чистых чисел с 4 делителями: %d\n", len(m)-invalidCount)
	maxN := 0
	for k := range m {
		if k > maxN {
			maxN = k
		}
	}
	fmt.Printf("Максимальное число в карте: %d\n", maxN)
}
