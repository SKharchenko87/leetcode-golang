package w

var bitToRem = [8]int{1, 7, 11, 13, 17, 19, 23, 29}

var remToBit = [30]int{
	-1, 0, -1, -1, -1, -1, -1, 1, -1, -1, -1, 2, -1, 3, -1, -1,
	-1, 4, -1, 5, -1, -1, -1, 6, -1, -1, -1, -1, -1, 7,
}

var wheelSteps = [8]int{6, 4, 2, 4, 2, 4, 6, 2}

type Jump struct {
	byteMove int
	nextBit  uint8
}

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	size := (limit / 30) + 1
	data := make([]byte, size)
	for i := range data {
		data[i] = 0xFF
	}
	data[0] &= ^byte(1 << 0)

	for n := 7; n*n <= limit; n++ {
		rem := n % 30
		bitIdx := remToBit[rem]

		if bitIdx == -1 || data[n/30]&(1<<uint(bitIdx)) == 0 {
			continue
		}

		jumps := genJumpsForP(n)
		strikeOut(data, n, limit, jumps)
	}

	return collectPrimes(data, limit)
}

func genJumpsForP(p int) [8]Jump {
	var jumps [8]Jump

	for i := 0; i < 8; i++ {
		currentRem := bitToRem[i]

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

		jumps[i] = Jump{
			byteMove: nextFull/30 - currentRem/30,
			nextBit:  uint8(nextBit),
		}
	}
	return jumps
}

func strikeOut(data []byte, p int, limit int, jumps [8]Jump) {
	n := p * p
	if n > limit {
		return
	}

	limitBytes := len(data)
	byteIdx := n / 30
	bitIdx := uint8(remToBit[n%30])

	for byteIdx < limitBytes {
		data[byteIdx] &= ^(1 << bitIdx)
		j := jumps[bitIdx]
		byteIdx += j.byteMove
		bitIdx = j.nextBit
	}
}

func collectPrimes(data []byte, limit int) []int {
	primes := make([]int, 3, limit/10)
	primes[0] = 2
	primes[1] = 3
	primes[2] = 5

	for i := 0; i < len(data); i++ {
		b := data[i]
		if b == 0 {
			continue
		}
		for bit := 0; bit < 8; bit++ {
			if b&(1<<uint(bit)) != 0 {
				p := int(i)*30 + bitToRem[bit]
				if p <= limit && p > 1 {
					primes = append(primes, p)
				}
			}
		}
	}
	return primes
}

var arr [1e5 + 1]int

func init() {
	var limit = 50000

	primes := Sieve(limit)

	for i := 0; i < len(primes)-1; i++ {
		d := primes[i] * primes[i] * primes[i]
		if d > 1e5 {
			break
		}
		arr[d] = (1 + primes[i]) * (1 + primes[i]*primes[i])
	}

	for i := 0; i < len(primes)-1; i++ {
		if primes[i]*primes[i+1] > 1e5 {
			break
		}
		for j := i + 1; j < len(primes); j++ {
			d := primes[i] * primes[j]
			if d > 1e5 {
				break
			}
			arr[d] = d + primes[i] + primes[j] + 1
		}
	}
}

func sumFourDivisors(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		result += arr[nums[i]]
	}
	return result
}
